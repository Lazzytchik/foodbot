package errors

type ServiceError struct {
	RealError error
	ShowError error

	handler ErrorHandler
}

func (e ServiceError) Error() string {
	e.handler(e.RealError)

	return e.ShowError.Error()
}
