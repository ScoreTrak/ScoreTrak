package orm

type NoRowsAffectedError struct {
	msg string // description of error
}

func (e *NoRowsAffectedError) Error() string { return e.msg }
