package authhandler

import authusecase "github.com/Squid3rd/FirstGoProject/modules/auth/authUsecase"

type (
	AuthGrpcHandler struct {
		authusecase authusecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authusecase authusecase.AuthUsecaseService) authusecase.AuthUsecaseService {
	return &AuthGrpcHandler{authusecase}
}
