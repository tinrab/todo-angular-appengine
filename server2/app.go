package todo

import (
	"errors"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/urlfetch"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	ID           string `json:"id"`
	SessionToken string `json:"sessionToken"`
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

type TodoHandler func(http.ResponseWriter, *http.Request, context.Context, string)

var (
	clientID string
)

func init() {
	clientID = os.Getenv("CLIENT_ID")

	r := mux.NewRouter()

	r.HandleFunc("/api/signin/{token}", signin).Methods("GET")
	r.HandleFunc("/api/todos", authenticate(create)).Methods("POST")
	r.HandleFunc("/api/todos", authenticate(getAll)).Methods("GET")
	r.HandleFunc("/api/todos/{id}", authenticate(update)).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", authenticate(delete)).Methods("DELETE")

	http.Handle("/api/", cors.AllowAll().Handler(r))
}

func authenticate(handler TodoHandler) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		sessionToken := r.Header.Get("X-Session-Token")
		ctx := appengine.NewContext(r)

		if len(userID) == 0 || len(sessionToken) == 0 {
			responseError(w, "Unauthenticated", http.StatusBadRequest)
			return
		}

		session, err := memcache.Get(ctx, userID)
		if err != nil || string(session.Value) != sessionToken {
			responseError(w, "Unauthenticated", http.StatusBadRequest)
			return
		}

		handler(w, r, ctx, userID)
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	googleToken := vars["token"]
	ctx := appengine.NewContext(r)

	userID, err := verifyToken(ctx, googleToken)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := generateToken()

	session := &memcache.Item{
		Key:        userID,
		Value:      []byte(token),
		Expiration: 1 * time.Hour,
	}
	if err := memcache.Set(ctx, session); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON(w, User{userID, token})
}

const tokenInfoURL = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="

func verifyToken(ctx context.Context, token string) (string, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get(tokenInfoURL + token)
	if err != nil {
		return "nil", err
	}

	bodyJSON, err := readJSON(resp.Body)
	if err != nil {
		return "", err
	}

	if aud, ok := bodyJSON["aud"].(string); ok {
		if clientID != aud {
			return "", errors.New("invalid client id")
		}
	} else {
		return "", errors.New("invalid id token")
	}

	if sub, ok := bodyJSON["sub"].(string); ok {
		return sub, nil
	}

	return "", errors.New("invalid id token")
}
