package handlers

import (
	"github.com/gin-gonic/gin"
	"zatrasz75/tz_market/configs"
	"zatrasz75/tz_market/internal/repository"
	"zatrasz75/tz_market/pkg/logger"
)

// ErrorResponse структура для отображения ошибок API
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Server -.
type Server struct {
	l    logger.LoggersInterface
	repo *repository.Store
	cfg  *configs.Config
}

// @title Swagger API
// @version 1.0
// @description ТЗ market_backend.
// @description Задание (Golang + PostgreSQL + Gin)

// @contact.name Михаил Токмачев
// @contact.url https://t.me/Zatrasz
// @contact.email zatrasz@ya.ru

// @BasePath /

// NewRouter -.
func NewRouter() *gin.Engine {
	r := gin.Default()

	return r
}
