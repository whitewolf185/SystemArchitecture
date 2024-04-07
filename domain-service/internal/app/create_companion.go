package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description Создание сущности маршрута у пользователя.
// @ID domain_service_create_companion
// @Accept json
// @Produce json
// @Param input body domain.CreateCompanionRequest true "Параметры, по которым создаем информацию"
// @Success 200 {object} domain.Companion
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/CreateCompanion [post]
func (i Implementation) CreateCompanion(ctx context.Context, req *domain.CreateCompanionRequest) (*domain.Companion, error) {
	panic("not implemented")
}
