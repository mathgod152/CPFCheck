package database

import (
	"database/sql"
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CpfInterface = (*CpfRepository)(nil)

type CpfRepository struct {
	Db *sql.DB
}

func (c *CpfRepository) Create(cpfData *entity.CpfEntity) (entity.CpfEntity, error) {
	stmt, err := c.Db.Prepare("INSERT INTO cpf (name, city, state, cpf_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return entity.CpfEntity{}, err
	}
	_, err = stmt.Exec(cpfData.Name, cpfData.City, cpfData.State, convertIntArrayToString(cpfData.CpfNumber))
	if err != nil {
		return entity.CpfEntity{}, err
	}
	return entity.CpfEntity{
		Name:      cpfData.Name,
		City:      cpfData.City,
		State:     cpfData.State,
		CpfNumber: cpfData.CpfNumber,
	}, nil
}

func (c *CpfRepository) GetCpfs() ([]entity.CpfEntity, error) {
	return nil, nil
}

func (c *CpfRepository) GetCpf(cpf []int) (entity.CpfEntity, error) {
	return entity.CpfEntity{}, nil
}

func (c *CpfRepository) Update(cpfData *entity.CpfEntity, cpf []int) (entity.CpfEntity, error) {
	return entity.CpfEntity{}, nil
}

func (c *CpfRepository) Delete(cpf []int) (bool, error) {
	return true, nil
}

func convertIntArrayToString(cpf []int) string {
	cpfStr := ""
	for i, num := range cpf {
		cpfStr += fmt.Sprintf("%d", num)
		if i == 2 || i == 5 {
			cpfStr += "."
		} else if i == 8 {
			cpfStr += "-"
		}
	}
	return cpfStr
}