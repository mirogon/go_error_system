package error_handler

import (
	"log"

	es "github.com/mirogon/go_error_system"
)

type ErrorHandlerImpl struct{}

func (errorHandler ErrorHandlerImpl) ProcessError(e es.Error) error {
	log.Printf("Error: %s\n", e.Error())
	return e
}
