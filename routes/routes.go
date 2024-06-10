package routes

import (
	"net/http"
	"strings"
	"tansan/auth"
	"tansan/user/userHandler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    e.GET("/users", userHandler.GetUsers, authMiddleware)
    e.POST("/users", userHandler.CreateUser, authMiddleware)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return c.JSON(http.StatusUnauthorized, "Missing token")
        }

        // Remove "Bearer " prefix from the token
        token = strings.TrimPrefix(token, "Bearer ")

        claims, err := auth.ValidateToken(token)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, err.Error())
        }

        c.Set("user", claims.Email)
        return next(c)
    }
}


