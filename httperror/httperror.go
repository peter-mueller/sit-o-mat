package httperror

import (
	"fmt"
	"net/http"

	"gocloud.dev/gcerrors"
)

// Error is the default error type of the sit-o-mat application
type Error struct {
	Status  int
	Message string
	Err     error `json:"-"`
}

// Wrap an existing error into an httperror and add an description what failed
func Wrap(message string, err error) Error {
	return Error{
		Status:  Parse(err),
		Err:     fmt.Errorf("%v: %v", message, err),
		Message: fmt.Sprintf("%v: %v", message, err),
	}
}

// Parse an error to an status code
func Parse(err error) int {
	code := gcerrors.Code(err)
	if code == gcerrors.AlreadyExists {
		return http.StatusConflict
	}
	if code == gcerrors.NotFound {
		return http.StatusNotFound
	}
	if code == gcerrors.PermissionDenied {
		return http.StatusForbidden
	}
	if code == gcerrors.InvalidArgument {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func (e Error) Error() string {
	return e.Err.Error()
}

// Unwrap the contained error
func (e Error) Unwrap() error {
	return e.Err
}
