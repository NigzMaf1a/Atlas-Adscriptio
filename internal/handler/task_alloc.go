package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	taskalloc "github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/task_alloc"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/queries"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/scripts"
)

func CreateTaskAllocation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		task, err := scripts.DecodeJSON[taskalloc.TaskAlloc](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if err := taskalloc.CreateTaskAllocation(ctx, db, queries.Task_Queries.CreateTask, task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := scripts.EncodeJSON(w, task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
