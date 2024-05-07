package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/config"
	"github.com/whitewolf185/SystemArchitecture/client-service/internal/gen/client_service/public/model"
	customerrors "github.com/whitewolf185/SystemArchitecture/client-service/pkg/custom_errors"
)

// @Tags client_service
// @Description Login пользователя
// @ID client_service_login
// @Accept json
// @Produce json
// @Param input body domain.LoginPayload true "Данные пользователя. Поля username и password не могут быть пустыми."
// @Failure default  {object}  customerrors.ErrCodes
// @Router /person/Login [post]
func (i Implementation) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	if req == nil || req.Payload == nil || req.Payload.UserName == "" || req.Payload.Password == "" {
		return nil, customerrors.CodesBadRequest(fmt.Errorf("username or password cannot be empty"))
	}

	person, err := i.personRepo.GetUserByName(ctx, model.Persons{
		Username: req.Payload.UserName,
	})
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrCannotLoginUser):
			return nil, customerrors.CodesUnauthorized(err)
		}
		return nil, fmt.Errorf("cannot login user: %w", err)
	}

	if person.UserPassword != req.Payload.Password {
		return nil, customerrors.CodesUnauthorized(fmt.Errorf("passwords not equal"))
	}

	expireTime := time.Hour * 24
	tokenString, err := createJWT(person, expireTime)
	if err != nil {
		return nil, fmt.Errorf("cannot generate jwt: %w", err)
	}

	result := domain.LoginResponse{
		Token:  tokenString,
		MaxAge: int(expireTime / time.Second),
	}
	return &result, nil
}

type jwtClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func createJWT(person *model.Persons, expireTime time.Duration) (string, error) {
	claims := jwtClaims{
		ID: person.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetValue(config.JwtSecret)))
	if err != nil {
		return "", fmt.Errorf("signing jwt failure: %w", err)
	}

	return tokenString, nil
}
