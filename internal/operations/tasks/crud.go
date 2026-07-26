package tasks

import (
	"context"
	"database/sql"
	"errors"
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
		fmt.Println("Error occurred while creating task")
		log.Fatal(err)
		return err
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
		log.Println("Error occurred while updating task status")
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("No rows updated")
			return err
		}
		log.Println("Error occurred while updating task status")
		return err
	}

	if aff == 0 {
		log.Println("No rows affected")
		return errors.New("No records found")
	}

	return nil
}

func ReadTasks(ctx context.Context, db *sql.DB, query string, id int64) ([]Task, error) {
	tasks := []Task{}
	rows, err := db.QueryContext(ctx, query, id)

	if err != nil {
		log.Println("Error occurred while reading tasks")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t Task

		err := rows.Scan(
			t.TaskId,
			t.RegistrationId,
			t.TaskDetail,
			t.TaskStatus,
			t.StartTime,
			t.EndTime,
		)

		if err != nil {
			log.Println("Error while scanning rows")
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error while scanning rows")
		return nil, err
	}

	return tasks, nil
}
