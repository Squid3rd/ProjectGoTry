package authhandler

import (
	"github.com/Squid3rd/FirstGoProject/config"
	authusecase "github.com/Squid3rd/FirstGoProject/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authusecase authusecase.AuthUsecaseService
	}
)

func NewAuthHandler(cfg *config.Config, authusecase authusecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authusecase}
}
