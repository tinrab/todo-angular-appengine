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

func createTodoHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	// Read todo from request body
	var todo Todo
	if err := readJSON(r.Body, &todo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not read todo", http.StatusBadRequest)
		return
	}
	todo.UserID = userID
	todo.CreatedAt = time.Now()

	// Store todo
	key := datastore.NewIncompleteKey(ctx, "Todo", nil)
	if key, err := datastore.Put(ctx, key, &todo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not create todo", http.StatusInternalServerError)
	} else {
		todo.ID = strconv.FormatInt(key.IntID(), 10)
		responseJSON(w, todo)
	}
}

func listTodosHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	var todos []Todo
	// Query todos by user's ID and order them by creation time
	query := datastore.NewQuery("Todo").
		Filter("UserID =", userID).
		Order("-CreatedAt")

	// Execute query
	if keys, err := query.GetAll(ctx, &todos); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not read todos", http.StatusInternalServerError)
	} else {
		// Return empty array instead of 'null'
		if len(todos) == 0 {
			responseJSON(w, []Todo{})
			return
		}
		// Set string IDs
		for i := range todos {
			todos[i].ID = strconv.FormatInt(keys[i].IntID(), 10)
		}
		responseJSON(w, todos)
	}
}

func updateTodoHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	// Parse ID
	id := mux.Vars(r)["id"]
	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responseError(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Get old todo
	var todo Todo
	if err := getOwningTodo(ctx, userID, todoID, &todo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not read old todo", http.StatusBadRequest)
		return
	}

	// Read new todo from request body
	var newTodo Todo
	if err := readJSON(r.Body, &newTodo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	// Update todo
	todo.Title = newTodo.Title
	key := datastore.NewKey(ctx, "Todo", "", todoID, nil)
	if _, err := datastore.Put(ctx, key, &todo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not update todo", http.StatusInternalServerError)
		return
	}

	todo.ID = id
	responseJSON(w, todo)
}

func deleteTodoHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	// Parse ID
	id := mux.Vars(r)["id"]
	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responseError(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// Get todo to check if it can be deleted
	var todo Todo
	if err := getOwningTodo(ctx, userID, todoID, &todo); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not read todo", http.StatusInternalServerError)
		return
	}

	// Delete todo
	key := datastore.NewKey(ctx, "Todo", "", todoID, nil)
	if err := datastore.Delete(ctx, key); err != nil {
		log.Errorf(ctx, "%v", err)
		responseError(w, "Could not delete todo", http.StatusInternalServerError)
		return
	}

	todo.ID = id
	responseJSON(w, todo)
}

func getOwningTodo(ctx context.Context, userID string, id int64, todo *Todo) error {
	// Fetch todo
	key := datastore.NewKey(ctx, "Todo", "", id, nil)
	if err := datastore.Get(ctx, key, todo); err != nil {
		return err
	}
	// Check if it belongs to the current user
	if todo.UserID != userID {
		return errors.New("Not own todo")
	}
	return nil
}
