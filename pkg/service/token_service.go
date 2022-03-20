package service

import (
	"fmt"
	"log"
	"reflect"
	"encoding/json"
	"github.com/golang-jwt/jwt"

	"github.com/nhatlang19/go-monorepo/pkg/helper"
)

type TokenService interface {
	ValidateIDToken(string) (*helper.UserInfo, error)
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
	fmt.Println("var1 = ", reflect.TypeOf(claims["user"]))
	jsonbody, err := json.Marshal(claims["user"])
    if err != nil {
        return nil, err
    }

	user := helper.UserInfo{}
	if err := json.Unmarshal(jsonbody, &user); err != nil {
        return nil, err
    }
	return &user, nil
}