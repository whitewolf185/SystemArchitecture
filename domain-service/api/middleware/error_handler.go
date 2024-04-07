package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ggicci/httpin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	customerrors "github.com/whitewolf185/SystemArchitecture/domain-service/pkg/custom_errors"
)

const contextDeadline = time.Minute * 5

// ErrHandler структура используется для того, чтобы возможно было хендлить ошибки из ручек
type ErrHandler struct {
	handlers domain.Handlers
}

// Конструктор для ErrHandler
func NewErrorHandler(handlers domain.Handlers) ErrHandler {
	return ErrHandler{
		handlers: handlers,
	}
}

// ErrMiddleware - функция-хендлер. Принимает в себя тип ручки, которая используется в хендлере
func (em ErrHandler) ErrMiddleware(handleType domain.HandlerType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Method", string(handleType))
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
			err = errors.Wrapf(err, "Method %s ->", string(handleType))
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
func (em ErrHandler) checkExclusiveFiles(res interface{}, w http.ResponseWriter, handleType domain.HandlerType) {
	if res == nil {
		w.Write(nil)
		w.WriteHeader(http.StatusOK)
		return
	}

	switch res.(type) {
	default:
		toSend, err := json.Marshal(res)
		if err != nil {
			err = errors.Wrapf(err, "Method %s ->", string(handleType))
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Errorln("unmarshal response error ", err.Error())
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(toSend)
	}
	w.WriteHeader(http.StatusOK)
}

func (em ErrHandler) handleTypeSwitcher(ctx context.Context, _ *http.Request, handleType domain.HandlerType) (interface{}, error) {
	inputQuery := ctx.Value(httpin.Input)
	switch handleType {
	case domain.GetClientByID:
		if inputQuery == nil {
			return em.handlers.GetClientByID(ctx, nil)
		}
		return em.handlers.GetClientByID(ctx, inputQuery.(*domain.GetPersonByIDRequest))
	case domain.CreateUser:
		if inputQuery == nil {
			return em.handlers.CreateUser(ctx, nil)
		}
		return em.handlers.CreateUser(ctx, inputQuery.(*domain.CreateUserRequest))
	case domain.SearchUserByUserName:
		if inputQuery == nil {
			return em.handlers.SearchUserByUserName(ctx, nil)
		}
		return em.handlers.SearchUserByUserName(ctx, inputQuery.(*domain.SearchUserByUserNameRequest))
	case domain.DeleteUserByID:
		if inputQuery == nil {
			return em.handlers.DeleteUserByID(ctx, nil)
		}
		return em.handlers.DeleteUserByID(ctx, inputQuery.(*domain.DeleteUserByIDRequest))
	}
	return nil, customerrors.ErrUnknownType
}
