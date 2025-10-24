package exception

type ApplicationError interface {
	GetStatusCode() int
	GetErrorCode() int
	GetErrorName() string
	Error() string
}
