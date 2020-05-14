package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserRole ...
// type UserRole struct {
// 	RoleID  int `json:"role_id"`
// 	NamaRole string `json:"nama_role"`
// }

func generateTokenPair() (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// userRoles := []UserRole { { RoleID: 1, NamaRole: "SUPER ADMINISTRATOR"}, }

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
  // claims["aud"] = "Web"
	// claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  // claims["user_id"] = 1
	// claims["user_fullname"] = "ADMIN EMONICA"	
	// claims["user_role"] = userRoles	
	// claims["bagian"] = "TEKNIK"
  // claims["jabatan"] = "STAF"
  // claims["id_jabatan"] = 12
	// claims["id_bagian"] = 4
			
	claims["id"] = 1
	claims["email"] = "superadmin@twiscoder.com"
	claims["type"] = "admin"
	claims["iat"] = time.Now().Add(time.Minute * 15).Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// claims["sub"] = 1
	// claims["name"] = "Jon Doe"
	// claims["admin"] = true
	// claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
