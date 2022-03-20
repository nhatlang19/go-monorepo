package service

import (
	"fmt"
	"os"
	"strconv"

	helper "github.com/nhatlang19/go-monorepo/pkg/helper"
	"github.com/nhatlang19/go-monorepo/pkg/model"
	grpc_client "github.com/nhatlang19/go-monorepo/services/user/client"
	"github.com/nhatlang19/go-monorepo/services/user/repository"

	"log"
)

type UserService interface {
	Save(model.User) (model.User, error)
	Login(email string, password string) (model.User, error, int)
}

type userService struct {
	userRepository repository.UserRepository
	mailClient     grpc_client.MailClient
}

func NewUserService(r repository.UserRepository, m grpc_client.MailClient) UserService {
	return userService{
		userRepository: r,
		mailClient:     m,
	}
}

func (u userService) Login(email string, password string) (model.User, error, int) {
	log.Print("[UserService]...Login")

	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return user, fmt.Errorf("User Not Found"), 404
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return user, fmt.Errorf("Password incorrect"), 403
	}

	return user, nil, 200
}

func (u userService) Save(user model.User) (model.User, error) {
	log.Print("[UserService]...Save")
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hashedPassword
	user, err = u.userRepository.Save(user)
	if err != nil {
		panic(err)
	}

	enable_mail, err := strconv.ParseBool(os.Getenv("SEND_MAIL"))
	if err != nil {
		panic(err)
	}

	if enable_mail {
		u.mailClient.HandleRegisterMail(user)
	}

	return user, err
}
