package home

import (
	"ispick-project-21022023/pkg/apiresponse"
	"ispick-project-21022023/pkg/helper"
	"ispick-project-21022023/pkg/model"
	"ispick-project-21022023/pkg/user"
	"net/http"
)

func HomeContent(us user.UserService, hs HomeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appHomeScreen model.AppHomeScreen

		// Get JWT token from request
		requestToken := helper.ExtractToken(r)

		// Validate user based on JWT token
		user_id, err := us.ValidateJwtToken(r.Context(), requestToken)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "Un-authorized", http.StatusUnauthorized, nil)
			return
		}

		// Get user details
		var userDetails model.UserDetails
		userDetails, err = us.GetUserById(r.Context(), user_id)
		if err != nil {
			apiresponse.PrintJsonResponse(w, "User not found", http.StatusNotFound, nil)
			return
		}
		appHomeScreen.User_details = userDetails

		if userDetails.Is_approved == 1 {
			appHomeScreen.AccountStatus = "approved"
		} else {
			appHomeScreen.AccountStatus = "un-approved"
		}
		if userDetails.Is_completed == 1 {
			appHomeScreen.KycStatus = "completed"
		} else {
			appHomeScreen.KycStatus = "incomplete"
		}

		// Get App Version Details
		var appVersion model.AppVersion
		appVersion, _ = hs.AppVersion(r.Context())

		appHomeScreen.App_versions = appVersion
		apiresponse.PrintJsonResponse(w, "Details fetched successfully!", http.StatusOK, appHomeScreen)
	}
}
