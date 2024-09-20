package http

import (
	"backend-test/internal/app/controller"
	"backend-test/internal/config"
	"backend-test/internal/constants"
	"backend-test/internal/health"
	"backend-test/pkg/logger"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	configuration      *config.Configuration
	healthController   *health.HealthController
	registerController *controller.RegisterController
	logger             *logger.Log
	server             *http.Server
}

func NewServer(healthController *health.HealthController, register *controller.RegisterController, configuration *config.Configuration, logger *logger.Log) *Server {
	return &Server{
		configuration:      configuration,
		logger:             logger,
		healthController:   healthController,
		registerController: register,
	}
}

func (s *Server) Start() {
	s.setupRoutes()
	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.configuration.Port),
	}
	fmt.Println(1, "Http server started on port: %d", s.configuration.Port)
	if err := s.server.ListenAndServe(); err != nil {
		switch err {
		case http.ErrServerClosed:
			s.logger.Infof("Http server closed")
		default:
			s.logger.Infof("Server-Start" + err.Error())
		}
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Infof("Shutting down http server...")
	return s.server.Shutdown(ctx)
}

func (s *Server) setupRoutes() {
	http.HandleFunc(constants.HealthResource, s.healthController.HealthCheck)
	//http.HandleFunc(constants.RegisterEntrance, s.registerController.RegisterEntrance)
	http.HandleFunc(constants.RegisterEntrance, s.handleRegisterEntrance)
	http.HandleFunc(constants.RegisterExit, s.registerController.RegisterExit)
}

func (s *Server) handleRegisterEntrance(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.registerController.RegisterEntrance(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
