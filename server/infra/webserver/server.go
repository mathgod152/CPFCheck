package web

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mathgod152/CFPcheck/infra/implemantation"
	handlers "github.com/mathgod152/CFPcheck/infra/webserver/handllers"
	"github.com/mathgod152/CFPcheck/internal/entity"
	"github.com/mathgod152/CFPcheck/internal/usecase"
)

var _ entity.Server = (*Server)(nil)

type Server struct {
	CpfValidator   *usecase.CpfValidatorUseCase
	Cpf            *usecase.CpfUsecase
	CnpjValidator  *usecase.CnpjValidatorUseCase
	Cnpj           *usecase.CnpjUsecase
	StartTime      time.Time
	RequestCounter *implemantation.RequestCounter
}

func (s *Server) Start(port string) error {
	// Inicializa o contador de requisições
	s.RequestCounter = implemantation.NewRequestCounter()
	s.StartTime = time.Now()

	// Configura o roteador
	apiRouter := &handlers.Router{
		StartTime:      s.StartTime,
		RequestCounter: s.RequestCounter,
		CpfValidator:   s.CpfValidator,
		Cpf:            s.Cpf,
	}

	app := fiber.New()

	//CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // APENAS PARA DEV
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Middleware para contar requisições
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
	router.Put("/addblocklistcpf/:cpf", apiRouter.AddToBlockListCpfHandler)
	router.Get("/blocklistcpfs", apiRouter.GetBlocklistCpfsHandler)
	router.Put("/removeblocklistcpf/:cpf", apiRouter.RemoveToBlockListCpfHandler)

	// CNPJ
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
