package usecase_test

import (
	"errors"
	"math"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/mathgod152/CFPcheck/infra/database"
	"github.com/mathgod152/CFPcheck/internal/dto"
	"github.com/mathgod152/CFPcheck/internal/usecase"
	"github.com/stretchr/testify/assert"
)

var (
	CpfImplementation *database.CpfRepository
)

func TestNewCpf(t *testing.T) (){
	cpfUseCase := &usecase.CpfUseCase{
		CpfInterface: CpfImplementation,
	}
	f := fuzz.New()

	newCpf := generateCpf(f)

	createCpfResponse, err := cpfUseCase.NewCpf(newCpf)
	if len(newCpf.CpfNumber) > 11 || len(newCpf.CpfNumber) < 11 {
		expectedError := errors.New("CPF Invalido")
		assert.Equal(t, expectedError, err)
	}
	if len(newCpf.CpfNumber) == 11{
		assert.Equal(t, newCpf.Id, createCpfResponse.Id)
		assert.Equal(t, newCpf.Name, createCpfResponse.Name)
		assert.Equal(t, newCpf.State, createCpfResponse.State)
		assert.Equal(t, newCpf.City, createCpfResponse.City)
		assert.Equal(t, newCpf.CpfNumber, createCpfResponse.CpfNumber)
		assert.NoError(t, err)
	}
}


func TestSelectCpfs(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfEntity: CpfImplementation,
	}
	f := fuzz.New()
	newCpf := generateCpf(f)
	_, err := cpfUseCase.NewCpf(newCpf) //Garante que um CPF vai existir

	cpfs, err := cpfUseCase.SelectCpfs()

	assert.NoError(t, err)
	assert.NotNil(t, cpfs)
	assert.Greater(t, len(cpfs), 0)
}

func TestSelectCpfById(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfEntity: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	newCpf, err := cpfUseCase.NewCpf(newCpf) //Garante que um CPF vai existir

	testId := newCpf.Id
	cpf, err := cpfUseCase.SelectById(testId)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.NotNil(t, cpf)
		assert.Equal(t, testId, cpf.Id)
		assert.NoError(t, err)
	}
}

func TestUpdateCpf(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfEntity: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	newCpf, err := cpfUseCase.NewCpf(newCpf) //Garante que um CPF vai existir

	updatedCpf := updateCpf(f)
	updatedCpf.Id = newCpf.Id 

	updateResponse, err := cpfUseCase.Update(updatedCpf)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.Equal(t, updatedCpf.Id, updateResponse.Id)
		assert.Equal(t, updatedCpf.Name, updateResponse.Name)
		assert.Equal(t, updatedCpf.State, updateResponse.State)
		assert.Equal(t, updatedCpf.City, updateResponse.City)
		assert.Equal(t, updatedCpf.CpfNumber, updateResponse.CpfNumber)
		assert.NoError(t, err)
	}
}

func TestDeleteCpf(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfEntity: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	newCpf, err := cpfUseCase.NewCpf(newCpf) //Garante que um CPF vai existir

	testId := newCpf.Id  
	err = cpfUseCase.Delete(testId)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.NoError(t, err)
	}
}

func generateCpf(f *fuzz.Fuzzer) *dto.CpfDTO {
	newCpf := &dto.CpfDTO{}
	f.Fuzz(&newCpf.Id)
	f.Fuzz(&newCpf.Name)
	f.Fuzz(&newCpf.State)
	f.Fuzz(&newCpf.City)
	for len(newCpf.CpfNumber) < 9 {
		var newNumber int
		f.Fuzz(&newNumber)
		newNumber = int(math.Abs(float64(newNumber))) % 10
		newCpf.CpfNumber = append(newCpf.CpfNumber, newNumber)
	}
	return newCpf
}

func updateCpf(f *fuzz.Fuzzer) *dto.CpfDTO {
	newCpf := &dto.CpfDTO{}
	f.Fuzz(&newCpf.Name)
	f.Fuzz(&newCpf.State)
	f.Fuzz(&newCpf.City)
	for len(newCpf.CpfNumber) < 9 {
		var newNumber int
		f.Fuzz(&newNumber)
		newNumber = int(math.Abs(float64(newNumber))) % 10
		newCpf.CpfNumber = append(newCpf.CpfNumber, newNumber)
	}
	return newCpf
}
