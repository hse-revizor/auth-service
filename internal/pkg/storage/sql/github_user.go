package sql

import (
	"context"

	"github.com/hse-revizor/auth-service/internal/pkg/models"
)

func (s *Storage) CreateUser(ctx context.Context, user *models.GitHubUser) (*models.GitHubUser, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	if err := tr.Create(&user).Error; err != nil {
		if isDuplicateError(err) {
			return nil, ErrEntityExists
		}
		return nil, err
	}

	return user, nil
}

func (s *Storage) FindUserByGitHubID(ctx context.Context, githubID int) (*models.GitHubUser, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)

	var user models.GitHubUser
	if err := tr.Where("github_id = ?", githubID).First(&user).Error; err != nil {
		if isNotFoundError(err) {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}

	return &user, nil
}
