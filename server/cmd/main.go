package main

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/mathgod152/CFPcheck/cmd/dbinit"
	"github.com/mathgod152/CFPcheck/infra/database"
	"github.com/mathgod152/CFPcheck/infra/implemantation"
	web "github.com/mathgod152/CFPcheck/infra/webserver"
	"github.com/mathgod152/CFPcheck/internal/entity"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

var (
	serverImpl entity.Server
	DB *sql.DB
	err error
)

func init(){

	var(
		CpfValidateImplementation = &implemantation.CpfValidatorImplementation{}
		CpfImplementation = &database.CpfRepository{Db: dbinit.DB}
		runCpfValidator = &usecase.CpfValidatorUseCase{
			CpfValidatorEntity: CpfValidateImplementation, // Injetando a implementação 
		}
		runCpf = &usecase.CpfUsecase{
			CpfInterface: CpfImplementation,
		}
	)

	serverImpl = &web.Server{
		CpfValidator: runCpfValidator,
		Cpf: runCpf,
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