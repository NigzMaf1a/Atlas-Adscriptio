package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	mydb "github.com/NigzMaf1a/Atlas-Adscriptio/internal/db"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/handler"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/middleware"
)

func main() {
	// Database connection
	db, err := mydb.ConnectDB()

	if err != nil {
		log.Fatalf("opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	log.Println("Database connected.")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/login", handler.LoginUser(db))
	mux.HandleFunc("POST /api/reg/post", handler.AddUser(db))
	mux.HandleFunc("GET /api/reg/get", handler.ReadUser(db))
	mux.HandleFunc("PATCH /api/reg/patch/{id}", handler.UpdateUserAccStatus(db))
	mux.HandleFunc("POST /api/sect/post", handler.CreateSector(db))
	mux.HandleFunc("GET /api/sect/get", handler.ReadSectors(db))
	mux.HandleFunc("PATCH /api/sect/patch/{id}", handler.UpdateSectorStatus(db))
	mux.HandleFunc("POST /api/roles/post", handler.CreateRole(db))
	mux.HandleFunc("GET /api/roles/get", handler.ReadRoles(db))
	mux.HandleFunc("PATCH /api/roles/patch/{id}", handler.UpdateRoleStatus(db))
	mux.HandleFunc("POST /api/task/post", handler.CreateTask(db))
	mux.HandleFunc("GET /api/task/get", handler.ReadTasks(db))
	mux.HandleFunc("PATCH /api/task/patch/{id}", handler.UpdateTaskStatus(db))
	mux.HandleFunc("POST /api/alloc/post", handler.CreateTaskAllocation(db))
	mux.HandleFunc("GET /api/alloc/get", handler.ReadTaskAllocations(db))
	mux.HandleFunc("GET /api/about/get", handler.ReadAbout(db))
	mux.HandleFunc("PATCH /api/about/patch", handler.UpdateAbout(db))
	mux.HandleFunc("GET /api/contact/get", handler.ReadContacts(db))
	mux.HandleFunc("PATCH /api/contact/patch", handler.UpdateAbout(db))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Server is running"))
	})

	// Server configuration
	server := &http.Server{
		Addr:              ":" + getPort(),
		Handler:           middleware.CORS(mux),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start server
	go func() {
		log.Printf("Server listening on http://localhost%s", server.Addr)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully.")
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
