package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// go run *.go
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env, not comming through %v", err)
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := &handler{}

	e.POST("/login", h.login)
	e.GET("/private", h.private, isLoggedIn)
	e.GET("/admin", h.private, isLoggedIn, isAdmin)
	e.POST("/token", h.token)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}