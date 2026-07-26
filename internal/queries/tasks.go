package queries

type TaskQueries struct {
	CreateTask       string
	UpdateTaskStatus string
	ReadTasks        string
}

var Task_Queries TaskQueries = TaskQueries{
	CreateTask: `INSERT INTO tasks
				 (registration_id,task_detail,task_status,start_time,end_time)
				 VALUES ($1,$2,$3,$4,$5)
	           `,
	UpdateTaskStatus: `UPDATE tasks SET task_status = $1 WHERE id = $2`,
	ReadTasks:        ``,
}
