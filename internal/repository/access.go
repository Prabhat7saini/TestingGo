package repository

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"

	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// repositoryaccess holds shared dependencies across all repositories
type RepositoryAccess struct {
	DB     *gorm.DB
	Redis  sharedRedis.Client
	Logger *zap.Logger
	Config *config.Env
}

// newrepositoryaccess initializes repositoryaccess
func NewRepositoryAccess(db *gorm.DB, redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env) *RepositoryAccess {
	return &RepositoryAccess{
		DB:     db,
		Redis:  redis,
		Logger: logger,
		Config: cfg,
	}
}
