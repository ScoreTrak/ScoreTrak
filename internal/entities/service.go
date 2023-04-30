// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
)

// Service is the model entity for the Service schema.
type Service struct {
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
	// CompetitionID holds the value of the "competition_id" field.
	CompetitionID string `json:"competition_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceQuery when eager-loading is set.
	Edges        ServiceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServiceEdges holds the relations/edges for other nodes in the graph.
type ServiceEdges struct {
	// Competition holds the value of the competition edge.
	Competition *Competition `json:"competition,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CompetitionOrErr returns the Competition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceEdges) CompetitionOrErr() (*Competition, error) {
	if e.loadedTypes[0] {
		if e.Competition == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: competition.Label}
		}
		return e.Competition, nil
	}
	return nil, &NotLoadedError{edge: "competition"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldPause, service.FieldHidden:
			values[i] = new(sql.NullBool)
		case service.FieldID, service.FieldName, service.FieldDisplayName, service.FieldCompetitionID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case service.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case service.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				s.DisplayName = value.String
			}
		case service.FieldPause:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pause", values[i])
			} else if value.Valid {
				s.Pause = value.Bool
			}
		case service.FieldHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hidden", values[i])
			} else if value.Valid {
				s.Hidden = value.Bool
			}
		case service.FieldCompetitionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field competition_id", values[i])
			} else if value.Valid {
				s.CompetitionID = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Service.
// This includes values selected through modifiers, order, etc.
func (s *Service) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryCompetition queries the "competition" edge of the Service entity.
func (s *Service) QueryCompetition() *CompetitionQuery {
	return NewServiceClient(s.config).QueryCompetition(s)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("entities: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(s.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("pause=")
	builder.WriteString(fmt.Sprintf("%v", s.Pause))
	builder.WriteString(", ")
	builder.WriteString("hidden=")
	builder.WriteString(fmt.Sprintf("%v", s.Hidden))
	builder.WriteString(", ")
	builder.WriteString("competition_id=")
	builder.WriteString(s.CompetitionID)
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service
