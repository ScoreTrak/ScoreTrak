// Code generated by ent, DO NOT EDIT.

package team

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldPause holds the string denoting the pause field in the database.
	FieldPause = "pause"
	// FieldHidden holds the string denoting the hidden field in the database.
	FieldHidden = "hidden"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeHosts holds the string denoting the hosts edge name in mutations.
	EdgeHosts = "hosts"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// CompetitionTable is the table that holds the competition relation/edge.
	CompetitionTable = "teams"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "competition_id"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "team_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// HostsTable is the table that holds the hosts relation/edge.
	HostsTable = "hosts"
	// HostsInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostsInverseTable = "hosts"
	// HostsColumn is the table column denoting the hosts relation/edge.
	HostsColumn = "team_hosts"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldPause,
	FieldHidden,
	FieldCompetitionID,
	FieldName,
	FieldIndex,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "teams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"competition_teams",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"team_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// Order defines the ordering method for the Team queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByPause orders the results by the pause field.
func ByPause(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPause, opts...).ToFunc()
}

// ByHidden orders the results by the hidden field.
func ByHidden(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldHidden, opts...).ToFunc()
}

// ByCompetitionID orders the results by the competition_id field.
func ByCompetitionID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCompetitionID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByCompetitionField orders the results by competition field.
func ByCompetitionField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompetitionStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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
func newCompetitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompetitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CompetitionTable, CompetitionColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
	)
}
func newHostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HostsTable, HostsColumn),
	)
}
