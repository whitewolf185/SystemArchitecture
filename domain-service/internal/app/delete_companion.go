package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description Удаление сущности попутчика у пользователя.
// @ID domain_service_delete_companion
// @Accept json
// @Produce json
// @Param input query domain.DeleteCompanionRequest true "Параметры, по которым удаляем информацию"
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/DeleteCompanion [delete]
func (i Implementation) DeleteCompanion(ctx context.Context, req *domain.DeleteCompanionRequest) error {
	panic("not implemented")
}
