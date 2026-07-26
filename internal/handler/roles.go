package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/roles"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/queries"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/scripts"
)

func CreateRole(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		role, err := scripts.DecodeJSON[roles.Role](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if err := roles.CreateRole(ctx, db, queries.Role_Queries.CreateRole, role); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := scripts.EncodeJSON(w, role); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ReadRoles(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		roles, err := roles.ReadRoles(ctx, db, queries.Role_Queries.ReadRoles)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := scripts.EncodeJSON(w, roles); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateRoleStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
