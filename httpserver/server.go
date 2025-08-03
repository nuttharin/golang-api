package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type ginRestAPI struct {
	server *http.Server
	Host   string
	Port   int
	Router Router
}

func NewRestAPI(host string, port int, router *ginRouter) *ginRestAPI {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router.engine,
	}
	return &ginRestAPI{Host: host, Port: port, server: server}
}

// Start server
func (r *ginRestAPI) Start() error {
	fmt.Printf("Starting server on %s:%d\n", r.Host, r.Port)

	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	if err := r.gracefulShutdown(); err != nil {
		return err
	}

	return nil
}

// Manage Graceful Shutdown
func (r *ginRestAPI) gracefulShutdown() error {
	//  Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	fmt.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown:: %v\n", err)
		return err
	}
	fmt.Println("Server stopped.")
	return nil
}
