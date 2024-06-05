package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags domain_service
// @Description получение маршрута по фильрам.
// @ID domain_service_get_route_info
// @Accept json
// @Produce json
// @Param input query domain.GetRouteInfoRequestSwag true "Параметры, по которым получаем информацию"
// @Success 200 {array} domain.Route
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/GetRouteInfo [get]
func (i Implementation) GetRouteInfo(ctx context.Context, req *domain.GetRouteInfoRequest) ([]domain.Route, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("request cannot be empty"))
	}

	result, err := i.repo.GetRoutes(ctx, domain.Route{
		// ClientID: req.ClientID,
		Path: []string{req.OneOfPath},
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
