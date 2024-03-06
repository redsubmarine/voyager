package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/yangoneseok/voyager/db/sqlc"
)

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

var okResponse = gin.H{"result": "ok"}

type loginUserResponse struct {
	SessionID    uuid.UUID `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	User         userResponse
}

type renewAccessTokenResponse struct {
	AccessToken string `json:"access_token" binding:"required"`
}
