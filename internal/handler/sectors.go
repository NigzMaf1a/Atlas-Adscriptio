package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/sectors"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/queries"
	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/scripts"
)

func CreateSector(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		sector, err := scripts.DecodeJSON[sectors.Sector](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if err := sectors.CreateSector(ctx, db, queries.Sector_Queries.CreateSector, sector); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := scripts.EncodeJSON(w, sector); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ReadSectors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		sectors, err := sectors.ReadSectors(ctx, db, queries.Sector_Queries.ReadSectors)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := scripts.EncodeJSON(w, sectors); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateSectorStatus(db *sql.DB) http.HandlerFunc {
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

		if err := sectors.UpdateSectorStatus(ctx, db, queries.Role_Queries.UpdateRoleStatus, status, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
