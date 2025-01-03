package web

import (
	"fmt"
	"os"
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
	// Imprime o diretório de trabalho atual para verificação
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("Current working directory:", cwd)

	// Inicializa o contador de requisições
	s.RequestCounter = implemantation.NewRequestCounter()
	s.StartTime = time.Now()

	// Configura o roteador
	apiRouter := &handlers.Router{
		StartTime:      s.StartTime,
		RequestCounter: s.RequestCounter,
		CpfValidator:   s.CpfValidator,
		Cpf:            s.Cpf,
		CnpjValidator:  s.CnpjValidator,
		Cnpj:           s.Cnpj,
	}

	app := fiber.New()
	fmt.Println("Mudou")
	app.Static("/", "../public")

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // APENAS PARA DEV
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Middleware para contar requisições
	app.Use(s.RequestCounter.CountRequest())

	// Middleware de logging para debug (verificando as requisições)
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Request URL:", c.OriginalURL())
		return c.Next()
	})

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
	router.Put("/addblocklistcnpj/:cnpj", apiRouter.AddToBlockListCnpjHandler)
	router.Get("/blocklistcnpjs", apiRouter.GetBlocklistCnpjsHandler)
	router.Put("/removeblocklistcnpj/:cnpj", apiRouter.RemoveToBlockListCnpjHandler)

	// Servidor
	router.Get("/server-info", apiRouter.ServerInfoHandler)

	return app.Listen(port)
}
