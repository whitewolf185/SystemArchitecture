package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags domain_service
// @Description Удаление сущности попутчика у пользователя.
// @ID domain_service_delete_companion
// @Accept json
// @Produce json
// @Param input query domain.DeleteCompanionRequest true "Параметры, по которым удаляем информацию"
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/DeleteCompanion [delete]
func (i Implementation) DeleteCompanion(ctx context.Context, req *domain.DeleteCompanionRequest) error {
	if req == nil {
		return customerrors.CodesBadRequest(fmt.Errorf("request cannot be empty"))
	}
	err := i.repo.DeleteCompanion(ctx, req.ClientID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNotFound):
			return customerrors.CodesNotFound(err)
		}

		return err
	}

	return nil
}
