package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	OAuth interface{}
	JWT   interface {
		GenerateToken(claims jwt.Claims) (string, error)
		ValidateToken(token string) (*jwt.Token, error)
		ParseJWTToken(tokenString string) (*oauth2.Token, error)
	}
}

func NewAuthenticator(config *oauth2.Config, state string, secret, aud, iss string) Authenticator {
	return Authenticator{
		OAuth: NewOAuthAuthenticator(config, state),
		JWT:   NewJWTAuthenticator(secret, aud, iss),
	}
}
