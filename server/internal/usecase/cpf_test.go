package usecase_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) (bool, error){
	f := fuzz.New()
	newTrueCPF := generateTrueCPF(f)

	cpfCheck, err := CPFCheckUseCaseMock.UsecaseCPF.Check(*&newTrueCPF)

	assert.Nil(t, err)
	assert.Equal(t, cpfCheck, true)

	newFakeCPF := generateFakeCPF(f)

	cpfCheck, err = CPFCheckUseCaseMock.UsecaseCPF.Check(*&newFakeCPF)
	assert.Nil(t, err)
	assert.NotEqual(t, cpfCheck, true)
}

