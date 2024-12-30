package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CpfInterface = (*CpfRepository)(nil)

type CpfRepository struct {
	Db *sql.DB
}

func (c *CpfRepository) Create(cpfData *entity.CpfEntity) (entity.CpfEntity, error) {
	if len(cpfData.CpfNumber) != 11  &&  len(cpfData.CpfNumber) != 14 {
		return entity.CpfEntity{}, errors.New("CPF Invalido")
	}
	stmt, err := c.Db.Prepare("INSERT INTO cpf (name, city, state, cpf_number) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return entity.CpfEntity{}, err
	}
	_, err = stmt.Exec(cpfData.Name, cpfData.City, cpfData.State, cpfData.CpfNumber)
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
	rows, err := c.Db.Query("SELECT id, name, city, state, cpf_number FROM cpf")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cpfs := []entity.CpfEntity{}
	for rows.Next() {
		var id int
		var name, city, state, cpf_number string

		if err := rows.Scan(&id, &name, &city, &state, &cpf_number); err != nil {
			return nil, err
		}
		application := entity.CpfEntity{
			Id:        id,
			Name:      name,
			State:     city,
			City:      state,
			CpfNumber: cpf_number,
		}
		cpfs = append(cpfs, application)
	}
	return cpfs, nil
}

func (c *CpfRepository) GetCpf(cpf string) (entity.CpfEntity, error) {
	rows, err := c.Db.Query("SELECT id, name, city, state, cpf_number FROM cpf WHERE cpf_number = $1", cpf)
	if err != nil {
		return entity.CpfEntity{}, err
	}
	defer rows.Close()
	cpfs := entity.CpfEntity{}
	for rows.Next() {
		var id int
		var name, city, state, cpf_number string

		if err := rows.Scan(&id, &name, &city, &state, &cpf_number); err != nil {
			return entity.CpfEntity{}, err
		}
		cpfs = entity.CpfEntity{
			Id:        id,
			Name:      name,
			State:     city,
			City:      state,
			CpfNumber: cpf_number,
		}
	}
	return cpfs, nil
}

func (c *CpfRepository) Update(cpfData *entity.CpfEntity, cpf string) (entity.CpfEntity, error) {
	fmt.Println("CPF: ", cpf, "LEN: ", len(cpf))
	if len(cpf) != 11 && len(cpf) != 14 {
		return entity.CpfEntity{}, errors.New("CPF inválido")
	}

	query := "UPDATE cpf SET "
	args := []interface{}{}
	i := 1

	// Adiciona os campos dinamicamente
	if cpfData.Name != "" {
		query += "name = $" + strconv.Itoa(i) + ", "
		args = append(args, cpfData.Name)
		i++
	}
	if cpfData.City != "" {
		query += "city = $" + strconv.Itoa(i) + ", "
		args = append(args, cpfData.City)
		i++
	}
	if cpfData.State != "" {
		query += "state = $" + strconv.Itoa(i) + ", "
		args = append(args, cpfData.State)
		i++
	}
	if cpfData.CpfNumber != "" && cpfData.CpfNumber != cpf {
		query += "cpf_number = $" + strconv.Itoa(i) + ", "
		args = append(args, cpfData.CpfNumber)
		i++
	}

	// Verifica se ao menos um campo foi alterado
	if len(args) == 0 {
		return entity.CpfEntity{}, errors.New("nenhum campo para atualizar")
	}

	// Remove a última vírgula e adiciona a cláusula WHERE
	query = query[:len(query)-2] // Remove ", "
	query += " WHERE cpf_number = $" + strconv.Itoa(i)
	args = append(args, cpf)

	// Loga a query gerada e os argumentos para depuração
	fmt.Println("Query final: ", query)
	fmt.Println("Args: ", args)

	// Prepara e executa a query
	stmt, err := c.Db.Prepare(query)
	if err != nil {
		return entity.CpfEntity{}, err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return entity.CpfEntity{}, err
	}

	// Retorna o CPF atualizado
	return c.GetCpf(cpf)
}

func (c *CpfRepository) Delete(cpf string) (bool, error) {
	stmt, err := c.Db.Prepare("DELETE FROM cpf WHERE cpf_number = $1")
	if err != nil {
		return false, err
	}
	result, err := stmt.Exec(cpf)
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


