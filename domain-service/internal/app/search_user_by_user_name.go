package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/app/mapping"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

// @Tags client_service
// @Description удаление пользователя пользователя
// @ID client_service_swarch_user_by_user_name
// @Accept json
// @Produce json
// @Param input query domain.SearchUserByUserNameRequest true "Запрос на поиск пользователя по его нику"
// @Success 200 {array} domain.Person
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/SearchUserByUserName [get]
func (i Implementation) SearchUserByUserName(ctx context.Context, req *domain.SearchUserByUserNameRequest) ([]domain.Person, error) {
	if req == nil || req.UserNameIn == "" {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("cannot select persons with empty username"))
	}

	persons, err := i.personRepo.GetUserByUserNameIn(ctx, req.UserNameIn)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrUserNotFound):
			return nil, customerrors.CodesNotFound(err)
		}
		return nil, fmt.Errorf("cannot find users: %w", err)
	}

	return mapping.PersonsModelsToResponse(persons), nil
}
