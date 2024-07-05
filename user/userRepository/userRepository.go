package userRepository

import (
	"errors"
	"strings"
	"tansan/config"
	"tansan/user/userModel"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        db: config.DB,
    }
}

func (r *UserRepository) FindAllUsers() ([]userModel.User, error) {
    var users []userModel.User
    if result := r.db.Find(&users); result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*userModel.User, error) {
    var user userModel.User
    if result := r.db.Where("email = ?", email).First(&user); result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) SaveUser(user *userModel.User) error {
    if result := r.db.Create(user); result.Error != nil {
        if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
            return errors.New("email already exists")
        }
        return result.Error
    }
    return nil
}

func (r *UserRepository) UpdateUserPassword(email, newPassword string) error {
    var user userModel.User
    if result := r.db.Where("email = ?", email).First(&user); result.Error != nil {
        return result.Error
    }

    user.Password = newPassword
    if result := r.db.Save(&user); result.Error != nil {
        return result.Error
    }
    return nil
}
