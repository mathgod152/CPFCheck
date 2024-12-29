package database

import (
	"database/sql"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CpfInterface = (*CpfRepository)(nil)

type CpfRepository struct {
	Db *sql.DB
}

func (c *CpfRepository) Create(cpfData *entity.CpfEntity) (entity.CpfEntity, error) {
	return entity.CpfEntity{}, nil
}


func (c *CpfRepository) GetCpfs() ([]entity.CpfEntity, error) {
	return nil, nil
}


func (c *CpfRepository) GetCpf(cpf []int) (entity.CpfEntity, error) {
	return entity.CpfEntity{}, nil
}


func (c *CpfRepository) Update(cpfData *entity.CpfEntity) (entity.CpfEntity, error) {
	return entity.CpfEntity{}, nil
}

func (c *CpfRepository) Delete(cpf []int) (bool, error) {
	return true, nil
}
