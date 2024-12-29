package usecase

import (
	"errors"

	"github.com/mathgod152/CFPcheck/internal/dto"
	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfUseCase struct {
	CpfInterface entity.CpfInterface
}

func (c *CpfUseCase) NewCpf(input dto.CpfDTO) (dto.CpfDTO, error) {
	cpf := entity.CpfEntity{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CpfNumber: input.CpfNumber,
	}
	if _, err := c.CpfInterface.Create(&cpf); err != nil {
		return dto.CpfDTO{}, err
	}
	return dto.CpfDTO{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CpfNumber: input.CpfNumber,
	}, nil
}

func (c *CpfUseCase) SelectCpfs() ([]dto.CpfDTO, error) {
	entities, err := c.CpfInterface.GetCpfs()
	if err != nil {
		return nil, errors.New("Erro ao Receber os Dados")
	}
	var dtos []dto.CpfDTO
	if len(entities) == 0 {
		return nil, nil
	}
	for _, entity := range entities {
		dtos = append(dtos, dto.CpfDTO{
			Name:      entity.Name,
			City:      entity.City,
			State:     entity.State,
			CpfNumber: entity.CpfNumber,
		})
	}
	return dtos, nil
}

func (c *CpfUseCase) SelectById(cpf []int) (dto.CpfDTO, error) {
	entity, err := c.CpfInterface.GetCpf(cpf)
	if err != nil {
		return dto.CpfDTO{}, errors.New("Erro ao Receber os Dados")
	}
	return dto.CpfDTO{
		Name:      entity.Name,
		City:      entity.City,
		State:     entity.State,
		CpfNumber: entity.CpfNumber,
	}, nil
}

func (c *CpfUseCase) UpdateCpf(input dto.CpfDTO) (dto.CpfDTO, error) {
	cpf := entity.CpfEntity{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CpfNumber: input.CpfNumber,
	}
	updatedEntity, err := c.CpfInterface.Update(&cpf)
	if err != nil {
		return dto.CpfDTO{}, err
	}
	return dto.CpfDTO{
		Name:      updatedEntity.Name,
		City:      updatedEntity.City,
		State:     updatedEntity.State,
		CpfNumber: updatedEntity.CpfNumber,
	}, nil
}

func (c *CpfUseCase) DeleteCpf(cpf []int) (bool, error) {
	deletedEntity, err := c.CpfInterface.Delete(cpf)
	if err != nil {
		return deletedEntity, err
	}
	return deletedEntity, nil
}
