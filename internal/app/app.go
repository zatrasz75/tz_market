package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"
	"zatrasz75/tz_market/configs"
	_ "zatrasz75/tz_market/docs"
	"zatrasz75/tz_market/internal/handlers"
	"zatrasz75/tz_market/internal/middleware"
	"zatrasz75/tz_market/internal/repository"
	"zatrasz75/tz_market/pkg/logger"
	"zatrasz75/tz_market/pkg/postgres"
	"zatrasz75/tz_market/pkg/server"
)

func Run(cfg *configs.Config, l logger.LoggersInterface) {
	pg, err := postgres.New(cfg.DataBase.ConnStr, l, postgres.OptionSet(cfg.DataBase.PoolMax, cfg.DataBase.ConnAttempts, cfg.DataBase.ConnTimeout))
	if err != nil {
		l.Fatal("ошибка запуска - postgres.New:", err)
	}
	defer pg.Close()

	err = pg.Migrate(l)
	if err != nil {
		l.Fatal("ошибка миграции", err)
	}

	repo := repository.New(pg)

	router := handlers.NewRouter()
	router.Use(middleware.SetHeader)
	router.Use(middleware.CreateCorsMiddleware(cfg.Server.CORSAllowedOrigins))
	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerHandlers(router, l, repo, cfg)

	srv := server.New(router, server.OptionSet(cfg.Server.AddrHost, cfg.Server.AddrPort, cfg.Server.ReadTimeout, cfg.Server.WriteTimeout, cfg.Server.IdleTimeout, cfg.Server.ShutdownTime))
	go func() {
		err = srv.Start()
		if err != nil {
			l.Error("Остановка сервера:", err)
		}
	}()

	l.Info("Запуск сервера на http://" + cfg.Server.AddrHost + ":" + cfg.Server.AddrPort)
	l.Info("CORS ALLOWED ORIGINS %v", cfg.Server.CORSAllowedOrigins)
	l.Info("Документация Swagger API: http://" + cfg.Server.AddrHost + ":" + cfg.Server.AddrPort + "/swagger/index.html")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("принят сигнал прерывания прерывание %s", s.String())
	case err = <-srv.Notify():
		l.Error("получена ошибка сигнала прерывания сервера", err)
	}

	err = srv.Shutdown()
	if err != nil {
		l.Error("не удалось завершить работу сервера", err)
	}
}

func registerHandlers(r *gin.Engine, l logger.LoggersInterface, repo *repository.Store, cfg *configs.Config) {
	mainGroup := r.Group(
		"/en",
	)

	handlers.RegisterHomeHandlers(mainGroup, l, repo, cfg)
}
