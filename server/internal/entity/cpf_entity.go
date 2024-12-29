package entity

type CpfEntity struct {
	Id        int
	Name      string
	City      string
	State     string
	CpfNumber []int
}

type CpfValidatorInterface interface {
	ConverteToIntArray(cpf string) ([]int, error)
	Verify(cpfNumber []int) bool
}

type CpfInterface interface {
	Create(cpfData *CpfEntity) (CpfEntity, error)
	GetCpfs() ([]CpfEntity, error)
	GetCpf(cpf []int) (CpfEntity, error)
	Update(cpfData *CpfEntity, cpf []int) (CpfEntity, error)
	Delete(cpf []int) (bool, error)
}
