package dto

type CreateCurrencyDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCurrencyDTO struct {
	Name *string `json:"name"`
}

type CurrencyDTO struct {
	CurrencyID int    `json:"currency_id"`
	Name       string `json:"name"`
}


