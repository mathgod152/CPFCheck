package dto

type CnpjDTO struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	State     string `json:"state"`
	CnpjNumber string `json:"cnpjNumber"`
}
