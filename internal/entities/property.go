// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/property"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/service"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// Property is the model entity for the Property schema.
type Property struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CompetitionID holds the value of the "competition_id" field.
	CompetitionID string `json:"competition_id,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID string `json:"team_id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Status holds the value of the "status" field.
	Status property.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PropertyQuery when eager-loading is set.
	Edges              PropertyEdges `json:"edges"`
	service_properties *string
	selectValues       sql.SelectValues
}

// PropertyEdges holds the relations/edges for other nodes in the graph.
type PropertyEdges struct {
	// Competition holds the value of the competition edge.
	Competition *Competition `json:"competition,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Services holds the value of the services edge.
	Services *Service `json:"services,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CompetitionOrErr returns the Competition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PropertyEdges) CompetitionOrErr() (*Competition, error) {
	if e.loadedTypes[0] {
		if e.Competition == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: competition.Label}
		}
		return e.Competition, nil
	}
	return nil, &NotLoadedError{edge: "competition"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PropertyEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[1] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PropertyEdges) ServicesOrErr() (*Service, error) {
	if e.loadedTypes[2] {
		if e.Services == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: service.Label}
		}
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Property) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case property.FieldID, property.FieldCompetitionID, property.FieldTeamID, property.FieldKey, property.FieldValue, property.FieldStatus:
			values[i] = new(sql.NullString)
		case property.ForeignKeys[0]: // service_properties
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Property fields.
func (pr *Property) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case property.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case property.FieldCompetitionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field competition_id", values[i])
			} else if value.Valid {
				pr.CompetitionID = value.String
			}
		case property.FieldTeamID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				pr.TeamID = value.String
			}
		case property.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				pr.Key = value.String
			}
		case property.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				pr.Value = value.String
			}
		case property.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = property.Status(value.String)
			}
		case property.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_properties", values[i])
			} else if value.Valid {
				pr.service_properties = new(string)
				*pr.service_properties = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Property.
// This includes values selected through modifiers, order, etc.
func (pr *Property) GetValue(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryCompetition queries the "competition" edge of the Property entity.
func (pr *Property) QueryCompetition() *CompetitionQuery {
	return NewPropertyClient(pr.config).QueryCompetition(pr)
}

// QueryTeam queries the "team" edge of the Property entity.
func (pr *Property) QueryTeam() *TeamQuery {
	return NewPropertyClient(pr.config).QueryTeam(pr)
}

// QueryServices queries the "services" edge of the Property entity.
func (pr *Property) QueryServices() *ServiceQuery {
	return NewPropertyClient(pr.config).QueryServices(pr)
}

// Update returns a builder for updating this Property.
// Note that you need to call Property.Unwrap() before calling this method if this Property
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Property) Update() *PropertyUpdateOne {
	return NewPropertyClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Property entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Property) Unwrap() *Property {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("entities: Property is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Property) String() string {
	var builder strings.Builder
	builder.WriteString("Property(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("competition_id=")
	builder.WriteString(pr.CompetitionID)
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(pr.TeamID)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(pr.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(pr.Value)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Properties is a parsable slice of Property.
type Properties []*Property
