package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/app/mapping"
	customerrors "github.com/whitewolf185/SystemArchitecture/client-service/pkg/custom_errors"
)

// @Tags client_service
// @Description удаление пользователя пользователя
// @ID client_service_delete_user
// @Accept json
// @Produce json
// @Param input query domain.DeleteUserByIDRequest true "Запрос на удаление пользователя"
// @Success 200 {object} domain.Person
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/DeleteUserByID [delete]
func (i Implementation) DeleteUserByID(ctx context.Context, req *domain.DeleteUserByIDRequest) (*domain.Person, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}
	clientID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("clientID must be uuid: %w", err))
	}

	deletedPerson, err := i.personRepo.DeleteUserByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("cannot delete user: %w", err)
	}
	if deletedPerson == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("user does not exists, nothing to delete"))
	}

	return mapping.PersonModelToResponse(deletedPerson), nil
}
