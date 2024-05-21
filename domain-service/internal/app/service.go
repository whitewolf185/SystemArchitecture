package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @title Client service
// @version @0.9
// @description swagger для api gateway

// @basePath /api

type Repo interface {
	InsertRoute(ctx context.Context, dataToInsert interface{}) error
	InsertCompanion(ctx context.Context, dataToInsert interface{}) error

	DeleteRoute(ctx context.Context, idToDelete string) error
	DeleteCompanion(ctx context.Context, idToDelete string) error

	GetCompanions(ctx context.Context, filter domain.Companion) ([]domain.Companion, error)
	GetRoutes(ctx context.Context, filter domain.Route) ([]domain.Route, error)
}

// Implementation структура для реализации различных ручек
type Implementation struct {
	repo Repo
}

// NewImplementation конструктор для Implementation
func NewImplementation(repo Repo) (*Implementation, error) {
	return &Implementation{
		repo: repo,
	}, nil
}
