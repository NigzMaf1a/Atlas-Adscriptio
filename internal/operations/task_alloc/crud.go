package taskalloc

import (
	"context"
	"database/sql"
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
		log.Fatal("Error while querying the database")
		return err
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
	allocs := []TaskAlloc{}

	rows, err := db.QueryContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		log.Fatal("Error while querying the database")
		return nil, err
	}

	for rows.Next() {
		var t TaskAlloc

		err := rows.Scan(
			t.RegistrationId,
			t.TaskId,
			t.RegistrationId,
		)

		if err != nil {
			log.Fatal("Error occurred while scanning row")
			return nil, err
		}

		allocs = append(allocs, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error occurred while scanning rows")
		return nil, err
	}

	log.Println("Task allocations fetched successfully")

	return allocs, nil
}
