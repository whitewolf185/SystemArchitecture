package app

import (
	"context"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags domain_service
// @Description создание попутчика. У пользователя не получится создать несколько сущностей попутчиков
// @ID domain_service_create_companion
// @Accept json
// @Produce json
// @Param input body domain.CreateCompanionRequestPayload true "Параметры, по которым создаем информацию"
// @Success 200 {object} domain.Companion
// @Failure default  {object}  customerrors.ErrCodes
// @Router /domain/CreateCompanion [post]
func (i Implementation) CreateCompanion(ctx context.Context, req *domain.CreateCompanionRequest) (*domain.Companion, error) {
	if req == nil || req.Payload == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("request cannot be empty"))
	}

	if err := i.repo.InsertCompanion(ctx, *req.Payload); err != nil {
		return nil, fmt.Errorf("cannot insert: %w", err)
	}

	return &domain.Companion{
		ClientID:    req.Payload.ClientID,
		Destination: req.Payload.Destination,
	}, nil
}
