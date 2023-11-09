package services

import "github.com/dgrijalva/jwt-go"

//办法token逻辑

type jwtService struct {
}

var JwtService = new(jwtService)

//所有需要颁发token的用户模型必须实现这个接口

type jwtUser interface {
	GetUid() string
}

//CustomClaims 自定义 Claims

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

//CreateToken 生成Token

func (jwtService *jwtService) CreateToken(GuardName string, user jwtUser) (tokenData TokenOutPut, err error) {

}
