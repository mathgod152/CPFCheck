package entity

type CpfEntity struct {
	Id        int
	Name      string
	City      string
	State     string
	CpfNumber string
}

type CpfValidatorInterface interface {
	ConverteToIntArray(cpf string) ([]int, error)
	Verify(cpfNumber []int) bool
}

type CpfInterface interface {
	Create(cpfData *CpfEntity) (CpfEntity, error)
	GetCpfs() ([]CpfEntity, error)
	GetCpf(cpf string) (CpfEntity, error)
	Update(cpfData *CpfEntity, cpf string) (CpfEntity, error)
	Delete(cpf string) (bool, error)
	AddCpftoBlockList(cpf string) (bool, error)
	GetCpfBlockList()([]CpfEntity, error)
	RemoveCpfFromBlockList(cpf string) (bool, error)
}
