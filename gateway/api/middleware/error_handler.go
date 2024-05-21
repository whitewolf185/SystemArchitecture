package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ggicci/httpin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	domain_client_service "github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/internal/config"
	customerrors "github.com/whitewolf185/SystemArchitecture/gateway/pkg/custom_errors"
)

const contextDeadline = time.Minute * 5

// ErrHandler структура используется для того, чтобы возможно было хендлить ошибки из ручек
type ErrHandler struct {
	personHandlers domain_client_service.Handlers
	domainHandlers domain_domain.Handlers
}

// Конструктор для ErrHandler
func NewErrorHandler(personHandlers domain_client_service.Handlers, domainHandlers domain_domain.Handlers) ErrHandler {
	return ErrHandler{
		personHandlers: personHandlers,
		domainHandlers: domainHandlers,
	}
}

// ErrMiddleware - функция-хендлер. Принимает в себя тип ручки, которая используется в хендлере
func (em ErrHandler) ErrMiddleware(handleType fmt.Stringer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Method", handleType.String())
		logrus.Infof("method: %v", handleType)
		ctx, cancel := context.WithTimeout(r.Context(), contextDeadline)
		defer cancel()

		res, err := em.handleTypeSwitcher(ctx, r, handleType)
		if err != nil {
			errToReturn := customerrors.ErrCodes{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
			switch v := err.(type) {
			case customerrors.ErrCodes:
				errToReturn.Code = v.Code
				errToReturn.Err = v.Err
				w.WriteHeader(v.Code)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			err = errors.Wrapf(err, "Method %s ->", handleType.String())
			logrus.Errorln(err.Error())

			errResutl, errMarshal := json.Marshal(errToReturn)
			if errMarshal != nil {
				logrus.Errorln(errMarshal)
			}
			w.Write(errResutl)

			return
		}
		em.checkExclusiveFiles(res, w, handleType)
	}
}

// checkExclusiveFiles функция делает специфическую отправку файлов (не json), если того требует логика
func (em ErrHandler) checkExclusiveFiles(res interface{}, w http.ResponseWriter, handleType fmt.Stringer) {
	if res == nil {
		w.Write(nil)
		w.WriteHeader(http.StatusOK)
		return
	}

	switch v := res.(type) {
	case *domain_client_service.LoginResponse:
		cookie := http.Cookie{
			Name:   "auth",
			Value:  v.Token,
			Path:   "/",
			MaxAge: v.MaxAge,
		}
		http.SetCookie(w, &cookie)
	default:
		toSend, err := json.Marshal(res)
		if err != nil {
			err = errors.Wrapf(err, "Method %s ->", handleType.String())
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Errorln("unmarshal response error ", err.Error())
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(toSend)
	}
	w.WriteHeader(http.StatusOK)
}

type jwtClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func (em ErrHandler) handleTypeSwitcher(ctx context.Context, req *http.Request, handleType fmt.Stringer) (interface{}, error) {
	inputQuery := ctx.Value(httpin.Input)
	switch handleTypeCast := handleType.(type) {
	case domain_domain.HandlerType:
		// preparing coockie check
		c, err := req.Cookie("auth")
		if err != nil {
			return nil, customerrors.CodesUnauthorized(fmt.Errorf("cookie err: %w", err))
		}
		var claims jwtClaims
		token, err := jwt.ParseWithClaims(c.Value, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetValue(config.JwtSecret)), nil
		})
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrSignatureInvalid):
				return nil, customerrors.CodesUnauthorized(fmt.Errorf("wrong signature: %w", err))
			}
			return nil, customerrors.CodesBadRequest(fmt.Errorf("empty cookie"))
		}
		if !token.Valid {
			return nil, customerrors.CodesUnauthorized(fmt.Errorf("token is not valid"))
		}

		switch handleTypeCast {
		case domain_domain.GetCompanionInfo:
			if inputQuery == nil {
				return em.domainHandlers.GetCompanionInfo(ctx, nil)
			}
			toInput := inputQuery.(*domain_domain.GetCompanionInfoRequest)
			toInput.ClientID = claims.ID
			return em.domainHandlers.GetCompanionInfo(ctx, toInput)
		case domain_domain.GetRouteInfo:
			if inputQuery == nil {
				return em.domainHandlers.GetRouteInfo(ctx, nil)
			}
			toInput := inputQuery.(*domain_domain.GetRouteInfoRequest)
			toInput.ClientID = claims.ID
			return em.domainHandlers.GetRouteInfo(ctx, toInput)
		case domain_domain.CreateRoute:
			if inputQuery == nil {
				return em.domainHandlers.CreateRoute(ctx, nil)
			}
			toInput := inputQuery.(*domain_domain.CreateRouteRequest)
			toInput.Payload.ClientID = claims.ID
			return em.domainHandlers.CreateRoute(ctx, toInput)
		case domain_domain.CreateCompanion:
			if inputQuery == nil {
				return em.domainHandlers.CreateCompanion(ctx, nil)
			}
			toInput := inputQuery.(*domain_domain.CreateCompanionRequest)
			toInput.Payload.ClientID = claims.ID
			return em.domainHandlers.CreateCompanion(ctx, toInput)
		case domain_domain.DeleteRoute:
			if inputQuery == nil {
				err := em.domainHandlers.DeleteRoute(ctx, nil)
				return nil, err
			}
			err := em.domainHandlers.DeleteRoute(ctx, inputQuery.(*domain_domain.DeleteRouteRequest))
			return nil, err
		case domain_domain.DeleteCompanion:
			if inputQuery == nil {
				err := em.domainHandlers.DeleteCompanion(ctx, nil)
				return nil, err
			}
			err := em.domainHandlers.DeleteCompanion(ctx, inputQuery.(*domain_domain.DeleteCompanionRequest))
			return nil, err
		}

	case domain_client_service.HandlerType:
		switch handleTypeCast {
		case domain_client_service.Login:
			if inputQuery == nil {
				return em.personHandlers.Login(ctx, nil)
			}
			return em.personHandlers.Login(ctx, inputQuery.(*domain_client_service.LoginRequest))
		}

	}

	return nil, customerrors.ErrUnknownType
}
