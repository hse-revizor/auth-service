package auth

import (
	"context"

	"github.com/hse-revizor/auth-service/internal/pkg/models"
)

type storage interface {
	CreateUser(context.Context, *models.GitHubUser) (*models.GitHubUser, error)
	FindUserByGitHubID(context.Context, int) (*models.GitHubUser, error)
}

type Service struct {
	storage storage
}

func New(storage storage) *Service {
	return &Service{storage: storage}
}
