package todo

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

type AuthenticatedHandler func(context.Context, http.ResponseWriter, *http.Request, string)

func authenticate(handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		// Get session token from header
		sessionToken := r.Header.Get("Authorization")
		if len(sessionToken) == 0 {
			responseError(w, "Invalid session token", http.StatusUnauthorized)
			return
		}

		// Fetch user's ID from Memcache
		sessionItem, err := memcache.Get(ctx, "session:"+sessionToken)
		if err != nil {
			log.Errorf(ctx, "%v", err)
			responseError(w, "Could not authenticate", http.StatusUnauthorized)
			return
		}

		// Call handler function
		userID := string(sessionItem.Value)
		handler(ctx, w, r, userID)
	}
}
