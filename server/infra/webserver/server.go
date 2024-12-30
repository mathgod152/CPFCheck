package web

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/mathgod152/CFPcheck/infra/webserver/handllers"
	"github.com/mathgod152/CFPcheck/internal/entity"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

var _ entity.Server = (*Server)(nil) 

type Server struct {
	CpfValidator *usecase.CpfValidatorUseCase
	Cpf *usecase.CpfUsecase
}

func (s *Server) Start(port string) error {
	apiRouter := &handlers.Router{
		CpfValidator:   s.CpfValidator,
		Cpf: s.Cpf,
	}

	app := fiber.New()
	router := app.
		Group("/api/v1")

	router.Post("/cpfValidator", apiRouter.ValidateCpfHandler)
	router.Post("/savecpf", apiRouter.CreateCpfHandler)
	router.Get("/cpfs", apiRouter.GetAllCpfsHandler)
	router.Get("/cpf/:cpf", apiRouter.GetCpfHandler)
	router.Put("/cpf/:cpf", apiRouter.UpdateCpfHandler)
	router.Delete("/cpf/:cpf", apiRouter.DeleteCpfHandler)

	return app.Listen(port)
}