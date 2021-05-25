package docs

import "github.com/marimell09/stone-challenge/models"

// swagger:route POST /login login postLogin
// Login with user information.
// responses:
//   200: loginResponse

// swagger:parameters postLogin
type loginParamsWrapper struct {
	// Credential body for user login
	// in:body
	Body models.Credentials
}

// Credential reponse body
// swagger:response loginResponse
type loginResponseWrapper struct {
	// in:body
	Body models.Credentials
}
