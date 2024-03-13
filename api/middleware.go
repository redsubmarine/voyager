package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(emptyAuthorizationError))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(invalidAuthorizationError))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		switch authorizationType {
		case authorizationTypeBearer:
			accessToken := fields[1]

			payload, err := s.tokenMaker.VerifyToken(accessToken)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			ctx.Set(authorizationPayloadKey, payload)
			ctx.Next()
		default:
			err := unsupportedAuthorizationError(authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		}
	}
}
