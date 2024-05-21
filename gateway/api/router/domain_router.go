package router

import (
	domain_client_service "github.com/whitewolf185/SystemArchitecture/client-service/api/domain"
	domain_domain "github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/gateway/api/middleware"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
)

func newDomainRouter(errHandler middleware.ErrHandler) chi.Router {
	domainRouter := chi.NewRouter()
	// GET
	domainRouter.With(httpin.NewInput(domain_domain.GetCompanionInfoRequest{})).Get("/"+domain_domain.GetCompanionInfo.String(),
		errHandler.ErrMiddleware(domain_domain.GetCompanionInfo),
	)
	domainRouter.With(httpin.NewInput(domain_domain.GetRouteInfoRequest{})).Get("/"+domain_domain.GetRouteInfo.String(),
		errHandler.ErrMiddleware(domain_domain.GetRouteInfo),
	)
	// POST
	domainRouter.With(httpin.NewInput(domain_domain.CreateRouteRequest{})).Post("/"+domain_domain.CreateRoute.String(),
		errHandler.ErrMiddleware(domain_domain.CreateRoute),
	)
	domainRouter.With(httpin.NewInput(domain_domain.CreateCompanionRequest{})).Post("/"+domain_domain.CreateCompanion.String(),
		errHandler.ErrMiddleware(domain_domain.CreateCompanion),
	)
	domainRouter.With(httpin.NewInput(domain_client_service.LoginRequest{})).Post("/"+domain_client_service.Login.String(),
		errHandler.ErrMiddleware(domain_client_service.Login),
	)
	// DELETE
	domainRouter.With(httpin.NewInput(domain_domain.DeleteRouteRequest{})).Delete("/"+domain_domain.DeleteRoute.String(),
		errHandler.ErrMiddleware(domain_domain.DeleteRoute),
	)
	domainRouter.With(httpin.NewInput(domain_domain.DeleteCompanionRequest{})).Delete("/"+domain_domain.DeleteCompanion.String(),
		errHandler.ErrMiddleware(domain_domain.DeleteCompanion),
	)
	return domainRouter
}
