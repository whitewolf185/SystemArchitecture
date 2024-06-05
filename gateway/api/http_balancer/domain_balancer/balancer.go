package domain_balancer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer/circuit_breaker"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

type CircuitBreaker interface {
	GetStatus() circuit_breaker.CircuitStatus
	CountError()
	CountHalfOpenedRequests() int
}

type balancer struct {
	apiName string
	client  http.Client
	apiHost string

	circuidBreaker CircuitBreaker
	keydbClient    *redis.Client
}

func New(apiName string, keydbClient *redis.Client, circuidBreaker CircuitBreaker) balancer {
	return balancer{
		apiName:        apiName,
		client:         http.Client{},
		apiHost:        fmt.Sprintf("%s:%s", config.GetValue(config.DomainServiceHost), config.GetValue(config.DomainServicePort)),
		keydbClient:    keydbClient,
		circuidBreaker: circuidBreaker,
	}
}

func (b balancer) circuidController(ctx context.Context, request *http.Request, keydbKey string) ([]byte, error) {
	var (
		resBody []byte
		err     error
	)
	circuidStatus := b.circuidBreaker.GetStatus()
	if circuidStatus == circuit_breaker.Opened {
		keydbResult, keydbErr := b.keydbClient.Get(ctx, keydbKey).Result()
		if keydbErr != nil {
			return nil, fmt.Errorf("keydb error: %s; service error: %w", keydbErr.Error(), err)
		}
		resBody = []byte(keydbResult)
		logrus.Info("the result was used by keydb")
	}
	if circuidStatus == circuit_breaker.Closed {
		resBody, err = http_balancer.RequestSender(request)
		if err != nil {
			if errors.As(err, &customerrors.ErrCodes{}) {
				return nil, err
			}
			b.circuidBreaker.CountError()
			return nil, err
		}
	}

	if circuidStatus == circuit_breaker.HalfOpened {
		requestsCounter := b.circuidBreaker.CountHalfOpenedRequests()
		if requestsCounter%2 == 0 {
			resBody, err = http_balancer.RequestSender(request)
			if err != nil {
				if errors.As(err, &customerrors.ErrCodes{}) {
					return nil, err
				}
				b.circuidBreaker.CountError()
				return nil, err
			}
		} else {
			keydbResult, keydbErr := b.keydbClient.Get(ctx, keydbKey).Result()
			if keydbErr != nil {
				return nil, fmt.Errorf("keydb error: %s; service error: %w", keydbErr.Error(), err)
			}
			resBody = []byte(keydbResult)
			logrus.Info("the result was used by keydb")
		}
	}

	return resBody, nil
}

// GetCompanionInfo - получение информации о попутчике
func (b balancer) GetCompanionInfo(ctx context.Context, req *domain_domain.GetCompanionInfoRequest) ([]domain_domain.Companion, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/GetCompanionInfo",
		RawQuery: query.Encode(),
	}
	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	keydbKey := fmt.Sprintf("%s|%s", u.Path, req.ClientID)
	resBody, err := b.circuidController(ctx, request, keydbKey)
	if err != nil {
		return nil, err
	}

	var result []domain_domain.Companion
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[GetCompanionInfo]: cannot unmarshal response: %w", err)
	}
	logrus.Info(result)
	b.setValueToKeydb(ctx, keydbKey, result)
	return result, nil
}

// GetRouteInfo - получение информации о маршруте
func (b balancer) GetRouteInfo(ctx context.Context, req *domain_domain.GetRouteInfoRequest) ([]domain_domain.Route, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/GetRouteInfo",
		RawQuery: query.Encode(),
	}
	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	keydbKey := fmt.Sprintf("%s|%s|%s", u.Path, req.ClientID, req.OneOfPath)
	resBody, err := b.circuidController(ctx, request, keydbKey)
	if err != nil {
		return nil, err
	}

	var result []domain_domain.Route
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[GetRouteInfo]: cannot unmarshal response: %w", err)
	}
	logrus.Info(result)
	b.setValueToKeydb(ctx, keydbKey, result)
	return result, nil
}

// CreateRoute - создание маршрута
func (b balancer) CreateRoute(ctx context.Context, req *domain_domain.CreateRouteRequest) (*domain_domain.Route, error) {
	if req == nil || req.Payload == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	requestBodyBytes, err := json.Marshal(req.Payload)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal request: %w", err)
	}

	u := url.URL{
		Scheme: "http",
		Host:   b.apiHost,
		Path:   "api/" + b.apiName + "/CreateRoute",
	}
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_domain.Route
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[CreateRoute]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateCompanion - создание попутчика
func (b balancer) CreateCompanion(ctx context.Context, req *domain_domain.CreateCompanionRequest) (*domain_domain.Companion, error) {
	if req == nil || req.Payload == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	requestBodyBytes, err := json.Marshal(req.Payload)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal request: %w", err)
	}

	u := url.URL{
		Scheme: "http",
		Host:   b.apiHost,
		Path:   "api/" + b.apiName + "/CreateCompanion",
	}
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_domain.Companion
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[CreateCompanion]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}

// DeleteRoute - удалить маршрут
func (b balancer) DeleteRoute(ctx context.Context, req *domain_domain.DeleteRouteRequest) error {
	if req == nil {
		return customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/DeleteRoute",
		RawQuery: query.Encode(),
	}
	request, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return fmt.Errorf("cannot build a request: %w", err)
	}

	_, err = http_balancer.RequestSender(request)
	return err
}

// DeleteCompanion - удалить попутчика
func (b balancer) DeleteCompanion(ctx context.Context, req *domain_domain.DeleteCompanionRequest) error {
	if req == nil {
		return customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/DeleteCompanion",
		RawQuery: query.Encode(),
	}
	request, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return fmt.Errorf("cannot build a request: %w", err)
	}

	_, err = http_balancer.RequestSender(request)

	return err
}

func (b balancer) setValueToKeydb(ctx context.Context, request string, response any) error {
	marshaledBinary, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("cannot marshal response while inserting into key-db: %w", err)
	}
	err = b.keydbClient.Set(ctx, request, marshaledBinary, time.Hour).Err()
	if err != nil {
		return fmt.Errorf("cannot add response into redis: %w", err)
	}
	return nil
}
