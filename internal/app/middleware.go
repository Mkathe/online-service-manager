package app

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *server) useMiddlewares() {
	s.app.Use(cors.New(cors.Config{
		AllowMethods:     "GET, POST, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		AllowOrigins:     "*",
	}))

	prometheus := fiberprometheus.New("chatty")
	prometheus.RegisterAt(s.app, "/metrics")
	s.app.Use(prometheus.Middleware)

	swaggerCfg := swagger.Config{
		BasePath: "/",
		FilePath: "swagger.yaml",
		Path:     "swagger",
		Title:    "EffMobile",
	}
	s.app.Use(swagger.New(swaggerCfg))
}
