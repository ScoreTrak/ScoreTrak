// Code generated by ent, DO NOT EDIT.

package teamreport

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldContainsFold(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldUpdateTime, v))
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldPoints, v))
}

// TeamID applies equality check predicate on the "team_id" field. It's identical to TeamIDEQ.
func TeamID(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldTeamID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotNull(FieldUpdateTime))
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldPoints, v))
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNEQ(FieldPoints, v))
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIn(FieldPoints, vs...))
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotIn(FieldPoints, vs...))
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGT(FieldPoints, v))
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGTE(FieldPoints, v))
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLT(FieldPoints, v))
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v int) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLTE(FieldPoints, v))
}

// TeamIDEQ applies the EQ predicate on the "team_id" field.
func TeamIDEQ(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEQ(FieldTeamID, v))
}

// TeamIDNEQ applies the NEQ predicate on the "team_id" field.
func TeamIDNEQ(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNEQ(FieldTeamID, v))
}

// TeamIDIn applies the In predicate on the "team_id" field.
func TeamIDIn(vs ...string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldIn(FieldTeamID, vs...))
}

// TeamIDNotIn applies the NotIn predicate on the "team_id" field.
func TeamIDNotIn(vs ...string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldNotIn(FieldTeamID, vs...))
}

// TeamIDGT applies the GT predicate on the "team_id" field.
func TeamIDGT(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGT(FieldTeamID, v))
}

// TeamIDGTE applies the GTE predicate on the "team_id" field.
func TeamIDGTE(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldGTE(FieldTeamID, v))
}

// TeamIDLT applies the LT predicate on the "team_id" field.
func TeamIDLT(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLT(FieldTeamID, v))
}

// TeamIDLTE applies the LTE predicate on the "team_id" field.
func TeamIDLTE(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldLTE(FieldTeamID, v))
}

// TeamIDContains applies the Contains predicate on the "team_id" field.
func TeamIDContains(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldContains(FieldTeamID, v))
}

// TeamIDHasPrefix applies the HasPrefix predicate on the "team_id" field.
func TeamIDHasPrefix(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldHasPrefix(FieldTeamID, v))
}

// TeamIDHasSuffix applies the HasSuffix predicate on the "team_id" field.
func TeamIDHasSuffix(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldHasSuffix(FieldTeamID, v))
}

// TeamIDEqualFold applies the EqualFold predicate on the "team_id" field.
func TeamIDEqualFold(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldEqualFold(FieldTeamID, v))
}

// TeamIDContainsFold applies the ContainsFold predicate on the "team_id" field.
func TeamIDContainsFold(v string) predicate.TeamReport {
	return predicate.TeamReport(sql.FieldContainsFold(FieldTeamID, v))
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TeamTable, TeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHostservicereports applies the HasEdge predicate on the "hostservicereports" edge.
func HasHostservicereports() predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HostservicereportsTable, HostservicereportsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostservicereportsWith applies the HasEdge predicate on the "hostservicereports" edge with a given conditions (other predicates).
func HasHostservicereportsWith(preds ...predicate.HostServiceReport) predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		step := newHostservicereportsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TeamReport) predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TeamReport) predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TeamReport) predicate.TeamReport {
	return predicate.TeamReport(func(s *sql.Selector) {
		p(s.Not())
	})
}
