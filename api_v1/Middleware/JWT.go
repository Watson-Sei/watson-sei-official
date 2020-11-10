package Middleware

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// jwt checker
func JWTChecker() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(Config.SECRETKEY)
			return b, nil
		})

		if err == nil {
			claims := token.Claims.(jwt.MapClaims)
			context.Set("username", claims["username"])
			context.Set("exp", claims["exp"])
			context.Set("token", token.Raw)
			context.Next()
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"err":err})
			context.Abort()
		}
	}
}