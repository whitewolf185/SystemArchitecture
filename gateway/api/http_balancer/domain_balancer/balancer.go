package domain_balancer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

type balancer struct {
	apiName string
	client  http.Client
	apiHost string

	keydbClient *redis.Client
}

func New(apiName string, servicePort string, keydbClient *redis.Client) balancer {
	return balancer{
		apiName:     apiName,
		client:      http.Client{},
		apiHost:     fmt.Sprintf("%s:%s", config.GetValue(config.DomainServiceHost), config.GetValue(config.DomainServicePort)),
		keydbClient: keydbClient,
	}
}

// GetCompanionInfo - получение информации о попутчике
func (b balancer) GetCompanionInfo(ctx context.Context, req *domain_domain.GetCompanionInfoRequest) ([]domain_domain.Companion, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := b.prepareQuery(req)
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
	resBody, err := b.requestSender(request)
	if err != nil {
		if errors.As(err, &customerrors.ErrCodes{}) {
			return nil, err
		}
		keydbResult, keydbErr := b.keydbClient.Get(ctx, keydbKey).Result()
		if keydbErr != nil {
			return nil, fmt.Errorf("keydb error: %s; service error: %w", keydbErr.Error(), err)
		}
		resBody = []byte(keydbResult)
		logrus.Info("the result was used by keydb")
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

	query := b.prepareQuery(req)
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
	resBody, err := b.requestSender(request)
	if err != nil {
		if errors.As(err, &customerrors.ErrCodes{}) {
			return nil, err
		}
		keydbResult, keydbErr := b.keydbClient.Get(ctx, keydbKey).Result()
		if keydbErr != nil {
			return nil, fmt.Errorf("keydb error: %s; service error: %w", keydbErr.Error(), err)
		}
		resBody = []byte(keydbResult)
		logrus.Info("the result was used by keydb")
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

	resBody, err := b.requestSender(request)
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

	resBody, err := b.requestSender(request)
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

	query := b.prepareQuery(req)
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

	_, err = b.requestSender(request)
	return err
}

// DeleteCompanion - удалить попутчика
func (b balancer) DeleteCompanion(ctx context.Context, req *domain_domain.DeleteCompanionRequest) error {
	if req == nil {
		return customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := b.prepareQuery(req)
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

	_, err = b.requestSender(request)

	return err
}

func (b balancer) prepareQuery(request any) url.Values {
	result := make(url.Values)
	if request == nil {
		return result
	}
	fields := reflect.TypeOf(request).Elem()
	// fmt.Println(fields.NumField())
	for i := 0; i < fields.NumField(); i++ {
		currField := fields.Field(i)
		// fmt.Println(currField)
		// fmt.Println(currField.Tag.Get("in"))
		if strings.HasPrefix(currField.Tag.Get("in"), "query=") {
			queryKey := strings.Split(currField.Tag.Get("in"), "=")[1]
			queryValue := reflect.ValueOf(request).Elem().Field(i).String()
			// fmt.Printf("queryKey: %s, queryValue: %s\n", queryKey, queryValue)
			result.Set(queryKey, queryValue)
		}
	}
	return result
}

func (b balancer) requestSender(request *http.Request) ([]byte, error) {
	client := &http.Client{}

	request.Header.Add("Content-Type", "application/json")
	logrus.Info(request)

	res, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("gateway: unkwodn error: %w", err)
	}
	if res == nil {
		return nil, fmt.Errorf("empty response")
	}
	logrus.Infof("status code: %d, type: %T", res.StatusCode, res.StatusCode)
	switch res.StatusCode {
	case http.StatusBadRequest:
		return nil, customerrors.CodesBadRequest(err)
	case http.StatusInternalServerError:
		return nil, err
	case http.StatusUnauthorized:
		return nil, customerrors.CodesUnauthorized(err)
	case http.StatusNotFound:
		return nil, customerrors.CodesNotFound(err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gateway: cannot ready body: %w", err)
	}
	if resBody == nil {
		return nil, nil
	}
	logrus.Info(string(resBody))

	return resBody, nil
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
