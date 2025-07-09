package repository

import (
	"github.com/redis/go-redis/v9"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// repositoryaccess holds shared dependencies across all repositories
type RepositoryAccess struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
	Config *config.Env
}

// newrepositoryaccess initializes repositoryaccess
func NewRepositoryAccess(db *gorm.DB, redis *redis.Client, logger *zap.Logger, cfg *config.Env) *RepositoryAccess {
	return &RepositoryAccess{
		DB:     db,
		Redis:  redis,
		Logger: logger,
		Config: cfg,
	}
}
