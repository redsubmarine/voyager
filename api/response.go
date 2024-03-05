package api

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/yangoneseok/voyager/db/sqlc"
)

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

var okResponse = gin.H{"result": "ok"}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	User        userResponse
}
