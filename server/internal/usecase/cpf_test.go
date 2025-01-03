package usecase_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/mathgod152/CFPcheck/cmd/dbinit"
	"github.com/mathgod152/CFPcheck/infra/database"
	"github.com/mathgod152/CFPcheck/internal/dto"
	"github.com/mathgod152/CFPcheck/internal/usecase"
	"github.com/stretchr/testify/assert"
)

var (
	CpfImplementation = &database.CpfRepository{Db: dbinit.DB}
)

func TestNewCpf(t *testing.T) (){
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}
	f := fuzz.New()

	newCpf := generateCpf(f)

	createCpfResponse, err := cpfUseCase.NewCpf(*newCpf)
	if len(newCpf.CpfNumber) != 11 && len(newCpf.CpfNumber) != 14 {
		expectedError := errors.New("CPF Invalido")
		assert.Equal(t, expectedError, err)
	}
	if len(newCpf.CpfNumber) == 11{
		assert.Equal(t, newCpf.Name, createCpfResponse.Name)
		assert.Equal(t, newCpf.State, createCpfResponse.State)
		assert.Equal(t, newCpf.City, createCpfResponse.City)
		assert.Equal(t, newCpf.CpfNumber, createCpfResponse.CpfNumber)
		assert.NoError(t, err)
	}
}


func TestSelectCpfs(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}
	f := fuzz.New()
	newCpf := generateCpf(f)
	_, err := cpfUseCase.NewCpf(*newCpf) //Garante que um CPF vai existir

	cpfs, err := cpfUseCase.SelectCpfs()

	assert.NoError(t, err)
	assert.NotNil(t, cpfs)
	assert.Greater(t, len(cpfs), 0)
}

func TestSelectCpfById(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	_, err := cpfUseCase.NewCpf(*newCpf) //Garante que um CPF vai existir

	testId := newCpf.CpfNumber
	cpf, err := cpfUseCase.SelectById(testId)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.NotNil(t, cpf)
		assert.Equal(t, testId, cpf.CpfNumber)
		assert.NoError(t, err)
	}
}

func TestUpdateCpf(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}
	f := fuzz.New()
	newCpf := generateCpf(f)
	newCpfToupdate, err := cpfUseCase.NewCpf(*newCpf) //Garante que um CPF vai existir
	if err != nil{
		expectedError := errors.New("CPF Invalido")
		assert.Equal(t, expectedError, err)
	}

	updatedCpf := updateCpf(f)
	cpfToUpdate := newCpfToupdate.CpfNumber 

	updateResponse, err := cpfUseCase.UpdateCpf(*updatedCpf, cpfToUpdate)
	fmt.Println("UpdateResponse : ", updateResponse.City)
	fmt.Println("UpdateCPF : ", updatedCpf.City)
	if len(newCpf.CpfNumber) != 11 && len(newCpf.CpfNumber) != 14 {
		expectedError := errors.New("CPF Invalido")
		assert.Equal(t, expectedError, err)
	}
	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.Equal(t, updatedCpf.CpfNumber, updateResponse.CpfNumber)
		assert.NoError(t, err)
	}
}

func TestDeleteCpf(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	_, err := cpfUseCase.NewCpf(*newCpf) //Garante que um CPF vai existir

	testId := newCpf.CpfNumber  
	response, err := cpfUseCase.DeleteCpf(testId)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.Equal(t, response, true)
		assert.NoError(t, err)
	}
}

func TestAddToBlockListCpf(t *testing.T) {
	cpfUseCase := &usecase.CpfUsecase{
		CpfInterface: CpfImplementation,
	}

	f := fuzz.New()
	newCpf := generateCpf(f)
	_, err := cpfUseCase.NewCpf(*newCpf) //Garante que um CPF vai existir

	testId := newCpf.CpfNumber  
	response, err := cpfUseCase.AddCpfToBlockList(testId)

	if err != nil {
		assert.Equal(t, errors.New("CPF not found"), err)
	} else {
		assert.Equal(t, response, true)
		assert.NoError(t, err)
	}
}

func generateCpf(f *fuzz.Fuzzer) *dto.CpfDTO {
    newCpf := &dto.CpfDTO{}
    f.Fuzz(&newCpf.Name)
    f.Fuzz(&newCpf.State)
    f.Fuzz(&newCpf.City)

    var cpfNumber []int
    for len(cpfNumber) < 11 {
        newNumber := rand.Intn(9) 
        cpfNumber = append(cpfNumber, newNumber)
    }
	fmt.Println("CPF GENERATE: ", string(convertIntArrayToString(cpfNumber)))
    newCpf.CpfNumber = string(convertIntArrayToString(cpfNumber))
    return newCpf
}

func updateCpf(f *fuzz.Fuzzer) *dto.CpfDTO {
	newCpf := &dto.CpfDTO{}
    f.Fuzz(&newCpf.Name)
    f.Fuzz(&newCpf.State)
    f.Fuzz(&newCpf.City)

    var cpfNumber []int
    for len(cpfNumber) < 11 {
        newNumber := rand.Intn(9) 
        cpfNumber = append(cpfNumber, newNumber)		
    }
	fmt.Println("CPF UPDATE: ", string(convertIntArrayToString(cpfNumber)))
    newCpf.CpfNumber = string(convertIntArrayToString(cpfNumber))
    return newCpf
}

