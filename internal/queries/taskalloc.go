package queries

type TaskAllocQueries struct {
	CreateAlloc string
	ReadAllocs  string
}

var Alloc_Queries TaskAllocQueries = TaskAllocQueries{
	CreateAlloc: `INSERT INTO task_allocations
				  (task_id,registration_id)
				  VALUES ($1,$2)
	             `,
	ReadAllocs: ``,
}
