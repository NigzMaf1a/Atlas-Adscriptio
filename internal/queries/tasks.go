package queries

type TaskQueries struct {
	CreateTask       string
	UpdateTaskStatus string
	ReadTasks        string
}

var Task_Queries = TaskQueries{
	CreateTask: `
		INSERT INTO tasks
		(
			registration_id,
			task_detail,
			task_status,
			start_time,
			end_time
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING task_id`,

	UpdateTaskStatus: `
		UPDATE tasks
		SET task_status = $1
		WHERE task_id = $2`,

	ReadTasks: `
		SELECT
			task_id,
			registration_id,
			task_detail,
			task_status,
			start_time,
			end_time
		FROM tasks`,
}
