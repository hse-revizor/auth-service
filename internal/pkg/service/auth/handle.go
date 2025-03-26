package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/hse-revizor/auth-service/internal/pkg/models"
	"github.com/hse-revizor/auth-service/internal/pkg/storage/sql"
	"github.com/hse-revizor/auth-service/internal/utils/json"
)

func (s *Service) CreateUser(ctx context.Context, user *models.GitHubUser) (*models.GitHubUser, error) {
	created, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityExists):
			return nil, ErrUserExists
		default:
			return nil, err
		}
	}
	log.Debug(fmt.Sprintf("Created GitHub user: %s", json.ToColorJson(created)))
	return created, nil
}

func (s *Service) GetUserByGitHubID(ctx context.Context, githubID int) (*models.GitHubUser, error) {
	user, err := s.storage.FindUserByGitHubID(ctx, githubID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrUserNotFound
		default:
			return nil, err
		}
	}
	return user, nil
}
