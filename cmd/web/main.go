package main

import (
	"apistructure/handlers"
	"apistructure/internal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

/*
func Routes(config *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)
	todoHandler := handlers.NewTodoHandler(config.Database)
	router.Route("/v1", func(r chi.Router) {
		r.Get("/todos", todoHandler.ListAllTodos)
		r.Get("/todos/{todoId:int}", todoHandler.FetchTodoByID)
	})

	return router
}*/

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := handlers.Routes(configuration)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Println("Serving application at PORT :" + configuration.Constants.PORT)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.PORT, router)) // Note, the port is usually gotten from the environment.
}
