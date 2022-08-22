package util

import "errors"

var ErrSkippedOperation = errors.New("this operation is not supported for unspecified platform. If you are using supported platform like kubernetes, or docker, make sure to specify it in scoretrak config")
