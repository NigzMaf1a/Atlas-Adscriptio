package taskalloc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateTaskAllocation(
	ctx context.Context,
	db *sql.DB,
	query string,
	t TaskAlloc,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		t.TaskId,
		t.RegistrationId,
	).Scan(&t.TaskAllocId)

	if err != nil {
		return fmt.Errorf("create task allocation: %w", err)
	}

	log.Println("Task allocation created successfully")
	return nil
}

func ReadTaskAllocations(
	ctx context.Context,
	db *sql.DB,
	query string,
	id int64,
) ([]TaskAlloc, error) {
	rows, err := db.QueryContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("query task allocations: %w", err)
	}
	defer rows.Close()

	var allocs []TaskAlloc

	for rows.Next() {
		var t TaskAlloc

		if err := rows.Scan(
			&t.TaskAllocId,
			&t.TaskId,
			&t.RegistrationId,
		); err != nil {
			return nil, fmt.Errorf("scan task allocation: %w", err)
		}

		allocs = append(allocs, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate task allocations: %w", err)
	}

	log.Println("Task allocations fetched successfully")

	return allocs, nil
}
