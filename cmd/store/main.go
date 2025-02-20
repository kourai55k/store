package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	web "github.com/kourai55k/store/internal/api/web/handlers/productHandler"
	"github.com/kourai55k/store/internal/config"
	"github.com/kourai55k/store/internal/repositories"
	"github.com/kourai55k/store/internal/services"
	prettySlog "github.com/kourai55k/store/pkg/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application")

	repo := repositories.NewInMemoryProductRepository()
	service := services.NewProductService(log, repo)
	productHandler := web.NewProductHandler(log, service)

	// repo.SaveProduct(models.Product{
	// 	ID:           1,
	// 	Name:         "Product 1",
	// 	Price:        100,
	// 	Stock:        true,
	// 	Params:       `{"weight": "10kg", "color": "red"}`,
	// 	CategoryName: "Category 1",
	// 	Measure:      "kg",
	// })
	// repo.SaveProduct(models.Product{
	// 	Name:         "Product 2",
	// 	Price:        100,
	// 	Stock:        true,
	// 	Params:       `{"weight": "10kg", "color": "red"}`,
	// 	CategoryName: "Category 1",
	// 	Measure:      "kg",
	// })
	// repo.SaveProduct(models.Product{
	// 	Name:         "Product 3",
	// 	Price:        100,
	// 	Stock:        true,
	// 	Params:       `{"weight": "10kg", "color": "red"}`,
	// 	CategoryName: "Category 1",
	// 	Measure:      "kg",
	// })

	r := http.NewServeMux()
	// Serve static files
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	// Routes for web
	r.HandleFunc("GET /products", productHandler.GetProducts)
	r.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	r.HandleFunc("GET /product/new", productHandler.CreateProductPage)
	r.HandleFunc("POST /product/new", productHandler.CreateProduct)

	// Routes for REST API

	// Start server
	log.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listening error", "error", err)

		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", "error", err)
		return
	}

	log.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := prettySlog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
