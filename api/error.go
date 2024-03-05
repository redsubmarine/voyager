package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	emptyAuthorizationError   = errors.New("authorization header is not provided")
	invalidAuthorizationError = errors.New("invalid authorization header format")
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func unsupportedAuthorizationError(aType string) error {
	return fmt.Errorf("unsupported authorization type %s", aType)
}
