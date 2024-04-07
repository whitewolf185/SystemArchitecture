package mapping

import (
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/domain-service/internal/gen/client_service/public/model"
)

func PersonModelToResponse(modelPerson *model.Persons) *domain.Person {
	if modelPerson == nil {
		return nil
	}
	return &domain.Person{
		ID:       modelPerson.ID.String(),
		UserName: modelPerson.Username,
		IsDriver: modelPerson.IsDriver,
	}
}

func PersonsModelsToResponse(persons []model.Persons) []domain.Person {
	if len(persons) == 0 {
		return nil
	}
	result := make([]domain.Person, 0, len(persons))
	for _, person := range persons {
		resultPerson := PersonModelToResponse(&person)
		result = append(result, *resultPerson)
	}

	return result
}
