package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description создание маршрута. Водителю не получится создать несколько маршрутов. Сущность маршрутов будет одна, но в нее можно будет поместить несколько пунктов назначения.
// @ID domain_service_get_route_info
// @Accept json
// @Produce json
// @Param input query domain.GetRouteInfoRequest true "Параметры, по которым получаем информацию"
// @Success 200 {object} domain.Route
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/GetRouteInfo [get]
func (i Implementation) GetRouteInfo(ctx context.Context, req *domain.GetRouteInfoRequest) (*domain.Route, error) {
	panic("not implemented")
}
