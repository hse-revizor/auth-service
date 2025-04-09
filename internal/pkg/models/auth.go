package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GitHubUser struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	GitHubID  int       `gorm:"unique;column:github_id" json:"github_id" `
	Login     string    `gorm:"unique;column:login" json:"login"`
	Email     *string   `json:"email"`
	Name      *string   `json:"name"`
	Company   *string   `json:"company"`
	Location  *string   `json:"location"`
	Bio       *string   `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (g *GitHubUser) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}
