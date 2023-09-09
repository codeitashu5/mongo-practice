package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"mongoPractice/handler"
	"net/http"
	"time"
)

type Server struct {
	chi.Router
	server *http.Server
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

// SetupRoutes provides all the routes that can be used
func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/mongo", func(mongo chi.Router) {
		mongo.Post("/health", handler.Health)

		// routes for the restaurant
		mongo.Route("/restaurant", func(restaurant chi.Router){
			restaurant.Get("/",handler.AllRestaurant)
			restaurant.Get("/{name}",handler.Restaurant)
		})

		// routes for the movie
		mongo.Route("/movie", func(movie chi.Router){
			movie.Get("/",handler.MovieWithIMDB)
			movie.Get("/cast",handler.MovieWithCast)
		})
	})
	return &Server{Router: router}
}

func (svc *Server) Run(port string) error {
	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.Router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return svc.server.ListenAndServe()
}

func (svc *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return svc.server.Shutdown(ctx)
}
