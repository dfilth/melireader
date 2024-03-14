package models

type CurrencyResponse struct {
	ID            string `json:"id"`
	Symbol        string `json:"symbol"`
	Description   string `json:"description"`
	DecimalPlaces int    `json:"decimal_places"`
}
