package exception

type InvariantError struct {
	StatusCode int
	ErrorCode  int
	Name       string
	Message    string
}

func NewInvariantError(errorCode int, message string) InvariantError {
	if errorCode == 0 {
		errorCode = 400
	}
	return InvariantError{400, errorCode, "INVARIANT_ERROR", message}
}

func (e InvariantError) GetStatusCode() int {
	return e.StatusCode
}

func (e InvariantError) GetErrorCode() int {
	return e.ErrorCode
}

func (e InvariantError) GetErrorName() string {
	return e.Name
}

func (e InvariantError) Error() string {
	return e.Message
}
