package queries

type TaskAllocQueries struct {
	CreateAlloc string
	ReadAllocs  string
}

var Alloc_Queries = TaskAllocQueries{
	CreateAlloc: `
		INSERT INTO task_allocations
		(
			task_id,
			registration_id
		)
		VALUES ($1,$2)
		RETURNING task_allocation_id`,

	ReadAllocs: `
		SELECT
			task_allocation_id,
			task_id,
			registration_id
		FROM task_allocations`,
}
