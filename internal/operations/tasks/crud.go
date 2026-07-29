package tasks

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateTask(
	ctx context.Context,
	db *sql.DB,
	query string,
	t Task,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		t.RegistrationId,
		t.TaskDetail,
		t.TaskStatus,
		t.StartTime,
		t.EndTime,
	).Scan(&t.TaskId)

	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	log.Println("Task created successfully")
	return nil
}

func UpdateTaskStatus(
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
		return fmt.Errorf("update task status: %w", err)
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	log.Println("Task status updated successfully")
	return nil
}

func ReadTasks(
	ctx context.Context,
	db *sql.DB,
	query string,
	id int64,
) ([]Task, error) {
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task

		if err := rows.Scan(
			&t.TaskId,
			&t.RegistrationId,
			&t.TaskDetail,
			&t.TaskStatus,
			&t.StartTime,
			&t.EndTime,
		); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}

	log.Println("Tasks fetched successfully")

	return tasks, nil
}
