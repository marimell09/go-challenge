package docs

import "github.com/marimell09/stone-challenge/models"

// swagger:route POST /transfers transfer addTransfer
// Create a new transfer for the logged user.
// responses:
//   200: transferResponse

// swagger:parameters addTransfer
type transferParamsWrapper struct {
	// Account body for new transfer creation
	// in:body
	Body models.Transfer
}

// Transfer creation reponse body
// swagger:response transferResponse
type transferResponseWrapper struct {
	// in:body
	Body models.Transfer
}

// swagger:route GET /transfers transfer getTransfers
// Get all transfer for the logged user.
// responses:
//   200: getTransfersResponse

// Get Transfer reponse body
// swagger:response getTransfersResponse
type getTransfersResponseWrapper struct {
	// in:body
	Body models.TransferList
}
