// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/round"
)

// Check is the model entity for the Check schema.
type Check struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Log holds the value of the "log" field.
	Log string `json:"log,omitempty"`
	// Error holds the value of the "error" field.
	Error string `json:"error,omitempty"`
	// Passed holds the value of the "passed" field.
	Passed bool `json:"passed,omitempty"`
	// RoundID holds the value of the "round_id" field.
	RoundID string `json:"round_id,omitempty"`
	// HostServiceID holds the value of the "host_service_id" field.
	HostServiceID string `json:"host_service_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckQuery when eager-loading is set.
	Edges        CheckEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CheckEdges holds the relations/edges for other nodes in the graph.
type CheckEdges struct {
	// Round holds the value of the round edge.
	Round *Round `json:"round,omitempty"`
	// Hostservice holds the value of the hostservice edge.
	Hostservice *HostService `json:"hostservice,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoundOrErr returns the Round value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckEdges) RoundOrErr() (*Round, error) {
	if e.loadedTypes[0] {
		if e.Round == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: round.Label}
		}
		return e.Round, nil
	}
	return nil, &NotLoadedError{edge: "round"}
}

// HostserviceOrErr returns the Hostservice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckEdges) HostserviceOrErr() (*HostService, error) {
	if e.loadedTypes[1] {
		if e.Hostservice == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hostservice.Label}
		}
		return e.Hostservice, nil
	}
	return nil, &NotLoadedError{edge: "hostservice"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Check) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case check.FieldPassed:
			values[i] = new(sql.NullBool)
		case check.FieldID, check.FieldLog, check.FieldError, check.FieldRoundID, check.FieldHostServiceID:
			values[i] = new(sql.NullString)
		case check.FieldCreateTime, check.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Check fields.
func (c *Check) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case check.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case check.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case check.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case check.FieldLog:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log", values[i])
			} else if value.Valid {
				c.Log = value.String
			}
		case check.FieldError:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error", values[i])
			} else if value.Valid {
				c.Error = value.String
			}
		case check.FieldPassed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field passed", values[i])
			} else if value.Valid {
				c.Passed = value.Bool
			}
		case check.FieldRoundID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field round_id", values[i])
			} else if value.Valid {
				c.RoundID = value.String
			}
		case check.FieldHostServiceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_service_id", values[i])
			} else if value.Valid {
				c.HostServiceID = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Check.
// This includes values selected through modifiers, order, etc.
func (c *Check) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryRound queries the "round" edge of the Check entity.
func (c *Check) QueryRound() *RoundQuery {
	return NewCheckClient(c.config).QueryRound(c)
}

// QueryHostservice queries the "hostservice" edge of the Check entity.
func (c *Check) QueryHostservice() *HostServiceQuery {
	return NewCheckClient(c.config).QueryHostservice(c)
}

// Update returns a builder for updating this Check.
// Note that you need to call Check.Unwrap() before calling this method if this Check
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Check) Update() *CheckUpdateOne {
	return NewCheckClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Check entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Check) Unwrap() *Check {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entities: Check is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Check) String() string {
	var builder strings.Builder
	builder.WriteString("Check(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("log=")
	builder.WriteString(c.Log)
	builder.WriteString(", ")
	builder.WriteString("error=")
	builder.WriteString(c.Error)
	builder.WriteString(", ")
	builder.WriteString("passed=")
	builder.WriteString(fmt.Sprintf("%v", c.Passed))
	builder.WriteString(", ")
	builder.WriteString("round_id=")
	builder.WriteString(c.RoundID)
	builder.WriteString(", ")
	builder.WriteString("host_service_id=")
	builder.WriteString(c.HostServiceID)
	builder.WriteByte(')')
	return builder.String()
}

// Checks is a parsable slice of Check.
type Checks []*Check
