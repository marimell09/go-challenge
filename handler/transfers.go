package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func transfers(router chi.Router) {
	router.Get("/", getAllTransfers)
}

func getAllTransfers(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			render.Render(w, r, ErrNotAuthorized)
			return
		}
		render.Render(w, r, ErrBadRequest)
		return
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			render.Render(w, r, ErrNotAuthorized)
			return
		}
		render.Render(w, r, ErrBadRequest)
		return
	}
	if !tkn.Valid {
		render.Render(w, r, ErrNotAuthorized)
		return
	}

	transfers, err := dbInstance.GetAllTransfers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, transfers); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
