package entity

type CnpjEntity struct {
	Id        int
	Name      string
	City      string
	State     string
	CnpjNumber string
}

type CnpjValidatorInterface interface {
	ConverteToIntArray(cnpj string) ([]int, error)
	Verify(cnpjNumber []int) bool
}

type CnpjInterface interface {
	Create(cnpjData *CnpjEntity) (CnpjEntity, error)
	GetCnpjs() ([]CnpjEntity, error)
	GetCnpj(cnpj string) (CnpjEntity, error)
	Update(cnpjData *CnpjEntity, cnpj string) (CnpjEntity, error)
	Delete(cnpj string) (bool, error)
	AddCnpjToBlockList(cnpj string)(bool, error)
	GetCnpjBlockList()([]CnpjEntity, error)
	RemoveCnpjFromBlockList(cnpj string) (bool, error)
}