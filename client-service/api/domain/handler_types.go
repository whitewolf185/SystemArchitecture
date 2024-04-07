package domain

import (
	"context"
)

type HandlerType string

const (
	GetClientByID        = HandlerType("GetClientByID")
	CreateUser           = HandlerType("CreateUser")
	SearchUserByUserName = HandlerType("SearchUserByUserName")
	DeleteUserByID       = HandlerType("DeleteUserByID")
)

// Интерфейс со всеми ручками сервиса
type Handlers interface {
	// GetClientByID - получение пользователя по его id
	GetClientByID(ctx context.Context, req *GetPersonByIDRequest) (*Person, error)
	// CreateUser - создание пользователя с указанием его пароля и логина
	CreateUser(ctx context.Context, req *CreateUserRequest) (*Person, error)
	// SearchUserByUserName - поиск пользователя по substring его userName. В итоге появляется массив
	SearchUserByUserName(ctx context.Context, req *SearchUserByUserNameRequest) ([]Person, error)
	// DeleteUserByID - удаляет пользователя по его id
	DeleteUserByID(ctx context.Context, req *DeleteUserByIDRequest) (*Person, error)
}

type GetPersonByIDRequest struct {
	ID string `in:"query=id"`
}
type Person struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	IsDriver bool   `json:"driver"`
}

type CreateUserPayload struct {
	UserName string `json:"user_name"`
	IsDriver bool   `json:"is_driver"`
	Password string `json:"password"`
}
type CreateUserRequest struct {
	Payload *CreateUserPayload `json:"payload" in:"body=json"`
}

type SearchUserByUserNameRequest struct {
	UserNameIn string `in:"query=user_name_in"`
}

type DeleteUserByIDRequest struct {
	ID string `in:"query=id"`
}
