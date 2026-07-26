package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/tasks"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/queries"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/scripts"
)

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := scripts.ConvertToInteger(r.PathValue("id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := scripts.DecodeJSON[tasks.Task](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		tasklist, err := tasks.ReadTasks(ctx, db, queries.Task_Queries.ReadTasks, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := scripts.CheckOverlap(tasklist, task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tasks.CreateTask(ctx, db, queries.Task_Queries.CreateTask, task); err != nil {
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

func ReadTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		id, err := scripts.ConvertToInteger(r.PathValue("id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		tasks, err := tasks.ReadTasks(ctx, db, queries.Task_Queries.ReadTasks, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := scripts.EncodeJSON(w, tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateTaskStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := scripts.ConvertToInteger(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		status, err := scripts.DecodeJSON[string](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if err := tasks.UpdateTaskStatus(ctx, db, queries.Task_Queries.UpdateTaskStatus, status, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
