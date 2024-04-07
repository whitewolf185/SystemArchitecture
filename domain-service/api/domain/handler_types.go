package domain

import (
	"context"
)

type HandlerType int

// route -- маршрут
// companion -- попутчик

//go:generate stringer -type=HandlerType -output=./handler_types_string.go
const (
	GetCompanionInfo HandlerType = iota
	GetRouteInfo
	CreateRoute
	CreateCompanion
	DeleteRoute
	DeleteCompanion
)

// Интерфейс со всеми ручками сервиса
type Handlers interface {
	// GetCompanionInfo - получение информации о попутчике
	GetCompanionInfo(ctx context.Context, req *GetCompanionInfoRequest) (*Companion, error)
	// GetRouteInfo - получение информации о маршруте
	GetRouteInfo(ctx context.Context, req *GetRouteInfoRequest) (*Route, error)
	// CreateRoute - создание маршрута
	CreateRoute(ctx context.Context, req *CreateRouteRequest) (*Route, error)
	// CreateCompanion - создание попутчика
	CreateCompanion(ctx context.Context, req *CreateCompanionRequest) (*Companion, error)
	// DeleteRoute - удалить маршрут
	DeleteRoute(ctx context.Context, req *DeleteRouteRequest) error
	// DeleteCompanion - удалить попутчика
	DeleteCompanion(ctx context.Context, req *DeleteCompanionRequest) error
}

type Companion struct {
}
type Route struct {
}

type GetCompanionInfoRequest struct {
	ClientID string `in:"query=client_id"`
}

type GetRouteInfoRequest struct {
	ClientID string `in:"query=client_id"`
}

type CreateRouteRequestPayload struct {
	ClientID string `json:"client_id" bson:"_id"`
}
type CreateRouteRequest struct {
	Payload *CreateRouteRequestPayload `json:"payload" in:"body=json"`
}

type CreateCompanionRequestPayload struct {
	ClientID string `json:"client_id" bson:"_id"`
}
type CreateCompanionRequest struct {
	Payload *CreateCompanionRequestPayload `json:"payload" in:"body=json"`
}

type DeleteRouteRequest struct {
}

type DeleteCompanionRequest struct {
}
