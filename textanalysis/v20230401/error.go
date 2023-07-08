package v20230401

import "fmt"

type TaskError struct {
	Information ErrorInformation
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("task failed: %s", e.Information.Message)
}
