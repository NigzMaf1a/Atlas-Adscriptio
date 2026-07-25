package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func CreateUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	u User,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		u.SectorId,
		u.RoleId,
		u.UserName,
		u.Email,
		u.Password,
		u.AccStatus,
		u.RegType,
		u.Location,
	).Scan(&u.UserId)

	if err != nil {
		log.Fatal("Error while querying the database")
		return err
	}

	log.Println("User created successfully")

	return nil
}

func ReadUsers(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]User, error) {
	users := []User{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		log.Fatal("Error while querying the database")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(
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
			log.Fatal("Error occurred while scanning row")
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error while scanning rows")
		return nil, err
	}

	log.Println("Users fetched successfully")
	return users, nil
}

func UpdateAccStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {
	res, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		log.Fatal("Error occurred while querying the database")
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		log.Fatal("Error occurred while updating row")
		return err
	}

	if aff == 0 {
		log.Fatal(("No rows affected"))
		return errors.New("Record not found for update")
	}

	log.Println("Account status updated successfully")

	return nil
}

func UpdateUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	user User,
) error {
	res, err := db.ExecContext(
		ctx,
		query,
		user.Password,
		user.Location,
		user.UserId,
	)

	if err != nil {
		log.Fatal("Error occurred while querying the database")
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		log.Fatal("Error occurred while updating row")
		return err
	}

	if aff == 0 {
		log.Fatal(("No rows affected"))
		return errors.New("Record not found for update")
	}

	log.Println("Account updated successfully")

	return nil
}
