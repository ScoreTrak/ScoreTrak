// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/host"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// HostService is the model entity for the HostService schema.
type HostService struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Pause holds the value of the "pause" field.
	Pause bool `json:"pause,omitempty"`
	// Hidden holds the value of the "hidden" field.
	Hidden bool `json:"hidden,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// PointBoost holds the value of the "point_boost" field.
	PointBoost int `json:"point_boost,omitempty"`
	// RoundUnits holds the value of the "round_units" field.
	RoundUnits int `json:"round_units,omitempty"`
	// RoundDelay holds the value of the "round_delay" field.
	RoundDelay int `json:"round_delay,omitempty"`
	// ServiceID holds the value of the "service_id" field.
	ServiceID string `json:"service_id,omitempty"`
	// HostID holds the value of the "host_id" field.
	HostID string `json:"host_id,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID string `json:"team_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostServiceQuery when eager-loading is set.
	Edges        HostServiceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HostServiceEdges holds the relations/edges for other nodes in the graph.
type HostServiceEdges struct {
	// Checks holds the value of the checks edge.
	Checks []*Check `json:"checks,omitempty"`
	// Properties holds the value of the properties edge.
	Properties []*Property `json:"properties,omitempty"`
	// Service holds the value of the service edge.
	Service *Service `json:"service,omitempty"`
	// Host holds the value of the host edge.
	Host *Host `json:"host,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ChecksOrErr returns the Checks value or an error if the edge
// was not loaded in eager-loading.
func (e HostServiceEdges) ChecksOrErr() ([]*Check, error) {
	if e.loadedTypes[0] {
		return e.Checks, nil
	}
	return nil, &NotLoadedError{edge: "checks"}
}

// PropertiesOrErr returns the Properties value or an error if the edge
// was not loaded in eager-loading.
func (e HostServiceEdges) PropertiesOrErr() ([]*Property, error) {
	if e.loadedTypes[1] {
		return e.Properties, nil
	}
	return nil, &NotLoadedError{edge: "properties"}
}

// ServiceOrErr returns the Service value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceEdges) ServiceOrErr() (*Service, error) {
	if e.loadedTypes[2] {
		if e.Service == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: service.Label}
		}
		return e.Service, nil
	}
	return nil, &NotLoadedError{edge: "service"}
}

// HostOrErr returns the Host value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceEdges) HostOrErr() (*Host, error) {
	if e.loadedTypes[3] {
		if e.Host == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.Host, nil
	}
	return nil, &NotLoadedError{edge: "host"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[4] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostService) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostservice.FieldPause, hostservice.FieldHidden:
			values[i] = new(sql.NullBool)
		case hostservice.FieldWeight, hostservice.FieldPointBoost, hostservice.FieldRoundUnits, hostservice.FieldRoundDelay:
			values[i] = new(sql.NullInt64)
		case hostservice.FieldID, hostservice.FieldName, hostservice.FieldDisplayName, hostservice.FieldServiceID, hostservice.FieldHostID, hostservice.FieldTeamID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostService fields.
func (hs *HostService) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostservice.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				hs.ID = value.String
			}
		case hostservice.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				hs.Name = value.String
			}
		case hostservice.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				hs.DisplayName = value.String
			}
		case hostservice.FieldPause:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pause", values[i])
			} else if value.Valid {
				hs.Pause = value.Bool
			}
		case hostservice.FieldHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hidden", values[i])
			} else if value.Valid {
				hs.Hidden = value.Bool
			}
		case hostservice.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				hs.Weight = int(value.Int64)
			}
		case hostservice.FieldPointBoost:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field point_boost", values[i])
			} else if value.Valid {
				hs.PointBoost = int(value.Int64)
			}
		case hostservice.FieldRoundUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field round_units", values[i])
			} else if value.Valid {
				hs.RoundUnits = int(value.Int64)
			}
		case hostservice.FieldRoundDelay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field round_delay", values[i])
			} else if value.Valid {
				hs.RoundDelay = int(value.Int64)
			}
		case hostservice.FieldServiceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_id", values[i])
			} else if value.Valid {
				hs.ServiceID = value.String
			}
		case hostservice.FieldHostID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_id", values[i])
			} else if value.Valid {
				hs.HostID = value.String
			}
		case hostservice.FieldTeamID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				hs.TeamID = value.String
			}
		default:
			hs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostService.
// This includes values selected through modifiers, order, etc.
func (hs *HostService) Value(name string) (ent.Value, error) {
	return hs.selectValues.Get(name)
}

// QueryChecks queries the "checks" edge of the HostService entity.
func (hs *HostService) QueryChecks() *CheckQuery {
	return NewHostServiceClient(hs.config).QueryChecks(hs)
}

// QueryProperties queries the "properties" edge of the HostService entity.
func (hs *HostService) QueryProperties() *PropertyQuery {
	return NewHostServiceClient(hs.config).QueryProperties(hs)
}

// QueryService queries the "service" edge of the HostService entity.
func (hs *HostService) QueryService() *ServiceQuery {
	return NewHostServiceClient(hs.config).QueryService(hs)
}

// QueryHost queries the "host" edge of the HostService entity.
func (hs *HostService) QueryHost() *HostQuery {
	return NewHostServiceClient(hs.config).QueryHost(hs)
}

// QueryTeam queries the "team" edge of the HostService entity.
func (hs *HostService) QueryTeam() *TeamQuery {
	return NewHostServiceClient(hs.config).QueryTeam(hs)
}

// Update returns a builder for updating this HostService.
// Note that you need to call HostService.Unwrap() before calling this method if this HostService
// was returned from a transaction, and the transaction was committed or rolled back.
func (hs *HostService) Update() *HostServiceUpdateOne {
	return NewHostServiceClient(hs.config).UpdateOne(hs)
}

// Unwrap unwraps the HostService entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hs *HostService) Unwrap() *HostService {
	_tx, ok := hs.config.driver.(*txDriver)
	if !ok {
		panic("entities: HostService is not a transactional entity")
	}
	hs.config.driver = _tx.drv
	return hs
}

// String implements the fmt.Stringer.
func (hs *HostService) String() string {
	var builder strings.Builder
	builder.WriteString("HostService(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hs.ID))
	builder.WriteString("name=")
	builder.WriteString(hs.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(hs.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("pause=")
	builder.WriteString(fmt.Sprintf("%v", hs.Pause))
	builder.WriteString(", ")
	builder.WriteString("hidden=")
	builder.WriteString(fmt.Sprintf("%v", hs.Hidden))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", hs.Weight))
	builder.WriteString(", ")
	builder.WriteString("point_boost=")
	builder.WriteString(fmt.Sprintf("%v", hs.PointBoost))
	builder.WriteString(", ")
	builder.WriteString("round_units=")
	builder.WriteString(fmt.Sprintf("%v", hs.RoundUnits))
	builder.WriteString(", ")
	builder.WriteString("round_delay=")
	builder.WriteString(fmt.Sprintf("%v", hs.RoundDelay))
	builder.WriteString(", ")
	builder.WriteString("service_id=")
	builder.WriteString(hs.ServiceID)
	builder.WriteString(", ")
	builder.WriteString("host_id=")
	builder.WriteString(hs.HostID)
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(hs.TeamID)
	builder.WriteByte(')')
	return builder.String()
}

// HostServices is a parsable slice of HostService.
type HostServices []*HostService
