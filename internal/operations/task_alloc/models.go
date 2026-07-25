package taskalloc

type TaskAlloc struct {
	TaskAllocId    int `json:"task_allocation_id"`
	TaskId         int `json:"task_id"`
	RegistrationId int `json:"registration_id"`
}
