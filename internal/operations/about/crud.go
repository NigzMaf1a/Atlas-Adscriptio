package about

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func GetAbout(
	ctx context.Context,
	db *sql.DB,
) ([]About, error) {
	about := []About{}

	rows, err := db.QueryContext(
		ctx,
		`
			SELECT
    			about_id,
    			about_detail
			FROM about
		`,
	)

	if err != nil {
		fmt.Println("An error occurred while querying about")
		return about, err
	}

	defer rows.Close()

	for rows.Next() {
		var a About

		err := rows.Scan(
			&a.AboutID,
			&a.AboutDetail,
		)

		if err != nil {
			fmt.Println("An error occurred while scanning about")
			return about, err
		}

		about = append(about, a)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("An error occurred while scanning about")
		return about, err
	}

	fmt.Println("About fetched successfully")

	return about, nil
}

func UpdateAbout(
	ctx context.Context,
	db *sql.DB,
	id int64,
	detail string,
) error {
	result, err := db.ExecContext(
		ctx,
		`
		UPDATE about
		SET about_detail = $1
		WHERE about_id = $2
		`,
		detail,
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

	fmt.Println("About updated successfully")

	return nil
}
