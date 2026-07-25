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
) error {
	result, err := db.ExecContext(
		ctx,
		query,
		status,
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
