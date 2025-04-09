package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hse-revizor/auth-service/internal/pkg/models"
	"github.com/hse-revizor/auth-service/internal/pkg/service/auth"
	"github.com/hse-revizor/auth-service/internal/utils/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type AuthHandler struct {
	cfg          *config.Config
	oauth2Config *oauth2.Config
	authService  *auth.Service
}

func NewAuthHandler(cfg *config.Config, authService *auth.Service) *AuthHandler {
	oauth2Config := &oauth2.Config{
		ClientID:     cfg.GitHub.ClientID,
		ClientSecret: cfg.GitHub.ClientSecret,
		RedirectURL:  cfg.GitHub.RedirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	return &AuthHandler{
		cfg:          cfg,
		oauth2Config: oauth2Config,
		authService:  authService,
	}
}

// @Summary     Домашняя страница
// @Description Отображает страницу с кнопкой для входа через GitHub
// @Tags        auth
// @Accept      json
// @Produce     html
// @Success     200 {object} gin.H
// @Router      / [get]
func (h *AuthHandler) HandleHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"LoginURL": "/api/v1/login",
	})
}

// @Summary     Вход через GitHub
// @Description Перенаправляет пользователя на страницу авторизации GitHub
// @Tags        auth
// @Accept      json
// @Produce     json
// @Success     307 {string} string "Redirect to GitHub"
// @Router      /login [get]
func (h *AuthHandler) HandleLogin(c *gin.Context) {
	url := h.oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// @Summary     Callback от GitHub
// @Description Обрабатывает ответ от GitHub после успешной авторизации
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       code query string true "Код авторизации от GitHub"
// @Success     200 {object} models.GitHubUser
// @Failure     500 {object} gin.H
// @Router      /auth/github/callback [get]
func (h *AuthHandler) HandleCallback(c *gin.Context) {
	ctx := context.Background()
	code := c.Query("code")

	token, err := h.oauth2Config.Exchange(ctx, code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	client := h.oauth2Config.Client(ctx, token)
	githubUser, err := h.fetchGitHubUser(client)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Проверяем, существует ли пользователь
	user, err := h.authService.GetUserByGitHubID(ctx, githubUser.GitHubID)
	if err != nil {
		if err == auth.ErrUserNotFound {
			// Создаем нового пользователя
			user, err = h.authService.CreateUser(ctx, githubUser)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"user":    user,
	})
}

func (h *AuthHandler) fetchGitHubUser(client *http.Client) (*models.GitHubUser, error) {
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	type GitHubResponse struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		Company   string `json:"company"`
		Location  string `json:"location"`
		Bio       string `json:"bio"`
		AvatarURL string `json:"avatar_url"`
	}

	var ghResp GitHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&ghResp); err != nil {
		return nil, err
	}

	// Преобразуем в нашу модель пользователя
	user := &models.GitHubUser{
		GitHubID:  ghResp.ID,
		Login:     ghResp.Login,
		Email:     &ghResp.Email,
		Name:      &ghResp.Name,
		Company:   &ghResp.Company,
		Location:  &ghResp.Location,
		Bio:       &ghResp.Bio,
		AvatarURL: ghResp.AvatarURL,
	}
	return user, nil
}
