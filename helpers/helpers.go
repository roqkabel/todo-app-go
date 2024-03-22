package helpers

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

type ResponseParams struct {
	StatusCode int
	Message    string
	Result     any
}

func Response(w http.ResponseWriter, r *http.Request, params ResponseParams) {
	render.Status(r, params.StatusCode)
	render.JSON(w, r, map[string]interface{}{
		"message": params.Message,
		"result":  params.Result,
	})
}

var AppTokenAuth = jwtauth.JWTAuth{}
