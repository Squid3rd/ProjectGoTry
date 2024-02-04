package authusecase

import authrepository "github.com/Squid3rd/FirstGoProject/modules/auth/authRepository"

type (
	AuthUsecaseService interface {
	}

	authUsecase struct {
		authRepository authrepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authrepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}
