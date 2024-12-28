package entity

type CpfEntity struct {
	Name string
	City string
	State string
	CpfNumber []int
}

type CpfInterface interface {
	Verify(cpfNumber []int) (bool, error)
	Create(cpfData CpfEntity) (CpfEntity, error)
	Read()([]CpfEntity)
	ReadByCpf(cpf []int)([]CpfEntity)
	Updae(cpfData CpfEntity) (CpfEntity, error)
}

