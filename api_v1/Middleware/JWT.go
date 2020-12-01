package Middleware

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// jwt checker
func JWTChecker() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(Config.ACCESS_TOKEN_SECRETKEY)
			return b, nil
		})

		// トークンが有効
		if err == nil {
			err := Models.BlackListChecker(token.Raw)
			if err == nil {
				context.JSON(http.StatusBadRequest, gin.H{"err": "トークンは無効です。"})
				context.Abort()
			} else {
				claims := token.Claims.(jwt.MapClaims)
				context.Set("exp", claims["exp"])
				context.Set("accessToken", token.Raw)
				context.Next()
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"err":err})
			context.Abort()
		}
	}
}

// Refresh Token 限定ミドルウェア
func RefreshChecker() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(Config.REFRESH_TOKEN_SECRETKEY)
			return b, nil
		})

		if err == nil {
			err := Models.BlackListChecker(token.Raw)
			if err == nil {
				context.JSON(http.StatusBadRequest, gin.H{"err": err})
				context.Abort()
			} else {
				claims := token.Claims.(jwt.MapClaims)
				context.Set("userId", claims["userId"])
				context.Set("exp", claims["exp"])
				context.Set("refreshToken", token.Raw)
				context.Next()
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"err":err})
			context.Abort()
		}
	}
}