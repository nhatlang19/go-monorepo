package service

import (
	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/repository"

	"log"
)

type UserService interface {
	Save(model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService (r repository.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (u userService) Save(user model.User) (model.User, error) {
	log.Print("[UserService]...Save")
	return u.userRepository.Save(user)
}
