package controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "oat-lab-module/cmd/docs"
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
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s *Server) initRoutes() {
	auth := s.router.Group("/auth")
	{
		auth.GET("/check-user", s.CheckUser)           // ✔
		auth.POST("/sign-in", s.SignIn)                // ✔
		auth.POST("/sign-up", s.SignUp)                // ✔
		auth.POST("/check-code", s.CheckCode)          // ✔
		auth.POST("/send-code-again", s.SendCodeAgain) // ✔
	}

	news := s.router.Group("/news")
	{
		news.POST("", s.CreateNews)   // ✔
		news.GET("", s.GetNews)       // ✔
		news.DELETE("", s.DeleteNews) // ✔
	}

	order := s.router.Group("/order")
	{
		order.POST("", s.CreateOrder) //
		order.GET("", s.GetOrder)     //
	}

	profile := s.router.Group("/profile")
	{
		profile.POST("", s.FillProfile)     //
		profile.GET("", s.GetProfileInfo)   //
		profile.DELETE("", s.DeleteProfile) //
	}

	items := s.router.Group("/item")
	{
		items.POST("", s.CreateItem) //
		items.GET("", s.GetItems)    //
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
