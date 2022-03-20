package service

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"

	"github.com/nhatlang19/go-monorepo/pkg/helper"
	"github.com/nhatlang19/go-monorepo/pkg/model"
)

type TokenService interface {
	ValidateIDToken(string) (*model.User, error)
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return tokenService{}
}

func (t tokenService) ValidateIDToken(tokenString string) (*helper.UserInfo, error) {
	jwtToken := helper.NewJwtToken()
	token, err := jwtToken.VerifyToken(tokenString)
	if err != nil {
		log.Printf("Unable to validate or parse idToken - Error: %v\n", err)
		return nil, fmt.Errorf("Unable to verify user from idToken")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Printf("Invalid JwtToken - Error: %v\n", err)
		return nil, fmt.Errorf("Invalid JwtToken")
	}

	user := claims["user"].(*helper.UserInfo)

	return user, nil
}
