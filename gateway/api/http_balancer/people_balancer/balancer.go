package people_balancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	domain_client_service "github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

type balancer struct {
	apiName    string
	apiBaseUrl string
}

func New(apiName string, servicePort string) balancer {
	return balancer{
		apiName:    apiName,
		apiBaseUrl: fmt.Sprintf("http://%s:%s/api/", config.GetValue(config.ClientServiceHost), servicePort),
	}
}

// GetClientByID - получение пользователя по его id
func (b balancer) GetClientByID(ctx context.Context, req *domain_client_service.GetPersonByIDRequest) (*domain_client_service.Person, error) {
	return nil, customerrors.CodesNotImplemented(fmt.Errorf("method not implemented"))
}

// CreateUser - создание пользователя с указанием его пароля и логина
func (b balancer) CreateUser(ctx context.Context, req *domain_client_service.CreateUserRequest) (*domain_client_service.Person, error) {
	return nil, customerrors.CodesNotImplemented(fmt.Errorf("method not implemented"))
}

// SearchUserByUserName - поиск пользователя по substring его userName. В итоге появляется массив
func (b balancer) SearchUserByUserName(ctx context.Context, req *domain_client_service.SearchUserByUserNameRequest) ([]domain_client_service.Person, error) {
	return nil, customerrors.CodesNotImplemented(fmt.Errorf("method not implemented"))
}

// DeleteUserByID - удаляет пользователя по его id
func (b balancer) DeleteUserByID(ctx context.Context, req *domain_client_service.DeleteUserByIDRequest) (*domain_client_service.Person, error) {
	return nil, customerrors.CodesNotImplemented(fmt.Errorf("method not implemented"))
}

// Login - ручка логина. На выход отдается jwt токен
func (b balancer) Login(ctx context.Context, req *domain_client_service.LoginRequest) (*domain_client_service.LoginResponse, error) {
	if req == nil || req.Payload == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}
	resBody, err := b.requestBuilder(req.Payload, "Login")
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

func (b balancer) requestBuilder(req any, path string) ([]byte, error) {
	client := &http.Client{}
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}
	requestBodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal request: %w", err)
	}
	logrus.Info(string(requestBodyBytes))
	request, err := http.NewRequest(http.MethodPost, b.apiBaseUrl+b.apiName+"/"+path, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot build a request: %w", err)
	}
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
