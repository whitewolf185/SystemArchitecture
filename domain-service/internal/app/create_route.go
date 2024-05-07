package app

import (
	"context"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags domain_service
// @Description создание маршрута. Водителю не получится создать несколько маршрутов. Сущность маршрутов будет одна, но в нее можно будет поместить несколько пунктов назначения.
// @ID domain_service_create_route
// @Accept json
// @Produce json
// @Param input body domain.CreateRouteRequestPayloadSwag true "Параметры, по которым создаем информацию"
// @Success 200 {object} domain.Route
// @Failure default  {object}  customerrors.ErrCodes
// @Router /domain/CreateRoute [post]
func (i Implementation) CreateRoute(ctx context.Context, req *domain.CreateRouteRequest) (*domain.Route, error) {
	if req == nil || req.Payload == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("request cannot be empty"))
	}

	if err := i.repo.InsertRoute(ctx, *req.Payload); err != nil {
		return nil, fmt.Errorf("cannot insert route: %w", err)
	}

	return &domain.Route{
		ClientID: req.Payload.ClientID,
		Path:     req.Payload.Path,
	}, nil
}
