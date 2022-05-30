package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type GracefulShutdown func(ctx context.Context)

type ServerConfig struct {
	Port             int
	ReadTimeout      int
	WriteTimeout     int
	IdleTimeout      int
	ShoutDownTimeout int
}

type Server interface {
	Start()
	AddGracefulShutdown(downFunc GracefulShutdown)
}

type server struct {
	s        *http.Server
	config   ServerConfig
	handler  http.Handler
	downFunc GracefulShutdown
}

func NewServer(config ServerConfig, handler http.Handler) Server {
	return &server{config: config, handler: handler}
}

func (s *server) Start() {
	s.s = &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", s.config.Port),
		WriteTimeout: time.Duration(s.config.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(s.config.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.IdleTimeout) * time.Second,
		Handler:      s.handler,
	}
	log.Printf("http server starting on port: %d\n", s.config.Port)
	go s.s.ListenAndServe()
	s.waitForSignal()
}

func (s *server) AddGracefulShutdown(downFunc GracefulShutdown) {
	s.downFunc = downFunc
}

func (s *server) waitForSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case sig := <-c:
			log.Printf("Got %s signal. Aborting...\n", sig)
			gracefulShutdown(s)
			return
		}
	}
}

func gracefulShutdown(s *server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.ShoutDownTimeout)*time.Second)
	defer cancel()
	if s.downFunc != nil {
		s.downFunc(ctx)
	}
	if err := s.s.Shutdown(ctx); err != nil {
		log.Fatalf("error in shutting down server: %s", err)
	}
}
