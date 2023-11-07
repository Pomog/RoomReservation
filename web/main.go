package main

import (
	"fmt"
	"log"
	"net/http"
	"udemyCourse1/pkg/config"
	"udemyCourse1/pkg/handlers"
	"udemyCourse1/pkg/render"
)

const port = ":8080"

func main() {
	// // 1. Create a new instance of the router
	// router := NewRouter()

	// // 2. Launch the server, listening to port 8080
	// router.Run(":8080")

	var app config.AppConfig
	tc, err := render.CreateTemplateCashe()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Server starting on port %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
