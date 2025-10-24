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
	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/storage/sqlite"
	"github.com/slayer/autorestart"
)

func main() {
	fmt.Println("Welcome...")
	// load config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		slog.Error("Error connecting Sqlite database.")
	}
	slog.Info("Storage initialised", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
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
