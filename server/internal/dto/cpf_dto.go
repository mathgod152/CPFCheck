package dto

type CpfDTO struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	State     string `json:"state"`
	CpfNumber string `json:"cpfNumber"`
}
