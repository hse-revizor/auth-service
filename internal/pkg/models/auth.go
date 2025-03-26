package models

import (
	"time"

	"github.com/google/uuid"
)

// GitHubUser представляет пользователя GitHub в системе
type GitHubUser struct {
	// ID пользователя в нашей системе
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	// ID пользователя в GitHub
	GitHubID int `gorm:"unique" json:"github_id" example:"12345678"`
	// Логин пользователя в GitHub
	Login string `gorm:"unique" json:"login" example:"octocat"`
	// Email пользователя
	Email string `json:"email" example:"octocat@github.com"`
	// Время создания записи
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" example:"2024-03-20T15:04:05Z"`
	// Время последнего обновления
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" example:"2024-03-20T15:04:05Z"`
}
