package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JwtJSON ...
type JwtJSON struct {
	Aud  string `json:"aud"`
	Exp int `json:"exp"`
	UserID int `json:"user_id"`
	UserFullname  string `json:"user_fullname"`
	UserRole  []UserRole `json:"user_role"`
	Bagian  string `json:"bagian"`
	Jabatan  string `json:"jabatan"`
	IDJabatan int `json:"id_jabatan"`
	IDBagian int `json:"id_bagian"`
}

// UserRole ...
type UserRole struct {
	RoleID  int `json:"role_id"`
	NamaRole string `json:"nama_role"`
}

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	// Claims: &JwtJSON{},
	// SigningKey: []byte("secret"),
	SigningKey: []byte(os.Getenv("JWT_KEY")),
})

// func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)
// 		// isAdmin := claims["admin"].(bool)
// 		userFullname := claims["user_fullname"]
// 		// user_fullname
// 		fmt.Println(user)

// 		// if isAdmin == false {
// 		// 	return echo.ErrUnauthorized
// 		// }
// 		if userFullname != "ADMIN EMONICA" {
// 			return echo.ErrUnauthorized
// 		}

// 		return next(c)
// 	}
// }

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userFullname := claims["user_fullname"]
		fmt.Println(user)

		if userFullname != "ADMIN EMONICA" {
			return echo.ErrUnauthorized
		}

		token, error := jwt.Parse(user.Raw, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if error != nil {
			// json.NewEncoder(w).Encode(ErrorMsg{Message: error.Error()})
			return c.String(http.StatusOK, "error!")
		}
		if token.Valid {
			fmt.Println("token", token.Valid)

			// var user User
			// mapstructure.Decode(token.Claims, &user)

			// vars := mux.Vars(req)
			// name := vars["userId"]
			// if name != user.Username {
			// 	// json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
			// 	return c.String(http.StatusOK, "Valid!")
			// 	// return
			// }

			// context.Set(req, "decoded", token.Claims)
			// next(w, req)
		} else {
			// json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
			return c.String(http.StatusOK, "Invalid authorization token!")
		}

		return next(c)
	}
}
