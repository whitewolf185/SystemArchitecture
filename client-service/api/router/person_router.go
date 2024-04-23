package router

import (
	"github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/client-service/api/middleware"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
)

func newPersonRouter(errHandler middleware.ErrHandler) chi.Router {
	personRouter := chi.NewRouter()
	// GET
	personRouter.With(httpin.NewInput(domain.GetPersonByIDRequest{})).Get("/"+string(domain.GetClientByID),
		errHandler.ErrMiddleware(domain.GetClientByID),
	)
	personRouter.With(httpin.NewInput(domain.SearchUserByUserNameRequest{})).Get("/"+string(domain.SearchUserByUserName),
		errHandler.ErrMiddleware(domain.SearchUserByUserName),
	)
	// POST
	personRouter.With(httpin.NewInput(domain.CreateUserRequest{})).Post("/"+string(domain.CreateUser),
		errHandler.ErrMiddleware(domain.CreateUser),
	)
	personRouter.With(httpin.NewInput(domain.LoginRequest{})).Post("/"+string(domain.Login),
		errHandler.ErrMiddleware(domain.Login),
	)
	// DELETE
	personRouter.With(httpin.NewInput(domain.DeleteUserByIDRequest{})).Delete("/"+string(domain.DeleteUserByID),
		errHandler.ErrMiddleware(domain.DeleteUserByID),
	)

	return personRouter
}
