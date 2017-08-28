package todo

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

type AuthenticatedHandler func(context.Context, http.ResponseWriter, *http.Request, string)

func authenticate(handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		sessionToken := r.Header.Get("Authorization")

		if len(sessionToken) == 0 {
			responseError(w, "Invalid session token", http.StatusUnauthorized)
			return
		}

		sessionItem, err := memcache.Get(ctx, "session:"+sessionToken)
		if err != nil {
			responseError(w, "Could not authenticate", http.StatusUnauthorized)
			return
		}

		userID := string(sessionItem.Value)
		handler(ctx, w, r, userID)
	}
}
