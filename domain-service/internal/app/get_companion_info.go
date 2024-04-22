package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags domain_service
// @Description получение попутчика по фильтрам.
// @ID domain_service_get_companion_info
// @Accept json
// @Produce json
// @Param input query domain.GetCompanionInfoRequest true "Параметры, по которым получаем информацию"
// @Success 200 {array} domain.Companion
// @Failure default  {object}  customerrors.ErrCodes
// @Router /domain/GetCompanionInfo [get]
func (i Implementation) GetCompanionInfo(ctx context.Context, req *domain.GetCompanionInfoRequest) ([]domain.Companion, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("request cannot be empty"))
	}

	result, err := i.repo.GetCompanions(ctx, domain.Companion{
		ClientID:    req.ClientID,
		Destination: req.Destination,
	})
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNotFound):
			return nil, customerrors.CodesNotFound(err)
		}

		return nil, err
	}

	return result, nil
}
