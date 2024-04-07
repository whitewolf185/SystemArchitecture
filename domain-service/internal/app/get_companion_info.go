package app

import (
	"context"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
)

// @Tags domain_service
// @Description создание попутчика. У пользователя не получится создать несколько сущностей попутчиков
// @ID domain_service_get_companion_info
// @Accept json
// @Produce json
// @Param input query domain.GetCompanionInfoRequest true "Параметры, по которым получаем информацию"
// @Success 200 {object} domain.Companion
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/GetCompanionInfo [get]
func (i Implementation) GetCompanionInfo(ctx context.Context, req *domain.GetCompanionInfoRequest) (*domain.Companion, error) {
	panic("not implemented")
}
