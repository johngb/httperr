package httperr

import "fmt"

// Err is a custom error type to handle returning errors with GCM
// http response codes
type Err struct {
	statusCode int
	retryAfter int
	err        error
}

// Error implements the Error interface.
func (r *Err) Error() string {
	var res string
	if r.statusCode != 0 {
		res += fmt.Sprintf("%d error", r.statusCode)
	}
	if r.err != nil {
		res += fmt.Sprintf(": %s", r.err.Error())
	}
	if r.retryAfter != 0 {
		res += fmt.Sprintf(", retry-after: %d", r.retryAfter)
	}

	return res
}

// NewHTTPError returns an Err type set up for the basics needed to
// handle HTTP errors.
func New(statusCode, retryAfter int, err error) *Err {
	return &Err{
		statusCode: statusCode,
		retryAfter: retryAfter,
		err:        err,
	}
}

func (e *Err) StatusCode() (statusCode int) {
	return e.statusCode
}

func (e *Err) RetryAfter() (retryAfter int) {
	return e.retryAfter
}
