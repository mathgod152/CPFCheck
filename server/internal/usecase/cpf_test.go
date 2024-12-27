package usecase_test

import (
	"fmt"
	"math/rand"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T){
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

func generateTrueCPF(f *fuzz.Fuzzer) ([]int) {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCPFFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCPFSecondVerifierDigit(cpf))

	return cpf
}
func generateFakeCPF(f *fuzz.Fuzzer) ([]int) {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCPFFakeFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCPFFakeSecondVerifierDigit(cpf))

	return cpf
}


func generateFirstNineNumbers() ([]int) {
	var cpfFirstNine []int
	for len(cpfFirstNine) < 9 {
		newNumber := rand.Intn(9)
		cpfFirstNine = append(cpfFirstNine, newNumber)
		fmt.Println("Novo Array: ", cpfFirstNine)
	}
	if allElementsEqual(cpfFirstNine){
		cpfFirstNine = generateFirstNineNumbers()
	}
	return cpfFirstNine
}

func generateCPFFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit =  (verifyDigit * 10) % 11
	return verifyDigit
}

func generateCPFSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit =  (verifyDigit * 10) % 11
	return verifyDigit
}

func generateCPFFakeFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit =  (verifyDigit * 10) % 11
	return verifyDigit + 1
}

func generateCPFFakeSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit =  (verifyDigit * 10) % 11
	return verifyDigit + 1
}


func allElementsEqual[T comparable](arr []T) bool {
	if len(arr) == 0 {
		return true // Considera vazio como "todos iguais"
	}

	first := arr[0]
	for _, v := range arr {
		if v != first {
			return false
		}
	}
	return true
}
