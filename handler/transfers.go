package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marimell09/stone-challenge/models"
)

func transfers(router chi.Router) {
	router.Use(CheckToken)
	router.Get("/", getAllTransfers)
	router.Post("/", addTransfer)
}

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		id, err := strconv.Atoi(claims.Account_id)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Invalid account Id")))
		}

		ctx := context.WithValue(r.Context(), accountIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllTransfers(w http.ResponseWriter, r *http.Request) {
	account_id := r.Context().Value(accountIdKey).(int)
	transfers, err := dbInstance.GetAllTransfers(account_id)

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, transfers); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func addTransfer(w http.ResponseWriter, r *http.Request) {
	transfer := &models.Transfer{}
	account_id := r.Context().Value(accountIdKey).(int)

	if err := render.Bind(r, transfer); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	transfer.Account_origin_id = account_id

	destination_account, errDest := dbInstance.GetAccountById(transfer.Account_destination_id)
	if errDest != nil {
		render.Render(w, r, ErrorUnprocessable(ErrInvalidAccountDestination))
		return
	}

	owner_account, _ := dbInstance.GetAccountById(account_id)
	var owner_balance float64
	owner_balance = owner_account.Balance - transfer.Amount
	if owner_balance < 0 {
		render.Render(w, r, ErrorUnprocessable(ErrInvalidBalance))
		return
	}

	if err := dbInstance.AddTransfer(transfer); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.UpdateAccountBalance(account_id, owner_balance); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	var destination_balance float64
	destination_balance = destination_account.Balance + transfer.Amount
	if err := dbInstance.UpdateAccountBalance(transfer.Account_destination_id, destination_balance); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, transfer); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
