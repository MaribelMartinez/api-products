package errors

import "fmt"

type ErrorFromMessage struct {
	Message string
}

func (e ErrorFromMessage) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}
