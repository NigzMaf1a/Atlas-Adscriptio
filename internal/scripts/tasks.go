package scripts

import (
	"errors"
	"strings"

	"github.com/NigzMaf1a/Atlas-Adscriptio/internal/operations/tasks"
)

func ValidateStartTime(t1, t2 tasks.Task) error {
	if strings.EqualFold(t1.TaskStatus, "Done") {
		return nil
	}

	if (t2.StartTime.Equal(t1.StartTime) || t2.StartTime.After(t1.StartTime)) &&
		(t2.StartTime.Equal(t1.EndTime) || t2.StartTime.Before(t1.EndTime)) {

		return errors.New("task start time conflicts with an active task")
	}

	return nil
}

func CheckOverlap(tasks []tasks.Task, newTask tasks.Task) error {
	for _, v := range tasks {
		if err := ValidateStartTime(v, newTask); err != nil {
			return WrapError(
				"Error while checking task overlap",
				err,
			)
		}
	}
	return nil
}
