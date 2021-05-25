package docs

import "github.com/marimell09/stone-challenge/models"

// swagger:route POST /accounts account addAccount
// Create a new account.
// responses:
//   200: accountResponse

// swagger:parameters addAccount
type accountParamsWrapper struct {
	// Account body for new account creation
	// in:body
	Body models.Account
}

// Account creation reponse body
// swagger:response accountResponse
type accountResponseWrapper struct {
	// in:body
	Body models.Account
}

// swagger:route GET /accounts account getAccounts
// Get all accounts.
// responses:
//   200: getAccountsResponse

// Get Accounts reponse body
// swagger:response getAccountsResponse
type getAccountsResponseWrapper struct {
	// in:body
	Body models.AccountList
}

// swagger:route GET /accounts/{account_id} account getAccount
// Get account by id.
// responses:
//   200: accountResponse

// swagger:parameters getAccount getAccountBalance
type getAccountParamsWrapper struct {
	// Account id for search
	// in:path
	Account_id int `json:"account_id"`
}

// swagger:route GET /accounts/{account_id}/balance account getAccountBalance
// Get account balance by id.
// responses:
//   200: getAccountBalanceResponse

// Get account balance reponse body
// swagger:response getAccountBalanceResponse
type getAccountBalanceResponseWrapper struct {
	// in:body
	Body struct {
		// example: 10.60
		Balance float64 `json:"balance"`
	}
}
