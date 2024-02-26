// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/scoretrak/scoretrak/internal/entities/hostservice"
	"github.com/scoretrak/scoretrak/internal/entities/hostservicereport"
	"github.com/scoretrak/scoretrak/internal/entities/service"
	"github.com/scoretrak/scoretrak/internal/entities/team"
	"github.com/scoretrak/scoretrak/internal/entities/teamreport"
)

// HostServiceReport is the model entity for the HostServiceReport schema.
type HostServiceReport struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// Passing holds the value of the "passing" field.
	Passing bool `json:"passing,omitempty"`
	// LatestCheckTime holds the value of the "latest_check_time" field.
	LatestCheckTime time.Time `json:"latest_check_time,omitempty"`
	// HostServiceID holds the value of the "host_service_id" field.
	HostServiceID string `json:"host_service_id,omitempty"`
	// ServiceID holds the value of the "service_id" field.
	ServiceID string `json:"service_id,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID string `json:"team_id,omitempty"`
	// TeamReportID holds the value of the "team_report_id" field.
	TeamReportID string `json:"team_report_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostServiceReportQuery when eager-loading is set.
	Edges        HostServiceReportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HostServiceReportEdges holds the relations/edges for other nodes in the graph.
type HostServiceReportEdges struct {
	// Hostservice holds the value of the hostservice edge.
	Hostservice *HostService `json:"hostservice,omitempty"`
	// Service holds the value of the service edge.
	Service *Service `json:"service,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Teamreport holds the value of the teamreport edge.
	Teamreport *TeamReport `json:"teamreport,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// HostserviceOrErr returns the Hostservice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceReportEdges) HostserviceOrErr() (*HostService, error) {
	if e.loadedTypes[0] {
		if e.Hostservice == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hostservice.Label}
		}
		return e.Hostservice, nil
	}
	return nil, &NotLoadedError{edge: "hostservice"}
}

// ServiceOrErr returns the Service value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceReportEdges) ServiceOrErr() (*Service, error) {
	if e.loadedTypes[1] {
		if e.Service == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: service.Label}
		}
		return e.Service, nil
	}
	return nil, &NotLoadedError{edge: "service"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceReportEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[2] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// TeamreportOrErr returns the Teamreport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostServiceReportEdges) TeamreportOrErr() (*TeamReport, error) {
	if e.loadedTypes[3] {
		if e.Teamreport == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teamreport.Label}
		}
		return e.Teamreport, nil
	}
	return nil, &NotLoadedError{edge: "teamreport"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostServiceReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostservicereport.FieldPassing:
			values[i] = new(sql.NullBool)
		case hostservicereport.FieldPoints:
			values[i] = new(sql.NullInt64)
		case hostservicereport.FieldID, hostservicereport.FieldHostServiceID, hostservicereport.FieldServiceID, hostservicereport.FieldTeamID, hostservicereport.FieldTeamReportID:
			values[i] = new(sql.NullString)
		case hostservicereport.FieldCreateTime, hostservicereport.FieldUpdateTime, hostservicereport.FieldLatestCheckTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostServiceReport fields.
func (hsr *HostServiceReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostservicereport.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				hsr.ID = value.String
			}
		case hostservicereport.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				hsr.CreateTime = value.Time
			}
		case hostservicereport.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				hsr.UpdateTime = value.Time
			}
		case hostservicereport.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				hsr.Points = int(value.Int64)
			}
		case hostservicereport.FieldPassing:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field passing", values[i])
			} else if value.Valid {
				hsr.Passing = value.Bool
			}
		case hostservicereport.FieldLatestCheckTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field latest_check_time", values[i])
			} else if value.Valid {
				hsr.LatestCheckTime = value.Time
			}
		case hostservicereport.FieldHostServiceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_service_id", values[i])
			} else if value.Valid {
				hsr.HostServiceID = value.String
			}
		case hostservicereport.FieldServiceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_id", values[i])
			} else if value.Valid {
				hsr.ServiceID = value.String
			}
		case hostservicereport.FieldTeamID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				hsr.TeamID = value.String
			}
		case hostservicereport.FieldTeamReportID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_report_id", values[i])
			} else if value.Valid {
				hsr.TeamReportID = value.String
			}
		default:
			hsr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostServiceReport.
// This includes values selected through modifiers, order, etc.
func (hsr *HostServiceReport) Value(name string) (ent.Value, error) {
	return hsr.selectValues.Get(name)
}

// QueryHostservice queries the "hostservice" edge of the HostServiceReport entity.
func (hsr *HostServiceReport) QueryHostservice() *HostServiceQuery {
	return NewHostServiceReportClient(hsr.config).QueryHostservice(hsr)
}

// QueryService queries the "service" edge of the HostServiceReport entity.
func (hsr *HostServiceReport) QueryService() *ServiceQuery {
	return NewHostServiceReportClient(hsr.config).QueryService(hsr)
}

// QueryTeam queries the "team" edge of the HostServiceReport entity.
func (hsr *HostServiceReport) QueryTeam() *TeamQuery {
	return NewHostServiceReportClient(hsr.config).QueryTeam(hsr)
}

// QueryTeamreport queries the "teamreport" edge of the HostServiceReport entity.
func (hsr *HostServiceReport) QueryTeamreport() *TeamReportQuery {
	return NewHostServiceReportClient(hsr.config).QueryTeamreport(hsr)
}

// Update returns a builder for updating this HostServiceReport.
// Note that you need to call HostServiceReport.Unwrap() before calling this method if this HostServiceReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (hsr *HostServiceReport) Update() *HostServiceReportUpdateOne {
	return NewHostServiceReportClient(hsr.config).UpdateOne(hsr)
}

// Unwrap unwraps the HostServiceReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hsr *HostServiceReport) Unwrap() *HostServiceReport {
	_tx, ok := hsr.config.driver.(*txDriver)
	if !ok {
		panic("entities: HostServiceReport is not a transactional entity")
	}
	hsr.config.driver = _tx.drv
	return hsr
}

// String implements the fmt.Stringer.
func (hsr *HostServiceReport) String() string {
	var builder strings.Builder
	builder.WriteString("HostServiceReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hsr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(hsr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(hsr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", hsr.Points))
	builder.WriteString(", ")
	builder.WriteString("passing=")
	builder.WriteString(fmt.Sprintf("%v", hsr.Passing))
	builder.WriteString(", ")
	builder.WriteString("latest_check_time=")
	builder.WriteString(hsr.LatestCheckTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("host_service_id=")
	builder.WriteString(hsr.HostServiceID)
	builder.WriteString(", ")
	builder.WriteString("service_id=")
	builder.WriteString(hsr.ServiceID)
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(hsr.TeamID)
	builder.WriteString(", ")
	builder.WriteString("team_report_id=")
	builder.WriteString(hsr.TeamReportID)
	builder.WriteByte(')')
	return builder.String()
}

// HostServiceReports is a parsable slice of HostServiceReport.
type HostServiceReports []*HostServiceReport
