package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marimell09/stone-challenge/db"
	"github.com/marimell09/stone-challenge/models"
	"github.com/marimell09/stone-challenge/utils"
)

var credentials_key = "credentials_key"
var expiration_time int = 3000

var secret_key string = os.Getenv("TOKEN_SECRET_KEY")

func login(router chi.Router) {
	router.Post("/", loginUser)
	//router.Use(UserContext)
}

/* func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		credentials := &models.Credentials{}

		if err := render.Bind(r, credentials); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), credentials_key, credentials)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
} */

type Claims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	credentials := &models.Credentials{}

	if err := render.Bind(r, credentials); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	secret, err := dbInstance.GetAccountByCpf(credentials.Cpf)
	var hashedPassword = utils.HashSecret(credentials.Secret)

	//compare := bcrypt.CompareHashAndPassword([]byte(secret), []byte(credentials.Secret))
	//fmt.Printf("compare: %s\n", compare)

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
		Cpf: credentials.Cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret_key))
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
