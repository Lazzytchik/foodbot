package errors

type ErrorHandler func(error)

type Factory struct {
	Handler ErrorHandler
}

func (f *Factory) Error(real, show error) ServiceError {
	return ServiceError{
		RealError: real,
		ShowError: show,
		handler:   f.Handler,
	}
}
