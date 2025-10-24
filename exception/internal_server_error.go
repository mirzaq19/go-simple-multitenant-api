package exception

type InternalServerError struct {
	StatusCode int
	ErrorCode  int
	Name       string
	Message    string
}

func NewInternalServerError(errorCode int, message string) InternalServerError {
	if errorCode == 0 {
		errorCode = 500
	}
	return InternalServerError{500, errorCode, "INTERNAL_SERVER_ERROR", message}
}

func (e InternalServerError) GetStatusCode() int {
	return e.StatusCode
}

func (e InternalServerError) GetErrorCode() int {
	return e.ErrorCode
}

func (e InternalServerError) GetErrorName() string {
	return e.Name
}

func (e InternalServerError) Error() string {
	return e.Message
}
