// package service_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/models"
// 	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/auth/service"
// )

// type MockUserRepo struct{ mock.Mock }

// func (m *MockUserRepo) FindByFields(ctx context.Context, conditions map[string]interface{}, selectFields ...string) (*models.UserLogin, error) {
// 	args := m.Called(ctx, conditions)
// 	if u := args.Get(0); u != nil {
// 		return u.(*models.UserLogin), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func TestAuthService_Login_InvalidCases(t *testing.T) {
// 	ctx := context.Background()

// 	t.Run("email not found", func(t *testing.T) {
// 		mockRepo := new(MockUserRepo)
// 		email := "ghost@example.com"
// 		password := "AnyP@ss1"
// 		conditions := map[string]interface{}{"email": email}

// 		mockRepo.On("FindByFields", mock.Anything, conditions).Return(nil, nil)

// 		svc := service.NewAuthService(mockRepo, func(hash, plain string) bool {
// 			return false // won't be called
// 		})

// 		token, err := svc.Login(ctx, email, password)
// 		assert.ErrorIs(t, err, service.ErrInvalidCredentials)
// 		assert.Empty(t, token)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("password invalid", func(t *testing.T) {
// 		mockRepo := new(MockUserRepo)
// 		email := "user@example.com"
// 		password := "wrongpass"
// 		conditions := map[string]interface{}{"email": email}

// 		user := &models.UserLogin{
// 			ID:           1,
// 			UUID:         uuid.New(),
// 			Email:        email,
// 			PasswordHash: "hashedPwd",
// 			IsActive:     true,
// 		}

// 		mockRepo.On("FindByFields", mock.Anything, conditions).Return(user, nil)

// 		svc := service.NewAuthService(mockRepo, func(hash, plain string) bool {
// 			return false // simulate mismatch
// 		})

// 		token, err := svc.Login(ctx, email, password)
// 		assert.ErrorIs(t, err, service.ErrInvalidCredentials)
// 		assert.Empty(t, token)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("valid login", func(t *testing.T) {
// 		mockRepo := new(MockUserRepo)
// 		email := "user@example.com"
// 		password := "correctpass"
// 		conditions := map[string]interface{}{"email": email}

// 		user := &models.UserLogin{
// 			ID:           1,
// 			UUID:         uuid.New(),
// 			Email:        email,
// 			PasswordHash: "hashedPwd",
// 			IsActive:     true,
// 		}

// 		mockRepo.On("FindByFields", mock.Anything, conditions).Return(user, nil)

// 		svc := service.NewAuthService(mockRepo, func(hash, plain string) bool {
// 			return true // simulate match
// 		})

// 		token, err := svc.Login(ctx, email, password)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "jwt-token", token)
// 		mockRepo.AssertExpectations(t)
// 	})
// }
package service