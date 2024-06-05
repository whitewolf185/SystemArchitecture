package people_balancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	domain_client_service "github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/http_balancer"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

type balancer struct {
	apiName string
	client  http.Client
	apiHost string
}

func New(apiName string) balancer {
	return balancer{
		apiName: apiName,
		client:  http.Client{},
		apiHost: fmt.Sprintf("%s:%s", config.GetValue(config.ClientServiceHost), config.GetValue(config.ClientServicePort)),
	}
}

// GetClientByID - получение пользователя по его id
func (b balancer) GetClientByID(ctx context.Context, req *domain_client_service.GetPersonByIDRequest) (*domain_client_service.Person, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/GetClientByID",
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_client_service.Person
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[Login]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}

// CreateUser - создание пользователя с указанием его пароля и логина
func (b balancer) CreateUser(ctx context.Context, req *domain_client_service.CreateUserRequest) (*domain_client_service.Person, error) {
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
		Path:   "api/" + b.apiName + "/CreateUser",
	}

	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_client_service.Person
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[Login]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}

// SearchUserByUserName - поиск пользователя по substring его userName. В итоге появляется массив
func (b balancer) SearchUserByUserName(ctx context.Context, req *domain_client_service.SearchUserByUserNameRequest) ([]domain_client_service.Person, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/SearchUserByUserName",
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result []domain_client_service.Person
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[Login]: cannot unmarshal response: %w", err)
	}
	return result, nil
}

// DeleteUserByID - удаляет пользователя по его id
func (b balancer) DeleteUserByID(ctx context.Context, req *domain_client_service.DeleteUserByIDRequest) (*domain_client_service.Person, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}

	query := http_balancer.PrepareQuery(req)
	u := url.URL{
		Scheme:   "http",
		Host:     b.apiHost,
		Path:     "api/" + b.apiName + "/DeleteUserByID",
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_client_service.Person
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[DeleteUserByID]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}

// Login - ручка логина. На выход отдается jwt токен
func (b balancer) Login(ctx context.Context, req *domain_client_service.LoginRequest) (*domain_client_service.LoginResponse, error) {
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
		Path:   "api/" + b.apiName + "/Login",
	}

	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}

	resBody, err := http_balancer.RequestSender(request)
	if err != nil {
		return nil, err
	}

	var result domain_client_service.LoginResponse
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("gateway[Login]: cannot unmarshal response: %w", err)
	}
	return &result, nil
}
