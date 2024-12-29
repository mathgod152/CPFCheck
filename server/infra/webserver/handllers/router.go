package handlers

import "github.com/mathgod152/CFPcheck/internal/usecase"

type Router struct {
	CpfValidator *usecase.CpfValidatorUseCase
	Cpf *usecase.CpfUsecase
}