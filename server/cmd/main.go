package main

import (
	"os"
	"os/signal"

	"github.com/mathgod152/CFPcheck/infra/implemantation"
	web "github.com/mathgod152/CFPcheck/infra/webserver"
	"github.com/mathgod152/CFPcheck/internal/entity"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

var (
	serverImpl entity.Server
)

func init(){

	var(
		CpfValidateImplementation = &implemantation.CpfValidatorImplementation{}
		runCpfValidator = &usecase.CpfValidatorUseCase{
			CpfValidatorEntity: CpfValidateImplementation, // Injetando a implementação 
		}
	)

	serverImpl = &web.Server{
		CpfValidator: runCpfValidator,
	}
}


func main() {
	go serverImpl.Start(":5000")

	listen()
}

func listen() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	os.Exit(130)
}