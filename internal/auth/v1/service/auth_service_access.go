package authserviceaccess

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"go.uber.org/zap"
		sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
)

type AuthServiceAccess struct {
	Redis  sharedRedis.Client
	Logger *zap.Logger
	Config *config.Env
}

func NewAuthServiceAccess(redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env) *AuthServiceAccess {
	return &AuthServiceAccess{
		Redis:  redis,
		Logger: logger,
		Config: cfg,
	}
}
