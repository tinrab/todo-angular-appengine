package todo

import (
	"crypto/rand"
	"errors"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/urlfetch"
)

type SignInResponse struct {
	UserID       string `json:"userId"`
	SessionToken string `json:"sessionToken"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// Verify ID token provided in header
	token := r.Header.Get("Authorization")
	userID, err := verifyToken(ctx, token)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Invalid ID token", http.StatusBadRequest)
		return
	}

	// Generate a new session token and store it in Memcache
	sessionToken := generateSessionToken()
	if err := memcache.Set(ctx, &memcache.Item{
		Key:        "session:" + sessionToken,
		Value:      []byte(userID),
		Expiration: 10 * time.Hour,
	}); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not start user session", http.StatusInternalServerError)
		return
	}

	// Return session data
	responseJSON(w, SignInResponse{userID, sessionToken})
}

func verifyToken(ctx context.Context, token string) (string, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + token)
	if err != nil {
		return "", err
	}
	var bodyJSON map[string]interface{}
	if err := readJSON(resp.Body, &bodyJSON); err != nil {
		return "", err
	}

	if aud, ok := bodyJSON["aud"].(string); ok {
		if clientID != aud {
			return "", errors.New("Invalid client ID")
		}
	} else {
		return "", errors.New("Invalid ID token")
	}
	if sub, ok := bodyJSON["sub"].(string); ok {
		return sub, nil
	}

	return "", errors.New("Invalid ID token")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-")

func generateSessionToken() string {
	const n = 64
	data := make([]byte, n)
	rand.Read(data)
	token := make([]rune, n)
	for i := range data {
		token[i] = letters[int(data[i])%len(letters)]
	}
	return string(token)
}
