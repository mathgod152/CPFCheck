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
}

func (s *Server) Start(port string) error {
	apiRouter := &handlers.Router{
		CpfValidator:   s.CpfValidator,
	}

	app := fiber.New()
	router := app.
		Group("/api/v1")

	router.Post("/cpfValidator", apiRouter.ValidateCpfHandler)

	return app.Listen(port)
}