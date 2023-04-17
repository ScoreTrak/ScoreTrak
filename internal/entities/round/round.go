// Code generated by ent, DO NOT EDIT.

package round

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the round type in the database.
	Label = "round"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// FieldRoundNumber holds the string denoting the round_number field in the database.
	FieldRoundNumber = "round_number"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldErr holds the string denoting the err field in the database.
	FieldErr = "err"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// EdgeChecks holds the string denoting the checks edge name in mutations.
	EdgeChecks = "checks"
	// Table holds the table name of the round in the database.
	Table = "rounds"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "rounds"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "competition_id"
	// ChecksTable is the table that holds the checks relation/edge.
	ChecksTable = "checks"
	// ChecksInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	ChecksInverseTable = "checks"
	// ChecksColumn is the table column denoting the checks relation/edge.
	ChecksColumn = "round_checks"
)

// Columns holds all SQL columns for round fields.
var Columns = []string{
	FieldID,
	FieldCompetitionID,
	FieldRoundNumber,
	FieldNote,
	FieldErr,
	FieldStartedAt,
	FieldFinishedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Order defines the ordering method for the Round queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCompetitionID orders the results by the competition_id field.
func ByCompetitionID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCompetitionID, opts...).ToFunc()
}

// ByRoundNumber orders the results by the round_number field.
func ByRoundNumber(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldRoundNumber, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByErr orders the results by the err field.
func ByErr(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldErr, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}

// ByChecksCount orders the results by checks count.
func ByChecksCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChecksStep(), opts...)
	}
}

// ByChecks orders the results by checks terms.
func ByChecks(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChecksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CompetitionTable, CompetitionColumn),
	)
}
func newChecksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChecksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChecksTable, ChecksColumn),
	)
}
