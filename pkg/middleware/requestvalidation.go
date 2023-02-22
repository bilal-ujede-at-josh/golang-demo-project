package middleware

import (
	"io"
	"ispick-project-21022023/pkg/apiresponse"
	"log"
	"net/http"
)

// var validate *validator.Validate

type Context struct {
	Mobile string
	Otp    int
}

func InputRequestValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.

		_, err := io.ReadAll(r.Body)

		if err != nil {
			apiresponse.PrintJsonResponse(w, "JSON body not found", http.StatusBadRequest, nil)
			// controller.PrintError(, w, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
