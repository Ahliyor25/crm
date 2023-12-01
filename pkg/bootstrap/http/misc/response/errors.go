package response

import "errors"

// Errors used in the project. Sorted alphabetically
var (
	// Default errors
	ErrAccessDenied       = errors.New("access is denied")
	ErrBadRequest         = errors.New("bad request")
	ErrDataNotFound       = errors.New("not found")
	ErrInternalServer     = errors.New("internal server error")
	ErrLimitExceeded      = errors.New("retry limit exceeded")
	ErrNoContent          = errors.New("no content")
	ErrNotImplementation  = errors.New("not implementation")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrSuccess            = errors.New("success")
	ErrTokenIsExpired     = errors.New("token is expired")
	ErrUnauthorized       = errors.New("unauthorized")
)

// Error statuses
var errorCode = map[string]int{
	// Default errors
	ErrAccessDenied.Error():       403,
	ErrBadRequest.Error():         400,
	ErrDataNotFound.Error():       404,
	ErrInternalServer.Error():     500,
	ErrLimitExceeded.Error():      429,
	ErrNoContent.Error():          201,
	ErrNotImplementation.Error():  501,
	ErrSomethingWentWrong.Error(): 500,
	ErrSuccess.Error():            200,
	ErrTokenIsExpired.Error():     401,
	ErrUnauthorized.Error():       401,
}
