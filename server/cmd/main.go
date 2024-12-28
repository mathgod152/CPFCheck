package main

import (
	"fmt"
	"math/rand"
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

func main() {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCpfFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfSecondVerifierDigit(cpf))

	fmt.Println("Cpf: ", cpf)
	fmt.Println("Cpf Sem os digitos: ", cpf[:9])
}
