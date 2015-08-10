package httperr

import "fmt"

// HTTPError is a custom error type to handle returning errors with GCM
// http response codes
type HTTPError struct {
	statusCode int
	retryAfter int
	err        error
}

// Error implements the Error interface.
func (r *HTTPError) Error() string {
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

// NewHTTPError returns an HTTPError type set up for the basics needed to
// handle HTTP errors.
func New(statusCode, retryAfter int, err error) *HTTPError {
	return &HTTPError{
		statusCode: statusCode,
		retryAfter: retryAfter,
		err:        err,
	}
}

func (e *HTTPError) StatusCode() (statusCode int) {
	return e.statusCode
}

func (e *HTTPError) SetRetryAfter() (retryAfter int) {
	return e.retryAfter
}
