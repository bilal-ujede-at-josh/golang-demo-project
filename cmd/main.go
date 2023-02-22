package main

import (
	"ispick-project-21022023/pkg/auth"
	"ispick-project-21022023/pkg/config"
	"ispick-project-21022023/pkg/database"
	"ispick-project-21022023/pkg/deps"
	"ispick-project-21022023/pkg/home"
	"ispick-project-21022023/pkg/jwt"
	"ispick-project-21022023/pkg/routes"
	"ispick-project-21022023/pkg/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	config.LoadEnvConfig()
	database.InitDatabase()
}

func main() {
	// init mux router
	r := mux.NewRouter()

	// Initialize dependencies
	appDb := database.GetDb()
	jwt := jwt.NewJwtService()
	auth := auth.NewAuthService(appDb, jwt)
	usr := user.NewUserService(appDb, jwt)
	home := home.NewHomeService(appDb, jwt)
	dependencies := deps.Dependencies{
		Auth: auth,
		User: usr,
		Home: home,
	}
	// // init handler functions with mux router
	routes.InitializeAuthRoutes(r, dependencies)
	routes.InitializeAccountRoutes(r, dependencies)

	// // Start server
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

	log.Println("Server started successfully!")
}
