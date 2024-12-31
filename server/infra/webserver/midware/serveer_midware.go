package midware

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type RequestCounter struct {
	QueryCount int64
	Mutex      sync.Mutex
}

// NewRequestCounter retorna uma nova instância de RequestCounter.
func NewRequestCounter() *RequestCounter {
	return &RequestCounter{}
}

// Middleware para contar requisições.
func (rc *RequestCounter) CountRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rc.Mutex.Lock()
		rc.QueryCount++
		rc.Mutex.Unlock()
		return c.Next()
	}
}

// GetQueryCount retorna a contagem de consultas de forma segura.
func (rc *RequestCounter) GetQueryCount() int64 {
	rc.Mutex.Lock()
	defer rc.Mutex.Unlock()
	return rc.QueryCount
}
