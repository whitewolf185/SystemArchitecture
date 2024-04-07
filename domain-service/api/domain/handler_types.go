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
	GetCompanionInfo(ctx context.Context, req *GetCompanionInfoRequest) ([]Companion, error)
	// GetRouteInfo - получение информации о маршруте
	GetRouteInfo(ctx context.Context, req *GetRouteInfoRequest) ([]Route, error)
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
	ClientID    string `bson:"_id" json:"client_id"`
	Destination string `json:"destination"`
}
type Route struct {
	ClientID string   `bson:"_id" json:"client_id"`
	Path     []string `json:"path"`
}

type GetCompanionInfoRequest struct {
	ClientID    string `in:"query=client_id"`
	Destination string `json:"destination" in:"query=destination"`
}

type GetRouteInfoRequest struct {
	ClientID  string `in:"query=client_id"`
	OneOfPath string `json:"one_of_path" in:"query=one_of_path"`
}

type CreateRouteRequestPayload struct {
	ClientID string   `json:"client_id" bson:"_id"`
	Path     []string `json:"path"`
}
type CreateRouteRequest struct {
	Payload *CreateRouteRequestPayload `json:"payload" in:"body=json"`
}

type CreateCompanionRequestPayload struct {
	ClientID    string `json:"client_id" bson:"_id"`
	Destination string `json:"destination"`
}
type CreateCompanionRequest struct {
	Payload *CreateCompanionRequestPayload `json:"payload" in:"body=json"`
}

type DeleteRouteRequest struct {
	ClientID string `json:"client_id" bson:"_id" in:"query=client_id"`
}

type DeleteCompanionRequest struct {
	ClientID string `json:"client_id" bson:"_id" in:"query=client_id"`
}
