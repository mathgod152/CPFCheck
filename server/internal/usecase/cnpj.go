package usecase

import (
	"errors"
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/dto"
	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CnpjUsecase struct {
	CnpjInterface entity.CnpjInterface
}

func (c *CnpjUsecase) NewCnpj(input dto.CnpjDTO) (dto.CnpjDTO, error) {
	cnpj := entity.CnpjEntity{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CnpjNumber: input.CnpjNumber,
	}
	if _, err := c.CnpjInterface.Create(&cnpj); err != nil {
		return dto.CnpjDTO{}, err
	}
	return dto.CnpjDTO{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CnpjNumber: input.CnpjNumber,
	}, nil
}

func (c *CnpjUsecase) SelectCnpjs() ([]dto.CnpjDTO, error) {
	entities, err := c.CnpjInterface.GetCnpjs()
	if err != nil {
		return nil, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
	}
	var dtos []dto.CnpjDTO
	if len(entities) == 0 {
		return nil, nil
	}
	for _, entity := range entities {
		dtos = append(dtos, dto.CnpjDTO{
			Name:      entity.Name,
			City:      entity.City,
			State:     entity.State,
			CnpjNumber: entity.CnpjNumber,
		})
	}
	return dtos, nil
}

func (c *CnpjUsecase) SelectById(cnpj string) (dto.CnpjDTO, error) {
	entity, err := c.CnpjInterface.GetCnpj(cnpj)
	if err != nil {
		return dto.CnpjDTO{}, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
	}
	return dto.CnpjDTO{
		Name:      entity.Name,
		City:      entity.City,
		State:     entity.State,
		CnpjNumber: entity.CnpjNumber,
	}, nil
}

func (c *CnpjUsecase) UpdateCnpj(input dto.CnpjDTO, cnpjNumber string) (dto.CnpjDTO, error) {
	cnpj := entity.CnpjEntity{
		Name:      input.Name,
		City:      input.City,
		State:     input.State,
		CnpjNumber: input.CnpjNumber,
	}
	updatedEntity, err := c.CnpjInterface.Update(&cnpj, cnpjNumber)
	if err != nil {
		return dto.CnpjDTO{}, err
	}
	return dto.CnpjDTO{
		Name:      updatedEntity.Name,
		City:      updatedEntity.City,
		State:     updatedEntity.State,
		CnpjNumber: updatedEntity.CnpjNumber,
	}, nil
}

func (c *CnpjUsecase) DeleteCnpj(cnpj string) (bool, error) {
	deletedEntity, err := c.CnpjInterface.Delete(cnpj)
	if err != nil {
		return deletedEntity, err
	}
	return deletedEntity, nil
}

func (c *CnpjUsecase) AddCnpjToBlockList(cnpj string) (bool, error) {
	addToBlockList, err := c.CnpjInterface.AddCnpjToBlockList(cnpj)
	if err != nil {
		return addToBlockList, err
	}
	return addToBlockList, nil
}

func (c *CnpjUsecase) RemoveCnpjFromBlockList(cnpj string) (bool, error) {
	addCnpjToBlockList, err := c.CnpjInterface.RemoveCnpjFromBlockList(cnpj)
	if err != nil {
		return addCnpjToBlockList, err
	}
	return addCnpjToBlockList, nil
}

func (c *CnpjUsecase) SelectBlockListCnpjs() ([]dto.CnpjDTO, error) {
	entities, err := c.CnpjInterface.GetCnpjBlockList()
	if err != nil {
		return nil, errors.New("Erro ao Receber os Dados" + fmt.Sprint(err))
	}
	var dtos []dto.CnpjDTO
	if len(entities) == 0 {
		return nil, nil
	}
	for _, entity := range entities {
		dtos = append(dtos, dto.CnpjDTO{
			Name:      entity.Name,
			City:      entity.City,
			State:     entity.State,
			CnpjNumber: entity.CnpjNumber,
		})
	}
	return dtos, nil
}

