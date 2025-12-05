package app

import (
	"effMobile/internal/store"
	"effMobile/pkg/config"
	"effMobile/pkg/db/postgres"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

func (s *server) generate() error {
	s.logger = hclog.New(&hclog.LoggerOptions{
		Name:       "effMobile",
		JSONFormat: true,
	})
	s.app = fiber.New()
	db, err := postgres.LoadDatabase(fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.Get().DBUser,
		config.Get().DBPassword,
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBName,
	))
	if err != nil {
		return err
	}
	s.db = db

	s.storeService = store.NewStoreService(s.db, s.logger)

	s.useMiddlewares()
	s.initRoutes()
	return nil
}
