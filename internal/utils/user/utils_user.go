package registration

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"errors"
)

var ErrUserNotFound = errors.New("user not found or no changes made")

func (user *models.User) UpdateUserStructProps(u models.User) {
	user.SectorID = u.SectorID
	user.RoleID = u.RoleID
	user.UserName = u.UserName
	user.Email = u.Email
	user.Password = u.Password
	user.AccStatus = u.AccStatus
	user.RegStatus = u.RegStatus
	user.Location = u.Location
}


func CreateUser(ctx context.Context, db *sql.DB, query string, u *models.User) error {

	err := db.QueryRowContext(
		ctx,
		query,
		u.SectorID,
		u.RoleID,
		u.UserName,
		u.Email,
		u.Password,
		u.AccStatus,
		u.RegStatus,
		u.Location,
	).Scan(&u.UserID)

	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return err
	}

	log.Println("User successfully created with ID:", u.UserID)
	return nil
}

func ReadUsers(ctx context.Context, db *sql.DB, query string) ([]models.User, error) {
	users := []models.User{}

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Printf("Error while reading users: %v\n", err)
		return nil, err 
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.UserID,
			&u.SectorID,
			&u.RoleID,
			&u.UserName,
			&u.Email,
			&u.Password,
			&u.AccStatus,
			&u.RegStatus,
			&u.Location,
		)

		if err != nil {
			log.Printf("Error while scanning record: %v\n", err)
			return nil, err 
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after finishing row iteration: %v\n", err)
		return nil, err
	}

	log.Println("User records fetched successfully")
	return users, nil
}

func (u *models.User) UpdateUser(ctx context.Context, db *sql.DB, query string, user *models.User) error {
	
	u.UpdateUserStructProps(user)

	id, err := strconv.ParseInt(u.UserID, 10, 64)
	if err != nil {
		log.Printf("Error occurred while parsing id: %v\n", err)
		return err
	}

	result, err := db.ExecContext(
		ctx,
		query,
		u.SectorID,
		u.RoleID,
		u.UserName,
		u.Email,
		u.Password,
		u.AccStatus,
		u.RegStatus,
		u.Location,
		id, 
	)
	if err != nil {
		log.Printf("Error occurred while querying the database: %v\n", err)
		return err
	}

	
	aff, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v\n", err)
		return err
	}

	if aff == 0 {
		log.Printf("Update failed: user with ID %d not found\n", id)
		return ErrUserNotFound
	}

	log.Printf("Successfully updated user: %+v\n", u)
	return nil
}