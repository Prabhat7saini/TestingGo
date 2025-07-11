package intitalizer

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	authhandler "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/handler"
	"go.uber.org/zap"
)

type BaseHandler struct {
	AuthHandler authhandler.AuthHandlerMethods
}

func NewBaseHandler(logger *zap.Logger, cfg *config.Env, baseService *BaseService) *BaseHandler {
	// authServiceAccess:=authserviceaccess.NewAuthServiceAccess(redis,logger,cfg)

	// NewBaseRepository(authServiceAccess)
	// NewUserRepository(authRepoAccess)
	return &BaseHandler{
		// User: NewUserRepository(access),
		// Add others similarly:
		// Order: NewOrderRepository(access),
		// Auth: NewAuthRepository(access),
		// UserLogin: userloginrepository.NewUserLoginRepository(authRepoAccess),
		// AuthServiceMethods:authservice.NewAuthService(),
		// UserLoginServiceMethods:userloginservice.NewUserLoginService(baseRepo.UserLogin,authServiceAccess),
		AuthHandler: authhandler.NewAuthHandler(baseService.AuthService),
	}
}
