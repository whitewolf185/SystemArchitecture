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

func (em ErrHandler) handleTypeSwitcher(ctx context.Context, _ *http.Request, handleType domain.HandlerType) (interface{}, error) {
	inputQuery := ctx.Value(httpin.Input)
	switch handleType {
	case domain.GetCompanionInfo:
		if inputQuery == nil {
			return em.handlers.GetCompanionInfo(ctx, nil)
		}
		return em.handlers.GetCompanionInfo(ctx, inputQuery.(*domain.GetCompanionInfoRequest))
	case domain.GetRouteInfo:
		if inputQuery == nil {
			return em.handlers.GetRouteInfo(ctx, nil)
		}
		return em.handlers.GetRouteInfo(ctx, inputQuery.(*domain.GetRouteInfoRequest))
	case domain.CreateRoute:
		if inputQuery == nil {
			return em.handlers.CreateRoute(ctx, nil)
		}
		return em.handlers.CreateRoute(ctx, inputQuery.(*domain.CreateRouteRequest))
	case domain.CreateCompanion:
		if inputQuery == nil {
			return em.handlers.CreateCompanion(ctx, nil)
		}
		return em.handlers.CreateCompanion(ctx, inputQuery.(*domain.CreateCompanionRequest))
	case domain.DeleteRoute:
		if inputQuery == nil {
			err := em.handlers.DeleteRoute(ctx, nil)
			return nil, err
		}
		err := em.handlers.DeleteRoute(ctx, inputQuery.(*domain.DeleteRouteRequest))
		return nil, err
	case domain.DeleteCompanion:
		if inputQuery == nil {
			err := em.handlers.DeleteCompanion(ctx, nil)
			return nil, err
		}
		err := em.handlers.DeleteCompanion(ctx, inputQuery.(*domain.DeleteCompanionRequest))
		return nil, err
	}

	return nil, customerrors.ErrUnknownType
}
