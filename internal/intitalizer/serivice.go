package intitalizer

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	authserviceaccess "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service"
	authservice "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service/auth_service"
	userloginservice "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service/user_login"
	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
)

type BaseService struct {
	UserLoginService userloginservice.UserLoginServiceMethods
	AuthService      authservice.AuthServiceMethods
}

func NewBaseService(redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env, baseRepo *BaseRepository) *BaseService {
	authServiceAccess := authserviceaccess.NewAuthServiceAccess(redis, logger, cfg)

	// authServiceAccess := authserviceaccess.NewAuthServiceAccess(redis, logger, cfg)

	userLoginSvc := userloginservice.NewUserLoginService(baseRepo.UserLogin, authServiceAccess)
	authSvc := authservice.NewAuthService(userLoginSvc, authServiceAccess)
	return &BaseService{
		UserLoginService: userLoginSvc,
		AuthService:      authSvc,
	}
}
