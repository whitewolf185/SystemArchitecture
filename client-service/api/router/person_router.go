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
	personRouter.With(httpin.NewInput(domain.GetPersonByIDRequest{})).Get("/"+domain.GetClientByID.String(),
		errHandler.ErrMiddleware(domain.GetClientByID),
	)
	personRouter.With(httpin.NewInput(domain.SearchUserByUserNameRequest{})).Get("/"+domain.SearchUserByUserName.String(),
		errHandler.ErrMiddleware(domain.SearchUserByUserName),
	)
	// POST
	personRouter.With(httpin.NewInput(domain.CreateUserRequest{})).Post("/"+domain.CreateUser.String(),
		errHandler.ErrMiddleware(domain.CreateUser),
	)
	personRouter.With(httpin.NewInput(domain.LoginRequest{})).Post("/"+domain.Login.String(),
		errHandler.ErrMiddleware(domain.Login),
	)
	// DELETE
	personRouter.With(httpin.NewInput(domain.DeleteUserByIDRequest{})).Delete("/"+domain.DeleteUserByID.String(),
		errHandler.ErrMiddleware(domain.DeleteUserByID),
	)

	return personRouter
}
