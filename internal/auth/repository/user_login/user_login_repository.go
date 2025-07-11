package userloginrepository

import (
	"context"

	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/repository"
	"gorm.io/gorm"
)

type UserLoginRepositoryMethods interface {
	FindUserByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) (*models.UserLogin, error)
}

type userLoginRepository struct {
	access *authrepositoryaccess.AuthRepositoryAccess
}

func NewUserLoginRepository(access *authrepositoryaccess.AuthRepositoryAccess) UserLoginRepositoryMethods {
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
