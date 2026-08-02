package contacts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func GetContacts(
	ctx context.Context,
	db *sql.DB,
) ([]Contact, error) {
	contacts := []Contact{}

	rows, err := db.QueryContext(
		ctx,
		``,
	)

	if err != nil {
		fmt.Println("An error occurred while querying about")
		return contacts, err
	}

	defer rows.Close()

	for rows.Next() {
		var c Contact

		err := rows.Scan(
			&c.ContactID,
			&c.Slack,
			&c.Instagram,
			&c.X,
			&c.Facebook,
			&c.Email,
			&c.Phone,
		)

		if err != nil {
			fmt.Println("An error occurred while scanning about")
			return contacts, err
		}

		contacts = append(contacts, c)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("An error occurred while scanning about")
		return contacts, err
	}

	fmt.Println("Contacts fetched successfully")

	return contacts, nil
}

func UpdateContact(
	ctx context.Context,
	db *sql.DB,
	c Contact,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx,
		``,
		&c.Slack,
		&c.Instagram,
		&c.X,
		&c.Facebook,
		&c.Email,
		&c.Phone,
		id,
	)

	if err != nil {
		fmt.Println("An error  occurred while querying for an update")
		return err
	}
	aff, err := result.RowsAffected()

	if err != nil {
		fmt.Println("An error  occurred while checking the rows affected")
		return err
	}

	if aff == 0 {
		fmt.Println("No rows affected")
		return errors.New("No rows affected")
	}

	fmt.Println("Contact updated successfully")

	return nil
}
