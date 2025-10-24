package exception

type NotFoundError struct {
	StatusCode int
	ErrorCode  int
	Name       string
	Message    string
}

func NewNotFoundError(errorCode int, message string) NotFoundError {
	if errorCode == 0 {
		errorCode = 404
	}
	return NotFoundError{404, errorCode, "NOT_FOUND_ERROR", message}
}

func (e NotFoundError) GetStatusCode() int {
	return e.StatusCode
}

func (e NotFoundError) GetErrorCode() int {
	return e.ErrorCode
}

func (e NotFoundError) GetErrorName() string {
	return e.Name
}

func (e NotFoundError) Error() string {
	return e.Message
}
