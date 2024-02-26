package executors

import "errors"

var (
	outcomeDefaultSetError  = errors.New("Unable to apply defaults to outcome")
	propertyUnmarshallError = errors.New("Unable to unmarshall properties")
)
