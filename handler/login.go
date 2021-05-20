package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marimell09/stone-challenge/db"
	"github.com/marimell09/stone-challenge/models"
	"github.com/marimell09/stone-challenge/utils"
)

var credentials_key = "credentials_key"
var expiration_time int = 5

var secret_key []byte = []byte(os.Getenv("TOKEN_SECRET_KEY"))

func login(router chi.Router) {
	router.Post("/", loginUser)
}

type Claims struct {
	Cpf        string `json:"cpf"`
	Account_id string `json:"account_id"`
	jwt.StandardClaims
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	credentials := &models.Credentials{}

	if err := render.Bind(r, credentials); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	id, secret, err := dbInstance.GetAccountByCpf(credentials.Cpf)
	var hashedPassword = utils.HashSecret(credentials.Secret)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotAuthorized)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if secret != hashedPassword {
		render.Render(w, r, ErrNotAuthorized)
	}

	expirationTime := time.Now().Add(time.Duration(expiration_time) * time.Minute)

	claims := &Claims{
		Cpf:        credentials.Cpf,
		Account_id: strconv.Itoa(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
