package roles

import (
	"context"
	"database/sql"
	"fmt"
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
		return fmt.Errorf("create role: %w", err)
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
		return fmt.Errorf("update role status: %w", err)
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	log.Println("Role status updated successfully")
	return nil
}

func ReadRoles(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Role, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query roles: %w", err)
	}
	defer rows.Close()

	var roles []Role

	for rows.Next() {
		var r Role

		if err := rows.Scan(
			&r.RoleId,
			&r.SectorId,
			&r.RoleTitle,
			&r.RoleStatus,
		); err != nil {
			return nil, fmt.Errorf("scan role: %w", err)
		}

		roles = append(roles, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate roles: %w", err)
	}

	log.Println("Roles fetched successfully")

	return roles, nil
}
