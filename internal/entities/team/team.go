// Code generated by ent, DO NOT EDIT.

package team

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldPause holds the string denoting the pause field in the database.
	FieldPause = "pause"
	// FieldHidden holds the string denoting the hidden field in the database.
	FieldHidden = "hidden"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// EdgeHosts holds the string denoting the hosts edge name in mutations.
	EdgeHosts = "hosts"
	// EdgeHostservices holds the string denoting the hostservices edge name in mutations.
	EdgeHostservices = "hostservices"
	// EdgeChecks holds the string denoting the checks edge name in mutations.
	EdgeChecks = "checks"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// HostsTable is the table that holds the hosts relation/edge.
	HostsTable = "hosts"
	// HostsInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostsInverseTable = "hosts"
	// HostsColumn is the table column denoting the hosts relation/edge.
	HostsColumn = "team_id"
	// HostservicesTable is the table that holds the hostservices relation/edge.
	HostservicesTable = "host_services"
	// HostservicesInverseTable is the table name for the HostService entity.
	// It exists in this package in order to avoid circular dependency with the "hostservice" package.
	HostservicesInverseTable = "host_services"
	// HostservicesColumn is the table column denoting the hostservices relation/edge.
	HostservicesColumn = "team_id"
	// ChecksTable is the table that holds the checks relation/edge.
	ChecksTable = "checks"
	// ChecksInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	ChecksInverseTable = "checks"
	// ChecksColumn is the table column denoting the checks relation/edge.
	ChecksColumn = "team_id"
	// PropertiesTable is the table that holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "team_id"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "teams"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "competition_id"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDisplayName,
	FieldPause,
	FieldHidden,
	FieldNumber,
	FieldCompetitionID,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultHidden holds the default value on creation for the "hidden" field.
	DefaultHidden bool
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Order defines the ordering method for the Team queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByPause orders the results by the pause field.
func ByPause(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPause, opts...).ToFunc()
}

// ByHidden orders the results by the hidden field.
func ByHidden(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldHidden, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByCompetitionID orders the results by the competition_id field.
func ByCompetitionID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCompetitionID, opts...).ToFunc()
}

// ByHostsCount orders the results by hosts count.
func ByHostsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHostsStep(), opts...)
	}
}

// ByHosts orders the results by hosts terms.
func ByHosts(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHostservicesCount orders the results by hostservices count.
func ByHostservicesCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHostservicesStep(), opts...)
	}
}

// ByHostservices orders the results by hostservices terms.
func ByHostservices(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostservicesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByPropertiesCount orders the results by properties count.
func ByPropertiesCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPropertiesStep(), opts...)
	}
}

// ByProperties orders the results by properties terms.
func ByProperties(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPropertiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}
func newHostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HostsTable, HostsColumn),
	)
}
func newHostservicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostservicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HostservicesTable, HostservicesColumn),
	)
}
func newChecksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChecksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChecksTable, ChecksColumn),
	)
}
func newPropertiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PropertiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PropertiesTable, PropertiesColumn),
	)
}
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompetitionTable, CompetitionColumn),
	)
}
