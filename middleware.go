package main

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// fmt.Println(os.Getenv("JWT_KEY"))

var jwtKey = os.Getenv("JWT_KEY")

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("local-ivan"),
})

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		fmt.Println(isAdmin)

		if isAdmin == false {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}
