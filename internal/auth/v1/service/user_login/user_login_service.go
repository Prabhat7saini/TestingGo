package userloginservice

import (
	"context"
	// "net/http"

	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"
	userloginrepository "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/repository/user_login"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/constants/exception"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/utils"
	// "go.uber.org/zap"

	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/dto"
	authserviceaccess "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/v1/service"
	// "gorm.io/gorm"
)

type UserLoginServiceMethods interface {
	GetUserByEmail(ctx context.Context, email string, password string) (IGetUserByEmailResponse, error)
	// Login(ctx context.Context, userLogin *models.UserLogin)
}

type userLoginService struct {
	repo   userloginrepository.UserLoginRepositoryMethods
	access *authserviceaccess.AuthServiceAccess
}

func NewUserLoginService(
	repo userloginrepository.UserLoginRepositoryMethods,
	access *authserviceaccess.AuthServiceAccess,
) UserLoginServiceMethods {
	return &userLoginService{
		repo:   repo,
		access: access,
	}
}

func (s *userLoginService) GetUserByEmail(
	ctx context.Context,
	email string,
	password string,
) (IGetUserByEmailResponse, error) {

	user, err := s.repo.FindUserByFields(ctx, map[string]interface{}{"email": email}, "id", "role_id","password_hash")
	if err != nil {
		return IGetUserByEmailResponse{} ,err
	}
	return IGetUserByEmailResponse{Email: user.Email,Password: user.PasswordHash},nil
}
