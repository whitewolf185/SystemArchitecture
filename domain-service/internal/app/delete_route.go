package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description Удаление сущности маршрута у пользователя.
// @ID domain_service_delete_route
// @Accept json
// @Produce json
// @Param input query domain.DeleteRouteRequest true "Параметры, по которым удаляем информацию"
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/DeleteRoute [delete]
func (i Implementation) DeleteRoute(ctx context.Context, req *domain.DeleteRouteRequest) error {
	panic("not implemented")
}
