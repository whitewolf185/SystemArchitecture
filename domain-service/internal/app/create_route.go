package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description Создание сущности маршрута у пользователя.
// @ID domain_service_create_route
// @Accept json
// @Produce json
// @Param input body domain.CreateRouteRequest true "Параметры, по которым создаем информацию"
// @Success 200 {object} domain.Route
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/CreateRoute [post]
func (i Implementation) CreateRoute(ctx context.Context, req *domain.CreateRouteRequest) (*domain.Route, error) {
	panic("not implemented")
}
