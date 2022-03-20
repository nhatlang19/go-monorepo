package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt"

	"github.com/nhatlang19/go-monorepo/pkg/helper"
	"github.com/nhatlang19/go-monorepo/services/user/model"
)

type TokenService interface {
	ValidateIDToken(string) (*model.User, error)
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return tokenService{}
}

func (t tokenService) ValidateIDToken(tokenString string) (*model.User, error) {
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

	fmt.Println("claims", claims)
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}
	fmt.Println("userId", userId)
	user := &model.User{
		ID: userId,
	}

	return user, nil
}
