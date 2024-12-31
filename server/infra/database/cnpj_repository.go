package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CnpjInterface = (*CnpjRepository)(nil)

type CnpjRepository struct {
	Db *sql.DB
}

func (c *CnpjRepository) Create(cnpjData *entity.CnpjEntity) (entity.CnpjEntity, error) {
	if len(cnpjData.CnpjNumber) != 14  &&  len(cnpjData.CnpjNumber) != 18 {
		return entity.CnpjEntity{}, errors.New("Cnpj Invalido")
	}
	stmt, err := c.Db.Prepare("INSERT INTO cnpj (name, city, state, cnpj_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return entity.CnpjEntity{}, err
	}
	_, err = stmt.Exec(cnpjData.Name, cnpjData.City, cnpjData.State, cnpjData.CnpjNumber)
	if err != nil {
		return entity.CnpjEntity{}, err
	}
	return entity.CnpjEntity{
		Name:      cnpjData.Name,
		City:      cnpjData.City,
		State:     cnpjData.State,
		CnpjNumber: cnpjData.CnpjNumber,
	}, nil
}

func (c *CnpjRepository) GetCnpjs() ([]entity.CnpjEntity, error) {
	rows, err := c.Db.Query("SELECT id, name, city, state, cnpj_number FROM cnpj")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cnpjs := []entity.CnpjEntity{}
	for rows.Next() {
		var id int
		var name, city, state, cnpj_number string

		if err := rows.Scan(&id, &name, &city, &state, &cnpj_number); err != nil {
			return nil, err
		}
		application := entity.CnpjEntity{
			Id:        id,
			Name:      name,
			State:     city,
			City:      state,
			CnpjNumber: cnpj_number,
		}
		cnpjs = append(cnpjs, application)
	}
	return cnpjs, nil
}

func (c *CnpjRepository) GetCnpj(cnpj string) (entity.CnpjEntity, error) {
	rows, err := c.Db.Query("SELECT id, name, city, state, cnpj_number FROM cnpj WHERE cnpj_number = $1", cnpj)
	if err != nil {
		return entity.CnpjEntity{}, err
	}
	defer rows.Close()
	cnpjs := entity.CnpjEntity{}
	for rows.Next() {
		var id int
		var name, city, state, cnpj_number string

		if err := rows.Scan(&id, &name, &city, &state, &cnpj_number); err != nil {
			return entity.CnpjEntity{}, err
		}
		cnpjs = entity.CnpjEntity{
			Id:        id,
			Name:      name,
			State:     city,
			City:      state,
			CnpjNumber: cnpj_number,
		}
	}
	fmt.Println("Cnpj RETURNS: ", cnpjs)
	return cnpjs, nil
}

func (c *CnpjRepository) Update(cnpjData *entity.CnpjEntity, cnpj string) (entity.CnpjEntity, error) {
	fmt.Println("Cnpj: ", cnpj, "LEN: ", len(cnpj))
	if len(cnpjData.CnpjNumber) != 14  &&  len(cnpjData.CnpjNumber) != 18 {
		return entity.CnpjEntity{}, errors.New("Cnpj Invalido")
	}

	query := "UPDATE cnpj SET "
	args := []interface{}{}
	i := 1

	if cnpjData.Name != "" {
		query += "name = $" + strconv.Itoa(i) + ", "
		args = append(args, cnpjData.Name)
		i++
	}
	if cnpjData.City != "" {
		query += "city = $" + strconv.Itoa(i) + ", "
		args = append(args, cnpjData.City)
		i++
	}
	if cnpjData.State != "" {
		query += "state = $" + strconv.Itoa(i) + ", "
		args = append(args, cnpjData.State)
		i++
	}
	if cnpjData.CnpjNumber != "" && cnpjData.CnpjNumber != cnpj {
		query += "cnpj_number = $" + strconv.Itoa(i) + ", "
		args = append(args, cnpjData.CnpjNumber)
		i++
	}

	if len(args) == 0 {
		return entity.CnpjEntity{}, errors.New("nenhum campo para atualizar")
	}

	query = query[:len(query)-2] // Remove ", "
	query += " WHERE cnpj_number = $" + strconv.Itoa(i)
	args = append(args, cnpj)

	fmt.Println("Query final: ", query)
	fmt.Println("Args: ", args)

	stmt, err := c.Db.Prepare(query)
	if err != nil {
		return entity.CnpjEntity{}, err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return entity.CnpjEntity{}, err
	}
	return c.GetCnpj(cnpjData.CnpjNumber)
}

func (c *CnpjRepository) Delete(cnpj string) (bool, error) {
	stmt, err := c.Db.Prepare("DELETE FROM cnpj WHERE cnpj_number = $1")
	if err != nil {
		return false, err
	}
	result, err := stmt.Exec(cnpj)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, errors.New("nenhum registro encontrado para deletar")
	}
	return true, nil
}


