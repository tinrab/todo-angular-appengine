package todo

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/net/context"
)

var (
	clientID string
)

func init() {
	clientID = os.Getenv("CLIENT_ID")

	r := mux.NewRouter()
	r.HandleFunc("/api/signin", signInHandler).
		Methods("POST")
	r.HandleFunc("/api/hello", authenticate(helloHandler)).
		Methods("GET")

	http.Handle("/api/", cors.AllowAll().Handler(r))
}

func helloHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	responseJSON(w, "Hello, "+userID)
}
