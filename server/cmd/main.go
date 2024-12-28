package main

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

func generateFirstNineNumbers() []int {
	var cpfFirstNine []int
	for len(cpfFirstNine) < 9 {
		newNumber := rand.Intn(9)
		cpfFirstNine = append(cpfFirstNine, newNumber)
		fmt.Println("Novo Array: ", cpfFirstNine)
	}
	return cpfFirstNine
}

func generateCpfFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	return verifyDigit
}

func generateCpfSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	return verifyDigit
}

func cpfstringToArray (cpf string) ([]int, error){
		// Remove todos os caracteres não numéricos
		cpfDigits := regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

		// Verifica se o CPF tem exatamente 11 dígitos
		if len(cpfDigits) != 11 {
			return nil, errors.New("o CPF deve conter exatamente 11 dígitos")
		}
	
		// Converte os dígitos para um array de inteiros
		intArray := make([]int, len(cpfDigits))
		for i, char := range cpfDigits {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, errors.New("erro ao converter CPF para inteiros")
			}
			intArray[i] = num
		}
	
		return intArray, nil
}

func main() {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCpfFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfSecondVerifierDigit(cpf))
	cpfarray, err := cpfstringToArray("503.278.618-78")
	if err != nil{
		fmt.Println("ERRO AO CONVERTER CPF: ", err)
	}
	fmt.Println("CpfArray: ", cpfarray)
	fmt.Println("Cpf: ", cpf)
	fmt.Println("Cpf Sem os digitos: ", cpf[:9])
}
