package controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "oat-lab-module/docs"
	"oat-lab-module/internal/pkg/storage/pg"
	"oat-lab-module/internal/utils/config"
)

type Server struct {
	router  *gin.Engine
	storage *pg.Storage
	cfg     *config.Config
}

func NewServer(storage *pg.Storage, cfg *config.Config) *Server {
	router := gin.Default()
	server := &Server{
		router:  router,
		storage: storage,
		cfg:     cfg,
	}
	server.initRoutes()
	return server
}

func (s *Server) InitSwagger() {
	url := ginSwagger.URL("5a37-92-124-163-102.ngrok-free.app/swagger/doc.json")
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s *Server) initRoutes() {
	auth := s.router.Group("/auth")
	{
		auth.POST("/authorize", s.Authorize)           //
		auth.POST("/check-code", s.CheckCode)          // ✔
		auth.POST("/send-code-again", s.SendCodeAgain) // ✔
	}

	catalog := s.router.Group("/catalog")
	{
		catalog.POST("/get-by-id", s.GetCatalogItemById)
		catalog.POST("", s.CreateCatalogItem)   // ✔
		catalog.POST("/get", s.GetCatalogItems) // ✔
		catalog.DELETE("", s.DeleteCatalogItem) // ✔
	}

	order := s.router.Group("/order")
	{
		order.POST("", s.CreateOrder) // TODO Реализовать
		order.GET("", s.GetOrder)     // TODO Реализовать
	}

	profile := s.router.Group("/profile")
	{
		profile.POST("", s.CreateProfile)   // ✔
		profile.GET("", s.GetProfileInfo)   // ✔
		profile.DELETE("", s.DeleteProfile) // ✔
	}
	s.router.Static("/media", "./media") // ✔

}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
