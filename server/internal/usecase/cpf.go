package usecase

import (
	"errors"
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/dto"
	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfUsecase struct {
	CpfInterface entity.CpfInterface
}

func (c *CpfUsecase) NewCpf(input dto.CpfDTO) (dto.CpfDTO, error) {
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

func (c *CpfUsecase) SelectCpfs() ([]dto.CpfDTO, error) {
	entities, err := c.CpfInterface.GetCpfs()
	if err != nil {
		return nil, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
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

func (c *CpfUsecase) SelectById(cpf string) (dto.CpfDTO, error) {
	entity, err := c.CpfInterface.GetCpf(cpf)
	if err != nil {
		return dto.CpfDTO{}, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
	}
	return dto.CpfDTO{
		Name:      entity.Name,
		City:      entity.City,
		State:     entity.State,
		CpfNumber: entity.CpfNumber,
	}, nil
}

func (c *CpfUsecase) UpdateCpf(input dto.CpfDTO, cpfNumber string) (dto.CpfDTO, error) {
	cpf := entity.CpfEntity{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CpfNumber: input.CpfNumber,
	}
	updatedEntity, err := c.CpfInterface.Update(&cpf, cpfNumber)
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

func (c *CpfUsecase) DeleteCpf(cpf string) (bool, error) {
	deletedEntity, err := c.CpfInterface.Delete(cpf)
	if err != nil {
		return deletedEntity, err
	}
	return deletedEntity, nil
}

func (c *CpfUsecase) AddCpfToBlockList(cpf string) (bool, error) {
	addCpfToBlockList, err := c.CpfInterface.AddCpftoBlockList(cpf)
	if err != nil {
		return addCpfToBlockList, err
	}
	return addCpfToBlockList, nil
}

func (c *CpfUsecase) RemoveCpfFromBlockList(cpf string) (bool, error) {
	addCpfToBlockList, err := c.CpfInterface.RemoveCpfFromBlockList(cpf)
	if err != nil {
		return addCpfToBlockList, err
	}
	return addCpfToBlockList, nil
}

func (c *CpfUsecase) SelectBlockListCpfs() ([]dto.CpfDTO, error) {
	entities, err := c.CpfInterface.GetCpfBlockList()
	if err != nil {
		return nil, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
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
