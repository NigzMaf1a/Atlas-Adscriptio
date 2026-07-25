package auth

import (
	"context"
	"database/sql"
	"log"
)

func Login(
	ctx context.Context,
	db *sql.DB,
	query string,
	l LoginCred,
) (User, error) {
	var u User

	err := db.QueryRowContext(
		ctx,
		query,
		l.Email,
		l.Password,
	).Scan(
		u.UserId,
		u.SectorId,
		u.RoleId,
		u.UserName,
		u.Email,
		u.Password,
		u.AccStatus,
		u.RegType,
		u.Location,
	)

	if err != nil {
		log.Println("Error occurred while querying the database")
		return u, err
	}

	log.Println("User read successfully")

	return u, nil
}
