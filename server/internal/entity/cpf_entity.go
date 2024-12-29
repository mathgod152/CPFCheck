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
	Create(cpfData CpfEntity) (CpfEntity, error)
	Read() []CpfEntity
	ReadByCpf(cpf []int) []CpfEntity
	Updae(cpfData CpfEntity) (CpfEntity, error)
}
