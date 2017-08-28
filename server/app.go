package todo

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	clientID string
)

func init() {
	// Read configuration environment variables
	clientID = os.Getenv("CLIENT_ID")

	// Register routes
	r := mux.NewRouter()
	r.HandleFunc("/api/signin", signInHandler).
		Methods("POST")
	r.HandleFunc("/api/todos", authenticate(createTodoHandler)).
		Methods("POST")
	r.HandleFunc("/api/todos", authenticate(listTodosHandler)).
		Methods("GET")
	r.HandleFunc("/api/todos/{id}", authenticate(updateTodoHandler)).
		Methods("POST")
	r.HandleFunc("/api/todos/{id}", authenticate(deleteTodoHandler)).
		Methods("DELETE")

	// Start HTTP server
	http.Handle("/", cors.AllowAll().Handler(r))
}
