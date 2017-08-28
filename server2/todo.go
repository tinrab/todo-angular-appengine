package todo

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Todo struct {
	ID        string    `json:"id" datastore:"-"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

func create(w http.ResponseWriter, r *http.Request, ctx context.Context, userID string) {
	body, err := readJSON(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var title string
	var ok bool
	if title, ok = body["title"].(string); !ok {
		responseError(w, "invalid json", http.StatusBadRequest)
		return
	}

	todo := &Todo{
		UserID:    userID,
		Title:     title,
		CreatedAt: time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Todo", nil)
	if newKey, err := datastore.Put(ctx, key, todo); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		todo.ID = strconv.FormatInt(newKey.IntID(), 10)
		responseJSON(w, todo)
	}
}

func getAll(w http.ResponseWriter, r *http.Request, ctx context.Context, userID string) {
	var todos []Todo
	if keys, err := datastore.NewQuery("Todo").
		Filter("UserID =", userID).
		Order("-CreatedAt").
		GetAll(ctx, &todos); err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
	} else {
		if len(todos) == 0 {
			responseJSON(w, []Todo{})
			return
		}

		for i := range todos {
			todos[i].ID = strconv.FormatInt(keys[i].IntID(), 10)
		}
		responseJSON(w, todos)
	}
}

func update(w http.ResponseWriter, r *http.Request, ctx context.Context, userID string) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := readJSON(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var title string
	var ok bool
	if title, ok = body["title"].(string); !ok || len(title) == 0 {
		responseError(w, "invalid json", http.StatusBadRequest)
		return
	}

	if todo, key, err := getOwnTodo(ctx, userID, id); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
	} else {
		todo.Title = title
		if _, err := datastore.Put(ctx, key, todo); err != nil {
			responseError(w, "could not update", http.StatusInternalServerError)
			return
		}

		responseJSON(w, todo)
	}
}

func delete(w http.ResponseWriter, r *http.Request, ctx context.Context, userID string) {
	vars := mux.Vars(r)
	id := vars["id"]

	if todo, key, err := getOwnTodo(ctx, userID, id); err == nil {
		if err := datastore.Delete(ctx, key); err != nil {
			responseError(w, "could not delete", http.StatusBadRequest)
			return
		}
		responseJSON(w, todo)
	} else {
		responseError(w, err.Error(), http.StatusBadRequest)
	}
}

func getOwnTodo(ctx context.Context, userID string, strID string) (*Todo, *datastore.Key, error) {
	id, err := strconv.ParseInt(strID, 10, 64)
	log.Infof(ctx, "%d\n", id)

	if err != nil {
		return nil, nil, errors.New("invalid id")
	}

	key := datastore.NewKey(ctx, "Todo", "", id, nil)
	todo := &Todo{}

	if err := datastore.Get(ctx, key, todo); err != nil {
		return nil, nil, errors.New("todo not found")
	}

	if todo.UserID != userID {
		return nil, nil, errors.New("not own todo")
	}

	todo.ID = strID
	return todo, key, nil
}
