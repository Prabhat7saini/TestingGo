// package user_model
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLogin struct {
	ID           int            `gorm:"primaryKey;column:id"`
	UUID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Name         *string        `gorm:"type:varchar(100)"`
	Email        string         `gorm:"type:varchar(255);not null"`
	PasswordHash string         `gorm:"type:text;not null"`
	RoleID       *int           `gorm:"column:role_id"`
	IsActive     bool           `gorm:"column:is_active;not null;default:true"`
	IsBlocked    bool           `gorm:"column:is_blocked;not null;default:false"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:current_timestamp"`
	CreatedBy    *int           `gorm:"column:created_by"`
	ModifiedAt   *time.Time     `gorm:"column:modified_at"`
	ModifiedBy   *int           `gorm:"column:modified_by"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy    *int           `gorm:"column:deleted_by"`
}

// TableName sets the actual PostgreSQL table name with schema
func (UserLogin) TableName() string {
	return "auth.userlogin"
}
