package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/app/mapping"
	customerrors "github.com/whitewolf185/SystemArchitecture/client-service/pkg/custom_errors"
)

// @Tags client_service
// @Description получение пользователя по его id
// @ID client_service_get_client_by_ID
// @Accept json
// @Produce json
// @Param input query domain.GetPersonByIDRequest true "ID клиента"
// @Success 200 {object} domain.Person
// @Failure default  {object}  customerrors.ErrCodes
// @Router /gateway/GetClientByID [get]
func (i Implementation) GetClientByID(ctx context.Context, req *domain.GetPersonByIDRequest) (*domain.Person, error) {
	if req == nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("empty request"))
	}
	clientID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("clientID must be uuid: %w", err))
	}
	person, err := i.personRepo.GetUserByID(ctx, clientID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrUserNotFound):
			return nil, customerrors.CodesNotFound(err)
		}
		return nil, fmt.Errorf("cannot get person: %w", err)
	}
	if person == nil {
		return nil, customerrors.CodesNotFound(fmt.Errorf("cannot find user"))
	}
	return mapping.PersonModelToResponse(person), nil
}
