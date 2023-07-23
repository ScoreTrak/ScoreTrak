// Code generated by ent, DO NOT EDIT.

package hostservice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the hostservice type in the database.
	Label = "host_service"
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
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldPointBoost holds the string denoting the point_boost field in the database.
	FieldPointBoost = "point_boost"
	// FieldRoundUnits holds the string denoting the round_units field in the database.
	FieldRoundUnits = "round_units"
	// FieldRoundDelay holds the string denoting the round_delay field in the database.
	FieldRoundDelay = "round_delay"
	// FieldServiceID holds the string denoting the service_id field in the database.
	FieldServiceID = "service_id"
	// FieldHostID holds the string denoting the host_id field in the database.
	FieldHostID = "host_id"
	// FieldTeamID holds the string denoting the team_id field in the database.
	FieldTeamID = "team_id"
	// EdgeChecks holds the string denoting the checks edge name in mutations.
	EdgeChecks = "checks"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the hostservice in the database.
	Table = "host_services"
	// ChecksTable is the table that holds the checks relation/edge.
	ChecksTable = "checks"
	// ChecksInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	ChecksInverseTable = "checks"
	// ChecksColumn is the table column denoting the checks relation/edge.
	ChecksColumn = "host_service_id"
	// PropertiesTable is the table that holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "host_service_id"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "host_services"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_id"
	// HostTable is the table that holds the host relation/edge.
	HostTable = "host_services"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "host_id"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "host_services"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_id"
)

// Columns holds all SQL columns for hostservice fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDisplayName,
	FieldPause,
	FieldHidden,
	FieldWeight,
	FieldPointBoost,
	FieldRoundUnits,
	FieldRoundDelay,
	FieldServiceID,
	FieldHostID,
	FieldTeamID,
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
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight int
	// DefaultPointBoost holds the default value on creation for the "point_boost" field.
	DefaultPointBoost int
	// DefaultRoundUnits holds the default value on creation for the "round_units" field.
	DefaultRoundUnits int
	// DefaultRoundDelay holds the default value on creation for the "round_delay" field.
	DefaultRoundDelay int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Order defines the ordering method for the HostService queries.
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

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByPointBoost orders the results by the point_boost field.
func ByPointBoost(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPointBoost, opts...).ToFunc()
}

// ByRoundUnits orders the results by the round_units field.
func ByRoundUnits(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldRoundUnits, opts...).ToFunc()
}

// ByRoundDelay orders the results by the round_delay field.
func ByRoundDelay(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldRoundDelay, opts...).ToFunc()
}

// ByServiceID orders the results by the service_id field.
func ByServiceID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldServiceID, opts...).ToFunc()
}

// ByHostID orders the results by the host_id field.
func ByHostID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldHostID, opts...).ToFunc()
}

// ByTeamID orders the results by the team_id field.
func ByTeamID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTeamID, opts...).ToFunc()
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

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByHostField orders the results by host field.
func ByHostField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostStep(), sql.OrderByField(field, opts...))
	}
}

// ByTeamField orders the results by team field.
func ByTeamField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), sql.OrderByField(field, opts...))
	}
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
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
	)
}
func newHostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HostTable, HostColumn),
	)
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
	)
}