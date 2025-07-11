package authrepositoryaccess


import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"

	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


type AuthRepositoryAccess struct {
	DB     *gorm.DB
	Redis  sharedRedis.Client
	Logger *zap.Logger
	Config *config.Env
}


func NewAuthRepoAccess(db *gorm.DB, redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env) *AuthRepositoryAccess {
	return &AuthRepositoryAccess{
		DB:     db,
		Redis:  redis,
		Logger: logger,
		Config: cfg,
	}
}
