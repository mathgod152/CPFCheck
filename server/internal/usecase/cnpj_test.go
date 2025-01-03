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
	CnpjImplementation = &database.CnpjRepository{Db: dbinit.DB}
)

func TestNewCnpj(t *testing.T) (){
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}
	f := fuzz.New()

	newCnpj := generateCnpj(f)

	createCnpjResponse, err := cnpjUseCase.NewCnpj(*newCnpj)
	if len(newCnpj.CnpjNumber) != 11 && len(newCnpj.CnpjNumber) != 14 {
		expectedError := errors.New("Cnpj Invalido")
		assert.Equal(t, expectedError, err)
	}
	if len(newCnpj.CnpjNumber) == 11{
		assert.Equal(t, newCnpj.Name, createCnpjResponse.Name)
		assert.Equal(t, newCnpj.State, createCnpjResponse.State)
		assert.Equal(t, newCnpj.City, createCnpjResponse.City)
		assert.Equal(t, newCnpj.CnpjNumber, createCnpjResponse.CnpjNumber)
		assert.NoError(t, err)
	}
}


func TestSelectCnpjs(t *testing.T) {
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}
	f := fuzz.New()
	newCnpj := generateCnpj(f)
	_, err := cnpjUseCase.NewCnpj(*newCnpj) //Garante que um Cnpj vai existir

	cnpjs, err := cnpjUseCase.SelectCnpjs()

	assert.NoError(t, err)
	assert.NotNil(t, cnpjs)
	assert.Greater(t, len(cnpjs), 0)
}

func TestSelectCnpjById(t *testing.T) {
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}

	f := fuzz.New()
	newCnpj := generateCnpj(f)
	_, err := cnpjUseCase.NewCnpj(*newCnpj) //Garante que um Cnpj vai existir

	testId := newCnpj.CnpjNumber
	cnpj, err := cnpjUseCase.SelectById(testId)

	if err != nil {
		assert.Equal(t, errors.New("Cnpj not found"), err)
	} else {
		assert.NotNil(t, cnpj)
		assert.Equal(t, testId, cnpj.CnpjNumber)
		assert.NoError(t, err)
	}
}

func TestUpdateCnpj(t *testing.T) {
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}
	f := fuzz.New()
	newCnpj := generateCnpj(f)
	newCnpjToupdate, err := cnpjUseCase.NewCnpj(*newCnpj) //Garante que um Cnpj vai existir
	if err != nil{
		expectedError := errors.New("Cnpj Invalido")
		assert.Equal(t, expectedError, err)
	}

	updatedCnpj := updateCnpj(f)
	cnpjToUpdate := newCnpjToupdate.CnpjNumber 

	updateResponse, err := cnpjUseCase.UpdateCnpj(*updatedCnpj, cnpjToUpdate)
	fmt.Println("UpdateResponse : ", updateResponse.City)
	fmt.Println("UpdateCnpj : ", updatedCnpj.City)
	if len(newCnpj.CnpjNumber) != 11 && len(newCnpj.CnpjNumber) != 14 {
		expectedError := errors.New("Cnpj Invalido")
		assert.Equal(t, expectedError, err)
	}
	if err != nil {
		assert.Equal(t, errors.New("Cnpj not found"), err)
	} else {
		assert.Equal(t, updatedCnpj.CnpjNumber, updateResponse.CnpjNumber)
		assert.NoError(t, err)
	}
}

func TestDeleteCnpj(t *testing.T) {
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}

	f := fuzz.New()
	newCnpj := generateCnpj(f)
	_, err := cnpjUseCase.NewCnpj(*newCnpj) //Garante que um Cnpj vai existir

	testId := newCnpj.CnpjNumber  
	response, err := cnpjUseCase.DeleteCnpj(testId)

	if err != nil {
		assert.Equal(t, errors.New("Cnpj not found"), err)
	} else {
		assert.Equal(t, response, true)
		assert.NoError(t, err)
	}
}

func TestAddToBlockList(t *testing.T) {
	cnpjUseCase := &usecase.CnpjUsecase{
		CnpjInterface: CnpjImplementation,
	}

	f := fuzz.New()
	newCnpj := generateCnpj(f)
	_, err := cnpjUseCase.NewCnpj(*newCnpj) //Garante que um Cnpj vai existir

	testId := newCnpj.CnpjNumber  
	response, err := cnpjUseCase.AddCnpjToBlockList(testId)

	if err != nil {
		assert.Equal(t, errors.New("Cnpj not found"), err)
	} else {
		assert.Equal(t, response, true)
		assert.NoError(t, err)
	}
}

func generateCnpj(f *fuzz.Fuzzer) *dto.CnpjDTO {
    newCnpj := &dto.CnpjDTO{}
    f.Fuzz(&newCnpj.Name)
    f.Fuzz(&newCnpj.State)
    f.Fuzz(&newCnpj.City)

    var cnpjNumber []int
    for len(cnpjNumber) < 11 {
        newNumber := rand.Intn(9) 
        cnpjNumber = append(cnpjNumber, newNumber)
    }
	fmt.Println("Cnpj GENERATE: ", string(convertIntArrayToString(cnpjNumber)))
    newCnpj.CnpjNumber = string(convertIntArrayToString(cnpjNumber))
    return newCnpj
}

func updateCnpj(f *fuzz.Fuzzer) *dto.CnpjDTO {
	newCnpj := &dto.CnpjDTO{}
    f.Fuzz(&newCnpj.Name)
    f.Fuzz(&newCnpj.State)
    f.Fuzz(&newCnpj.City)

    var cnpjNumber []int
    for len(cnpjNumber) < 11 {
        newNumber := rand.Intn(9) 
        cnpjNumber = append(cnpjNumber, newNumber)		
    }
	fmt.Println("Cnpj UPDATE: ", string(convertIntArrayToString(cnpjNumber)))
    newCnpj.CnpjNumber = string(convertIntArrayToString(cnpjNumber))
    return newCnpj
}

