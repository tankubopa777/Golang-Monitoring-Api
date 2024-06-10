package userUsecase

import (
	"tansan/user/userModel"
	"tansan/user/userRepository"
)

func GetAllUsers(repo *userRepository.UserRepository) ([]userModel.User, error) {
	return repo.FindAllUsers()
}

func CreateUser(repo *userRepository.UserRepository, user *userModel.User) error {
	return repo.SaveUser(user)
}
