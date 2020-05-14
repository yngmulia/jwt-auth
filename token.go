package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateTokenPair() (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = 1
	// claims["name"] = "Jon Doe"
	// claims["admin"] = true

	claims["id"] = 1
	claims["email"] = "superadmin@twiscode.com"
	claims["type"] = "admin"
	// claims["iat"] = time.Now().Add(time.Minute * 15).Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	fmt.Println("JWT_KEY", os.Getenv("JWT_KEY"))
	
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = 1
	// rtClaims["sub"] = 1
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
