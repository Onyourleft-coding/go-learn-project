package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"todo_list/app/common/response"
	"todo_list/app/services"
	"todo_list/global"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		//	Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		//	Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		fmt.Println("claims", claims)
		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}