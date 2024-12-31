package web

import (
	"time"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/mathgod152/CFPcheck/infra/webserver/handllers"
	"github.com/mathgod152/CFPcheck/infra/webserver/midware"
	"github.com/mathgod152/CFPcheck/internal/entity"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

var _ entity.Server = (*Server)(nil)

type Server struct {
	CpfValidator   *usecase.CpfValidatorUseCase
	Cpf            *usecase.CpfUsecase
	CnpjValidator   *usecase.CnpjValidatorUseCase
	Cnpj            *usecase.CnpjUsecase
	StartTime      time.Time
	RequestCounter *midware.RequestCounter
}

func (s *Server) Start(port string) error {
	// Inicializa o contador de requisições
	s.RequestCounter = midware.NewRequestCounter()
	s.StartTime = time.Now()

	// Configura o roteador
	apiRouter := &handlers.Router{
		StartTime:      s.StartTime,
		RequestCounter: s.RequestCounter,
		CpfValidator:   s.CpfValidator,
		Cpf:            s.Cpf,
	}

	app := fiber.New()

	app.Use(s.RequestCounter.CountRequest())

	// Roteamento da API
	router := app.Group("/api/v1")
	// CPF
	router.Post("/cpfValidator", apiRouter.ValidateCpfHandler)
	router.Post("/savecpf", apiRouter.CreateCpfHandler)
	router.Get("/cpfs", apiRouter.GetAllCpfsHandler)
	router.Get("/cpf/:cpf", apiRouter.GetCpfHandler)
	router.Put("/cpf/:cpf", apiRouter.UpdateCpfHandler)
	router.Delete("/cpf/:cpf", apiRouter.DeleteCpfHandler)

	//CNPJ
	router.Post("/cnpjValidator", apiRouter.ValidateCnpjHandler)
	router.Post("/savecnpj", apiRouter.CreateCnpjHandler)
	router.Get("/cnpjs", apiRouter.GetAllCnpjsHandler)
	router.Get("/cnpj/:cnpj", apiRouter.GetCnpjHandler)
	router.Put("/cnpj/:cnpj", apiRouter.UpdateCnpjHandler)
	router.Delete("/cnpj/:cnpj", apiRouter.DeleteCnpjHandler)

	// Servidor
	router.Get("/server-info", apiRouter.ServerInfoHandler)

	return app.Listen(port)
}
