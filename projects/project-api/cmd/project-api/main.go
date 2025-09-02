package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/config"
	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/http/handlers/student"
	"github.com/slayer/autorestart"
)

func main() {
	fmt.Println("Welcome...")
	// load config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	slog.Info("server started...", slog.String("address", cfg.Http.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	autorestart.StartWatcher()
	go func() {
		server.ListenAndServe()
	}()
	<-done

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut down server...", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown success...")

}
