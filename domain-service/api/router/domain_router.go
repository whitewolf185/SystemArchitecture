package router

import (
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/domain"
	"github.com/whitewolf185/SystemArchitecture/domain-service/api/middleware"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
)

func newDomainRouter(errHandler middleware.ErrHandler) chi.Router {
	domainRouter := chi.NewRouter()
	// GET
	domainRouter.With(httpin.NewInput(domain.GetCompanionInfoRequest{})).Get("/"+domain.GetCompanionInfo.String(),
		errHandler.ErrMiddleware(domain.GetCompanionInfo),
	)
	domainRouter.With(httpin.NewInput(domain.GetRouteInfoRequest{})).Get("/"+domain.GetRouteInfo.String(),
		errHandler.ErrMiddleware(domain.GetRouteInfo),
	)
	// POST
	domainRouter.With(httpin.NewInput(domain.CreateRouteRequest{})).Post("/"+domain.CreateRoute.String(),
		errHandler.ErrMiddleware(domain.CreateRoute),
	)
	domainRouter.With(httpin.NewInput(domain.CreateCompanionRequest{})).Post("/"+domain.CreateCompanion.String(),
		errHandler.ErrMiddleware(domain.CreateCompanion),
	)
	// DELETE
	domainRouter.With(httpin.NewInput(domain.DeleteRouteRequest{})).Delete("/"+domain.DeleteRoute.String(),
		errHandler.ErrMiddleware(domain.DeleteRoute),
	)
	domainRouter.With(httpin.NewInput(domain.DeleteCompanionRequest{})).Delete("/"+domain.DeleteCompanion.String(),
		errHandler.ErrMiddleware(domain.DeleteCompanion),
	)
	return domainRouter
}
