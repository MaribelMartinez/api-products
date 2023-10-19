package errors

import "fmt"

type ErrorFromMessage struct {
	Message    string
	StatusCode int
}

func (e ErrorFromMessage) Error() string {
	return fmt.Sprintf("error: %s - statusCode %d", e.Message, e.StatusCode)
}
