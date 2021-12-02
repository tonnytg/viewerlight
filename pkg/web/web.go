package web

import (
	"github.com/tonnytg/viewerlight/pkg/web/routes"
	"net/http"
)

// StartServer listen on port 8080 or custom if PORT exported
// use routes .go to define routes
func StartServer() {
	routes.CallRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
