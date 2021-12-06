package web

import (
	"github.com/tonnytg/viewerlight/pkg/web/routes"
	"net/http"
	"os"
)

// StartServer listen on port 8080 or custom if PORT exported
// use routes .go to define routes
func StartServer() {

	// import middleware routes
	routes.CallRoutes()
	port := "8080"
	// if PORT is exported use it
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	// if PORT is not exported use 8080 as default
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
