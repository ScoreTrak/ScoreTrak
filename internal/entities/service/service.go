// Code generated by ent, DO NOT EDIT.

package service

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/resolver"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
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
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// EdgeHostservices holds the string denoting the hostservices edge name in mutations.
	EdgeHostservices = "hostservices"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// Table holds the table name of the service in the database.
	Table = "services"
	// HostservicesTable is the table that holds the hostservices relation/edge.
	HostservicesTable = "host_services"
	// HostservicesInverseTable is the table name for the HostService entity.
	// It exists in this package in order to avoid circular dependency with the "hostservice" package.
	HostservicesInverseTable = "host_services"
	// HostservicesColumn is the table column denoting the hostservices relation/edge.
	HostservicesColumn = "service_id"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "services"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "competition_id"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDisplayName,
	FieldPause,
	FieldHidden,
	FieldType,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type resolver.Service) error {
	switch _type {
	case "ftp", "ssh", "winrm", "ping", "http", "ldap", "dns", "smb", "imap", "sql", "caldav":
		return nil
	default:
		return fmt.Errorf("service: invalid enum value for type field: %q", _type)
	}
}

// Order defines the ordering method for the Service queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByCompetitionID orders the results by the competition_id field.
func ByCompetitionID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCompetitionID, opts...).ToFunc()
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

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}
func newHostservicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostservicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HostservicesTable, HostservicesColumn),
	)
}
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompetitionTable, CompetitionColumn),
	)
}