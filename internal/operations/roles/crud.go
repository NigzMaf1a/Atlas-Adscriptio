package roles

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func CreateRole(
	ctx context.Context,
	db *sql.DB,
	query string,
	r Role,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		r.SectorId,
		r.RoleTitle,
		r.RoleStatus,
	).Scan(&r.RoleId)

	if err != nil {
		log.Fatal("Error occurred while querying database")
		return err
	}

	log.Println("Role created successfully")
	return nil
}

func UpdateRoleStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		log.Fatal("Error occurred while querying the database")
		return err
	}

	aff, err := result.RowsAffected()

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

	log.Println("Role status updated successfully")
	return nil
}

func ReadRoles(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Role, error) {
	roles := []Role{}

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
		var r Role

		err := rows.Scan(
			r.RoleId,
			r.SectorId,
			r.RoleTitle,
			r.RoleStatus,
		)

		if err != nil {
			log.Fatal("Error occurred while scanning row")
			return nil, err
		}

		roles = append(roles, r)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error occurred while scanning rows")
		return nil, err
	}

	log.Println("Roles fetched successfully")

	return roles, nil
}
