package main

import (
	myGrpc "github.com/AnKlvy/notifier-edunite/internal/services/grpc"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr     string
	notifier notifier.NotifyService
	server   *grpc.Server
}

func NewGRPCServer(addr string, notifier notifier.NotifyService) *GRPCServer {
	return &GRPCServer{addr: addr, notifier: notifier}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.server = grpc.NewServer()

	// register our grpc services
	newsService := s.notifier
	myGrpc.NewNotifierService(s.server, newsService)

	log.Println("Starting gRPC server on", s.addr)

	return s.server.Serve(lis)
}

// waitForShutdown блокирует до получения системного сигнала и корректно останавливает gRPC-сервер.
func waitForShutdown(server *grpc.Server) {
	// Создаём канал, в который пойдут сигналы ОС
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Received shutdown signal, gracefully stopping gRPC server...")
	server.GracefulStop()
}
