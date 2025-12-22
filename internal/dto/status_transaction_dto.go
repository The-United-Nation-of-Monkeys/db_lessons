package dto

type CreateStatusTransactionDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateStatusTransactionDTO struct {
	Name *string `json:"name"`
}

type StatusTransactionDTO struct {
	StatusTransactionID int    `json:"status_transaction_id"`
	Name                string `json:"name"`
}


