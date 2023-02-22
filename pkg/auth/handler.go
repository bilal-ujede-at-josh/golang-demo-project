package auth

import (
	"encoding/json"
	"ispick-project-21022023/pkg/apiresponse"
	"ispick-project-21022023/pkg/user"
	"log"
	"net/http"
)

func SendOtp(as AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode input json request body
		var sor SendOtpRequest
		err := json.NewDecoder(r.Body).Decode(&sor)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Error processing input ", http.StatusInternalServerError, nil)
			return
		}

		// Create OTP
		log.Println(sor.Mobile)
		err = as.SendOtp(r.Context(), sor.Mobile)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Error sending OTP ", http.StatusInternalServerError, nil)
			return
		}
		apiresponse.PrintJsonResponse(w, "OTP Send successfully ", http.StatusCreated, nil)
	}
}

func VerifyOtp(as AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		// Decode input json request body
		var vor VerifyOtpRequest
		err := json.NewDecoder(r.Body).Decode(&vor)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Error processing input ", http.StatusInternalServerError, nil)
			return
		}

		// Validate OTP
		log.Println(vor.Mobile, vor.Otp)
		err = as.VerifyOtp(r.Context(), vor.Mobile, vor.Otp)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Incorrect otp", http.StatusInternalServerError, nil)
			return
		}
		log.Println(user)

		// Generate get user_id from mobile
		token, err := as.JwtToken(r.Context(), vor.Mobile)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Incorrect otp", http.StatusInternalServerError, nil)
			return
		}
		apiresponse.PrintJsonResponse(w, "OTP Verified successfully ", http.StatusCreated, token)
	}
}
