// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/competition"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/team"
)

// HostGroup is the model entity for the HostGroup schema.
type HostGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Pause holds the value of the "pause" field.
	Pause bool `json:"pause,omitempty"`
	// Hidden holds the value of the "hidden" field.
	Hidden bool `json:"hidden,omitempty"`
	// CompetitionID holds the value of the "competition_id" field.
	CompetitionID int `json:"competition_id,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID int `json:"team_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostGroupQuery when eager-loading is set.
	Edges        HostGroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HostGroupEdges holds the relations/edges for other nodes in the graph.
type HostGroupEdges struct {
	// Competition holds the value of the competition edge.
	Competition *Competition `json:"competition,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Hosts holds the value of the hosts edge.
	Hosts []*Host `json:"hosts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CompetitionOrErr returns the Competition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostGroupEdges) CompetitionOrErr() (*Competition, error) {
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
func (e HostGroupEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[1] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// HostsOrErr returns the Hosts value or an error if the edge
// was not loaded in eager-loading.
func (e HostGroupEdges) HostsOrErr() ([]*Host, error) {
	if e.loadedTypes[2] {
		return e.Hosts, nil
	}
	return nil, &NotLoadedError{edge: "hosts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostgroup.FieldPause, hostgroup.FieldHidden:
			values[i] = new(sql.NullBool)
		case hostgroup.FieldID, hostgroup.FieldCompetitionID, hostgroup.FieldTeamID:
			values[i] = new(sql.NullInt64)
		case hostgroup.FieldName:
			values[i] = new(sql.NullString)
		case hostgroup.FieldCreateTime, hostgroup.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostGroup fields.
func (hg *HostGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hg.ID = int(value.Int64)
		case hostgroup.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				hg.CreateTime = value.Time
			}
		case hostgroup.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				hg.UpdateTime = value.Time
			}
		case hostgroup.FieldPause:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pause", values[i])
			} else if value.Valid {
				hg.Pause = value.Bool
			}
		case hostgroup.FieldHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hidden", values[i])
			} else if value.Valid {
				hg.Hidden = value.Bool
			}
		case hostgroup.FieldCompetitionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field competition_id", values[i])
			} else if value.Valid {
				hg.CompetitionID = int(value.Int64)
			}
		case hostgroup.FieldTeamID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				hg.TeamID = int(value.Int64)
			}
		case hostgroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				hg.Name = value.String
			}
		default:
			hg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostGroup.
// This includes values selected through modifiers, order, etc.
func (hg *HostGroup) Value(name string) (ent.Value, error) {
	return hg.selectValues.Get(name)
}

// QueryCompetition queries the "competition" edge of the HostGroup entity.
func (hg *HostGroup) QueryCompetition() *CompetitionQuery {
	return NewHostGroupClient(hg.config).QueryCompetition(hg)
}

// QueryTeam queries the "team" edge of the HostGroup entity.
func (hg *HostGroup) QueryTeam() *TeamQuery {
	return NewHostGroupClient(hg.config).QueryTeam(hg)
}

// QueryHosts queries the "hosts" edge of the HostGroup entity.
func (hg *HostGroup) QueryHosts() *HostQuery {
	return NewHostGroupClient(hg.config).QueryHosts(hg)
}

// Update returns a builder for updating this HostGroup.
// Note that you need to call HostGroup.Unwrap() before calling this method if this HostGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (hg *HostGroup) Update() *HostGroupUpdateOne {
	return NewHostGroupClient(hg.config).UpdateOne(hg)
}

// Unwrap unwraps the HostGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hg *HostGroup) Unwrap() *HostGroup {
	_tx, ok := hg.config.driver.(*txDriver)
	if !ok {
		panic("entities: HostGroup is not a transactional entity")
	}
	hg.config.driver = _tx.drv
	return hg
}

// String implements the fmt.Stringer.
func (hg *HostGroup) String() string {
	var builder strings.Builder
	builder.WriteString("HostGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hg.ID))
	builder.WriteString("create_time=")
	builder.WriteString(hg.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(hg.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("pause=")
	builder.WriteString(fmt.Sprintf("%v", hg.Pause))
	builder.WriteString(", ")
	builder.WriteString("hidden=")
	builder.WriteString(fmt.Sprintf("%v", hg.Hidden))
	builder.WriteString(", ")
	builder.WriteString("competition_id=")
	builder.WriteString(fmt.Sprintf("%v", hg.CompetitionID))
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(fmt.Sprintf("%v", hg.TeamID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(hg.Name)
	builder.WriteByte(')')
	return builder.String()
}

// HostGroups is a parsable slice of HostGroup.
type HostGroups []*HostGroup
