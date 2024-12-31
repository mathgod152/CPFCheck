package handlers

import (
	"time"

	"github.com/mathgod152/CFPcheck/infra/webserver/midware"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

type Router struct {
	CpfValidator *usecase.CpfValidatorUseCase
	Cpf *usecase.CpfUsecase
	Cnpj *usecase.CnpjUsecase
	CnpjValidator *usecase.CnpjValidatorUseCase
	StartTime      time.Time
	RequestCounter *midware.RequestCounter
}