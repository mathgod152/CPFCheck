package dto

type CpfDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	State string `json:"state"`
	CpfNumber []int `json:"cpfNumber"`
}
