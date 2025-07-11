package authrepositoryaccess

// import userloginrepository "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository/user_login"
import (
	"context"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"

	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	// userloginrepository "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository/user_login"
)

type AuthRepo struct {
	// User UserRepositoryMethods
	// Add other repositories like:
	// Order OrderRepositoryMethods
	// Auth  AuthRepositoryMethods
	// UserLogin userloginrepository.UserLoginRepositoryMethods
	// OrderLogin OrderRepositoryMethods
	// AuthLogin AuthRepositoryMethods
}

// package repository

// repositoryaccess holds shared dependencies across all repositories
type AuthRepositoryAccess struct {
	DB     *gorm.DB
	Redis  sharedRedis.Client
	Logger *zap.Logger
	Config *config.Env
}

// newrepositoryaccess initializes repositoryaccess
func NewAuthRepoAccess(db *gorm.DB, redis sharedRedis.Client, logger *zap.Logger, cfg *config.Env) *AuthRepositoryAccess {
	return &AuthRepositoryAccess{
		DB:     db,
		Redis:  redis,
		Logger: logger,
		Config: cfg,
	}
}

// package userloginrepository

// import (
// 	"context"

// 	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth"
// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository"
// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"
// 	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/repository"
// 	"gorm.io/gorm"
// )

type UserLoginRepositoryMethods interface {
	FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) (*models.UserLogin, error)
}

type userLoginRepository struct {
	access *AuthRepositoryAccess
}

func NewUserLoginRepository(access *AuthRepositoryAccess) UserLoginRepositoryMethods {
	return &userLoginRepository{access: access}
}

func (ul *userLoginRepository) FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) (*models.UserLogin, error) {
	var user models.UserLogin
	db := ul.access.DB

	query := db.WithContext(ctx).Model(&models.UserLogin{})

	// Optional select
	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	// Apply map conditions safely
	query = query.Where(conditions)

	// Execute query
	err := query.First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err // real DB error
	}
	return &user, nil
}
