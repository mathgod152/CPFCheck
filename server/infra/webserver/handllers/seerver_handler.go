package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)


func (r *Router) ServerInfoHandler(c *fiber.Ctx) error {
	uptime := time.Since(r.StartTime).String()
	queries := r.RequestCounter.GetQueryCount()

	return c.JSON(fiber.Map{
		"up_time":     uptime,
		"query_count": queries,
	})
}