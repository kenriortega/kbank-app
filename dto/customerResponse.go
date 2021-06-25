package dto

type CustomerResponse struct {
	ID          string `json:"_id" `
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
