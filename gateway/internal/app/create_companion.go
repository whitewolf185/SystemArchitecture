package app

import (
	"github.com/whitewolf185/SystemArchitecture/gateway/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

// @Tags domain_service
// @Description создание попутчика. У пользователя не получится создать несколько сущностей попутчиков
// @ID domain_service_create_companion
// @Accept json
// @Produce json
// @Param input body domain.CreateCompanionRequestPayloadSwag true "Параметры, по которым создаем информацию"
// @Success 200 {object} domain.Companion
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/CreateCompanion [post]

var (
	_ = domain.CreateCompanionRequestPayloadSwag{}
	_ = domain.Companion{}
	_ = customerrors.ErrCodes{}
)
