package controllers

import (
	"github.com/gin-gonic/gin"
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

func (s *Server) initRoutes() {
	auth := s.router.Group("/auth")
	{
		auth.POST("/signIn", s.SignIn)
		auth.POST("/signUp", s.SignUp)
		auth.POST("/checkCode", s.CheckCode)
	}

	news := s.router.Group("/news")
	{
		news.POST("/createNews", s.CreateNews)
		news.GET("/getNews", s.GetNews)
	}

	order := s.router.Group("/order")
	{
		order.POST("/createOrder", s.CreateOrder)
		order.GET("/getOrder", s.GetOrder)
	}

	profile := s.router.Group("/profile")
	{
		profile.POST("/fillProfile", s.FillProfile)
		profile.GET("/getProfileIngo", s.GetProfileInfo)
	}

	catalog := s.router.Group("/catalog")
	{
		catalog.GET("/getCatalog", s.GetCatalog)
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
