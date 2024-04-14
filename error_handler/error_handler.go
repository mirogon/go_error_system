package error_handler

import es "github.com/mirogon/go_error_system"

type ErrorHandler interface {
	ProcessError(e es.Error) error
}

var ErrHandler ErrorHandler = ErrorHandlerImpl{}
