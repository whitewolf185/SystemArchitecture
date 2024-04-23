package app

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/gen/client_service/public/model"
)

// @title Client service
// @version @0.9
// @description swagger для api к клиентскому сервису

// @basePath /api

type (
	PersonRepo interface {
		GetUserByID(ctx context.Context, clientID uuid.UUID) (*model.Persons, error)
		GetUserByUserNameIn(ctx context.Context, username string) ([]model.Persons, error)
		DeleteUserByID(ctx context.Context, clientID uuid.UUID) (*model.Persons, error)
		CreateUser(ctx context.Context, req model.Persons) (*model.Persons, error)
		GetUserByName(ctx context.Context, req model.Persons) (*model.Persons, error)
	}
)

// Implementation структура для реализации различных ручек
type Implementation struct {
	personRepo PersonRepo
}

// NewImplementation конструктор для Implementation
func NewImplementation(
	personRepo PersonRepo) (*Implementation, error) {
	return &Implementation{
		personRepo: personRepo,
	}, nil
}

func unescapeUrl(escapedUrl string) (string, error) {
	result, err := url.QueryUnescape(escapedUrl)
	if err != nil {
		return "", fmt.Errorf("unescape error: %w", err)
	}
	return result, nil
}
