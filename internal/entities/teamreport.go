// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/scoretrak/scoretrak/internal/entities/team"
	"github.com/scoretrak/scoretrak/internal/entities/teamreport"
)

// TeamReport is the model entity for the TeamReport schema.
type TeamReport struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID string `json:"team_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamReportQuery when eager-loading is set.
	Edges        TeamReportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TeamReportEdges holds the relations/edges for other nodes in the graph.
type TeamReportEdges struct {
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Hostservicereports holds the value of the hostservicereports edge.
	Hostservicereports []*HostServiceReport `json:"hostservicereports,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamReportEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[0] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// HostservicereportsOrErr returns the Hostservicereports value or an error if the edge
// was not loaded in eager-loading.
func (e TeamReportEdges) HostservicereportsOrErr() ([]*HostServiceReport, error) {
	if e.loadedTypes[1] {
		return e.Hostservicereports, nil
	}
	return nil, &NotLoadedError{edge: "hostservicereports"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TeamReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case teamreport.FieldPoints:
			values[i] = new(sql.NullInt64)
		case teamreport.FieldID, teamreport.FieldTeamID:
			values[i] = new(sql.NullString)
		case teamreport.FieldCreateTime, teamreport.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TeamReport fields.
func (tr *TeamReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teamreport.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				tr.ID = value.String
			}
		case teamreport.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				tr.CreateTime = value.Time
			}
		case teamreport.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				tr.UpdateTime = value.Time
			}
		case teamreport.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				tr.Points = int(value.Int64)
			}
		case teamreport.FieldTeamID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				tr.TeamID = value.String
			}
		default:
			tr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TeamReport.
// This includes values selected through modifiers, order, etc.
func (tr *TeamReport) Value(name string) (ent.Value, error) {
	return tr.selectValues.Get(name)
}

// QueryTeam queries the "team" edge of the TeamReport entity.
func (tr *TeamReport) QueryTeam() *TeamQuery {
	return NewTeamReportClient(tr.config).QueryTeam(tr)
}

// QueryHostservicereports queries the "hostservicereports" edge of the TeamReport entity.
func (tr *TeamReport) QueryHostservicereports() *HostServiceReportQuery {
	return NewTeamReportClient(tr.config).QueryHostservicereports(tr)
}

// Update returns a builder for updating this TeamReport.
// Note that you need to call TeamReport.Unwrap() before calling this method if this TeamReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TeamReport) Update() *TeamReportUpdateOne {
	return NewTeamReportClient(tr.config).UpdateOne(tr)
}

// Unwrap unwraps the TeamReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tr *TeamReport) Unwrap() *TeamReport {
	_tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("entities: TeamReport is not a transactional entity")
	}
	tr.config.driver = _tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TeamReport) String() string {
	var builder strings.Builder
	builder.WriteString("TeamReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(tr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(tr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", tr.Points))
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(tr.TeamID)
	builder.WriteByte(')')
	return builder.String()
}

// TeamReports is a parsable slice of TeamReport.
type TeamReports []*TeamReport
