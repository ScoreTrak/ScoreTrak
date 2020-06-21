package orm

type NoRowsAffected struct {
	msg string // description of error
}

func (e *NoRowsAffected) Error() string { return e.msg }
