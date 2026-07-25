package handler

import (
	"database/sql"
	"net/http"

	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/auth"
)

func AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u auth.User
	}
}
