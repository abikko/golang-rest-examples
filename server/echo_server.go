package server

import (
	"awesomeProject/config"
	"awesomeProject/shop/handlers"
	"awesomeProject/shop/repositories"
	"awesomeProject/shop/usecases"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.initializeShopHttpHandler()
	s.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (server *echoServer) initializeShopHttpHandler() {
	shopPostgresRepository := repositories.NewShopPostgresRepository(server.db)

	shopUsecase := usecases.NewShopUseCaseImpl(
		shopPostgresRepository,
	)

	shopHttpHandler := handlers.NewShopHttpHandler(shopUsecase)

	// Routers
	shopRouters := server.app.Group("v1/shop")
	shopRouters.POST("", shopHttpHandler.DetectShop)
}
