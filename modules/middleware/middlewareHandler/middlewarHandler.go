package middlewarehandler

import (
	"github.com/Squid3rd/FirstGoProject/config"
	middlewareusecase "github.com/Squid3rd/FirstGoProject/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareusecase middlewareusecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareusecase middlewareusecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareusecase: middlewareusecase,
	}
}
