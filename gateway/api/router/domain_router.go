package router

import (
	domain_client_service "github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/middleware"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
)

func newDomainRouter(errHandler middleware.ErrHandler) chi.Router {
	gatewayRouter := chi.NewRouter()
	// GET
	gatewayRouter.With(httpin.NewInput(domain_domain.GetCompanionInfoRequest{})).Get("/"+domain_domain.GetCompanionInfo.String(),
		errHandler.ErrMiddleware(domain_domain.GetCompanionInfo),
	)
	gatewayRouter.With(httpin.NewInput(domain_domain.GetRouteInfoRequest{})).Get("/"+domain_domain.GetRouteInfo.String(),
		errHandler.ErrMiddleware(domain_domain.GetRouteInfo),
	)
	gatewayRouter.With(httpin.NewInput(domain_client_service.GetPersonByIDRequest{})).Get("/"+domain_client_service.GetClientByID.String(),
		errHandler.ErrMiddleware(domain_client_service.GetClientByID),
	)
	gatewayRouter.With(httpin.NewInput(domain_client_service.SearchUserByUserNameRequest{})).Get("/"+domain_client_service.SearchUserByUserName.String(),
		errHandler.ErrMiddleware(domain_client_service.SearchUserByUserName),
	)
	// POST
	gatewayRouter.With(httpin.NewInput(domain_domain.CreateRouteRequest{})).Post("/"+domain_domain.CreateRoute.String(),
		errHandler.ErrMiddleware(domain_domain.CreateRoute),
	)
	gatewayRouter.With(httpin.NewInput(domain_domain.CreateCompanionRequest{})).Post("/"+domain_domain.CreateCompanion.String(),
		errHandler.ErrMiddleware(domain_domain.CreateCompanion),
	)
	gatewayRouter.With(httpin.NewInput(domain_client_service.LoginRequest{})).Post("/"+domain_client_service.Login.String(),
		errHandler.ErrMiddleware(domain_client_service.Login),
	)
	gatewayRouter.With(httpin.NewInput(domain_client_service.CreateUserRequest{})).Post("/"+domain_client_service.CreateUser.String(),
		errHandler.ErrMiddleware(domain_client_service.CreateUser),
	)
	// DELETE
	gatewayRouter.With(httpin.NewInput(domain_domain.DeleteRouteRequest{})).Delete("/"+domain_domain.DeleteRoute.String(),
		errHandler.ErrMiddleware(domain_domain.DeleteRoute),
	)
	gatewayRouter.With(httpin.NewInput(domain_domain.DeleteCompanionRequest{})).Delete("/"+domain_domain.DeleteCompanion.String(),
		errHandler.ErrMiddleware(domain_domain.DeleteCompanion),
	)
	gatewayRouter.With(httpin.NewInput(domain_client_service.DeleteUserByIDRequest{})).Delete("/"+domain_client_service.DeleteUserByID.String(),
		errHandler.ErrMiddleware(domain_client_service.DeleteUserByID),
	)
	return gatewayRouter
}
