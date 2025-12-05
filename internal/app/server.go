package app

import (
	"database/sql"
	"effMobile/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

type server struct {
	app          *fiber.App
	logger       hclog.Logger
	db           *sql.DB
	storeService ServiceRepository
}

func Run() error {
	var err error
	s := new(server)
	errChan := make(chan error)

	err = s.generate()
	if err != nil {
		return err
	}

	go func() {
		err = s.app.Listen(":" + config.Get().Port)
		if err != nil {
			errChan <- err
			s.logger.Error("Error starting server", "error", err)
		}
	}()

	select {
	case err = <-errChan:
		return err
	}
}
