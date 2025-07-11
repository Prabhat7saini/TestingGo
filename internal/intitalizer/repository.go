package intitalizer
// package repository

import (
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository"
	userloginrepository "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository/user_login"
	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseRepository struct {
	// User UserRepositoryMethods
	// Add other repositories like:
	// Order OrderRepositoryMethods
	// Auth  AuthRepositoryMethods
	// AuthRepo auth.AuthRepo
	// OrderLogin OrderRepositoryMethods
	// AuthLogin AuthRepositoryMethods
	UserLogin userloginrepository.UserLoginRepositoryMethods
}

// NewBaseRepository wires all individual repositories using shared access
func NewBaseRepository(db *gorm.DB, redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env) *BaseRepository {
	authRepoAccess := authrepositoryaccess.NewAuthRepoAccess(db, redis, logger, cfg)
	return &BaseRepository{
		// User: NewUserRepository(access),
		// Add others similarly:
		// Order: NewOrderRepository(access),
		// Auth: NewAuthRepository(access),
		UserLogin: userloginrepository.NewUserLoginRepository(authRepoAccess),
	}
}
