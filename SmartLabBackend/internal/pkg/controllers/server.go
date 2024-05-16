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
	swagURL := s.cfg.ServerNgrokHost + "/swagger/doc.json"
	url := ginSwagger.URL(swagURL)
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s *Server) initRoutes() {
	auth := s.router.Group("/auth")
	{
		auth.POST("/send-code", s.SendCode)            // ✔
		auth.POST("/check-code", s.CheckCode)          // ✔
		auth.POST("/send-code-again", s.SendCodeAgain) // ✔
	}

	catalog := s.router.Group("/catalog")
	{
		catalog.GET("/by-id", s.GetCatalogItemById) // ✔
		catalog.POST("", s.CreateCatalogItem)       // ✔
		catalog.GET("/get", s.GetCatalogItems)      // ✔
		catalog.DELETE("", s.DeleteCatalogItem)     // ✔
	}

	order := s.router.Group("/order")
	{
		order.POST("", s.CreateOrder) // TODO Реализовать
		order.GET("", s.GetOrder)     // TODO Реализовать
	}

	profile := s.router.Group("/med-card")
	{
		profile.POST("", s.CreateMedCard)   // ✔
		profile.GET("", s.GetMedCardsInfo)  // ✔
		profile.DELETE("", s.DeleteMedCard) // ✔
	}
	s.router.Static("/media", "./media") // ✔

}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
