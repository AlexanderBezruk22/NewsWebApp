package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Addr   string
	Config fiber.Config
	App    *fiber.App
}

func New(config fiber.Config, addr string) *Server {
	return &Server{Config: config, Addr: addr, App: fiber.New(config)}
}

func (s *Server) Run() error {
	if err := s.App.Listen(s.Addr); err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown(app *fiber.App) error {
	if err := app.Shutdown(); err != nil {
		return err
	}
	return nil
}
