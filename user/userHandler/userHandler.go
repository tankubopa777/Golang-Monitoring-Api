package userHandler

import (
	"net/http"
	"tansan/auth"
	"tansan/user/userModel"
	"tansan/user/userRepository"
	"tansan/user/userUsecase"
	"tansan/utils"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
    users, err := userUsecase.GetAllUsers(userRepository.NewUserRepository())
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
    var user userModel.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    err := userUsecase.CreateUser(userRepository.NewUserRepository(), &user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, user)
}


func Register(c echo.Context) error {
    var user userModel.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    user.Password = hashedPassword

    if err := userRepository.NewUserRepository().SaveUser(&user); err != nil {
        if err.Error() == "email already exists" {
            return c.JSON(http.StatusConflict, err.Error()) // HTTP 409 Conflict
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, user)
}


func Login(c echo.Context) error {
    var input userModel.User
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    user, err := userRepository.NewUserRepository().FindUserByEmail(input.Email)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, "Invalid email or password")
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        return c.JSON(http.StatusUnauthorized, "Invalid email or password")
    }

    token, err := auth.GenerateJWT(user.Email)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, echo.Map{
        "token": token,
    })
}
