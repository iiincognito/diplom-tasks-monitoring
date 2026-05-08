package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HTTPServer struct {
	router *chi.Mux
	cfg    *Config
}

func NewHTTPServer(cfg *Config) *HTTPServer {
	return &HTTPServer{
		router: chi.NewRouter(),
		cfg:    cfg,
	}
}

func (h *HTTPServer) Register(routes ...Route) {
	for _, route := range routes {
		h.router.MethodFunc(route.Method, route.Path, route.Handler)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {

	h.router.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("web")).ServeHTTP(w, r)
	}))

	httpServer := &http.Server{
		Addr:    h.cfg.Addr,
		Handler: h.router,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		err := httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	fmt.Printf("Listening on %s\n", h.cfg.Addr)

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), h.cfg.ShurdownTimeout)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			_ = httpServer.Close()
			return fmt.Errorf("shutting down http server: %w", err)
		}
		log.Println("http server stopped gracefully")
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("http server: %w", err)
		}
	}
	return nil
}
