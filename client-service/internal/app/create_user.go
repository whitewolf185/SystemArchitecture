package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/app/mapping"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/gen/client_service/public/model"
	customerrors "github.com/whitewolf185/SystemArchitecture/client-service/pkg/custom_errors"
)

// @Tags client_service
// @Description создание пользователя. Если пользователь с таким username уже существует, то будет выдана ошибка
// @ID client_service_create_user
// @Accept json
// @Produce json
// @Param input body domain.CreateUserPayload true "Данные для создания пользователя. Поля username и password не могут быть пустыми."
// @Success 200 {object} domain.Person
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/CreateUser [post]
func (i Implementation) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.Person, error) {
	if req == nil || req.Payload == nil || req.Payload.UserName == "" || req.Payload.Password == "" {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("username or password cannot be empty"))
	}

	createdPerson, err := i.personRepo.CreateUser(ctx, model.Persons{
		ID:           uuid.New(),
		Username:     req.Payload.UserName,
		UserPassword: req.Payload.Password,
		IsDriver:     req.Payload.IsDriver,
	})
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrUserAlreadyExists):
			return nil, customerrors.CodesBadRequest(err)
		}
		return nil, fmt.Errorf("cannot create user: %w", err)
	}

	return mapping.PersonModelToResponse(createdPerson), nil
}
