package user

import (
	"ispick-project-21022023/pkg/apiresponse"
	"ispick-project-21022023/pkg/helper"
	"net/http"
)

func AccountDetails(us UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract token from request
		tokenString := helper.ExtractToken(r)

		// Validate user request token
		user_id, err := us.ValidateJwtToken(r.Context(), tokenString)

		if err != nil {
			apiresponse.PrintJsonResponse(w, "Invalid token", http.StatusUnauthorized, nil)
			return
		}

		// get user details from user_id
		user, err := us.GetUserById(r.Context(), user_id)

		if err != nil {
			apiresponse.PrintJsonResponse(w, "User details not found", http.StatusNotFound, nil)
			return
		}
		apiresponse.PrintJsonResponse(w, "User details found", http.StatusOK, user)
	}
}
