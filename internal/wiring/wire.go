package wiring

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/victorsvart/vbi/internal/adapters/endpoints"
	"github.com/victorsvart/vbi/internal/adapters/postgres"
	"github.com/victorsvart/vbi/internal/services"
	"gorm.io/gorm"
)

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
}

var allowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

var allowedHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}

func routerInit() chi.Router {
	chi := chi.NewRouter()
	chi.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		middleware.Timeout(60*time.Second),
	)

	chi.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	return chi
}

func apiWalk(r chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("Routed %s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

func WireApp() chi.Router {
	db := postgres.Connect()
	api := routerInit()
	api.Route("/api/v1", func(r chi.Router) {
		postWire(r, db)
	})

	apiWalk(api)
	return api
}

func postWire(chi chi.Router, db *gorm.DB) services.PostService {
	repo := postgres.NewPostRepository(db)
	services := services.NewPostService(repo)
	endpoints.NewPostHandler(chi, services)
	return services
}
