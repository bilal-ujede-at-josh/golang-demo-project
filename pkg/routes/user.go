package routes

import (
	handler "ispick-project-21022023/pkg/auth"
	"ispick-project-21022023/pkg/deps"
	"ispick-project-21022023/pkg/home"
	userhandler "ispick-project-21022023/pkg/user"

	"github.com/gorilla/mux"
)

func InitializeAuthRoutes(r *mux.Router, dependencies deps.Dependencies) {
	r.HandleFunc("/send-otp", handler.SendOtp(dependencies.Auth))
	r.HandleFunc("/verify-otp", handler.VerifyOtp(dependencies.Auth))
}

func InitializeAccountRoutes(r *mux.Router, dependencies deps.Dependencies) {
	r.HandleFunc("/account-details", userhandler.AccountDetails(dependencies.User))
	r.HandleFunc("/home-content", home.HomeContent(dependencies.User, dependencies.Home))
}
