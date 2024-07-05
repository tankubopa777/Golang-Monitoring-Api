package userUsecase

import (
	"errors"
	"tansan/user/userModel"
	"tansan/user/userRepository"
	"tansan/utils"
)

func GetAllUsers(repo *userRepository.UserRepository) ([]userModel.User, error) {
	return repo.FindAllUsers()
}

func CreateUser(repo *userRepository.UserRepository, user *userModel.User) error {
	return repo.SaveUser(user)
}

func ChangeUserPassword(repo *userRepository.UserRepository, email, currentPassword, newPassword string) error {
	user, err := repo.FindUserByEmail(email)
	if err != nil {
		return err
	}

	// Check if current password is correct
	if !utils.CheckPasswordHash(currentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return repo.UpdateUserPassword(email, hashedPassword)
}
