package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func SetupWeb(controllers ...HttpController) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Use(middleware.Heartbeat("/healthz"))

	for _, resource := range controllers {
		r.Mount(resource.GetPath(), resource.GetRoutes())
	}

	return http.ListenAndServe(":8888", r)
}

type HttpController interface {
	GetPath() string
	GetRoutes() http.Handler
}
