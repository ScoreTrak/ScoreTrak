// Code generated by ent, DO NOT EDIT.

package competition

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldUpdateTime, v))
}

// Hidden applies equality check predicate on the "hidden" field. It's identical to HiddenEQ.
func Hidden(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldHidden, v))
}

// Pause applies equality check predicate on the "pause" field. It's identical to PauseEQ.
func Pause(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPause, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldDisplayName, v))
}

// RoundDuration applies equality check predicate on the "round_duration" field. It's identical to RoundDurationEQ.
func RoundDuration(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldRoundDuration, v))
}

// ToBeStartedAt applies equality check predicate on the "to_be_started_at" field. It's identical to ToBeStartedAtEQ.
func ToBeStartedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldToBeStartedAt, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStartedAt, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldFinishedAt, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldUpdateTime, v))
}

// HiddenEQ applies the EQ predicate on the "hidden" field.
func HiddenEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldHidden, v))
}

// HiddenNEQ applies the NEQ predicate on the "hidden" field.
func HiddenNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldHidden, v))
}

// PauseEQ applies the EQ predicate on the "pause" field.
func PauseEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPause, v))
}

// PauseNEQ applies the NEQ predicate on the "pause" field.
func PauseNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldPause, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldName, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Competition {
	return predicate.Competition(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldDisplayName, v))
}

// RoundDurationEQ applies the EQ predicate on the "round_duration" field.
func RoundDurationEQ(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldRoundDuration, v))
}

// RoundDurationNEQ applies the NEQ predicate on the "round_duration" field.
func RoundDurationNEQ(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldRoundDuration, v))
}

// RoundDurationIn applies the In predicate on the "round_duration" field.
func RoundDurationIn(vs ...float64) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldRoundDuration, vs...))
}

// RoundDurationNotIn applies the NotIn predicate on the "round_duration" field.
func RoundDurationNotIn(vs ...float64) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldRoundDuration, vs...))
}

// RoundDurationGT applies the GT predicate on the "round_duration" field.
func RoundDurationGT(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldRoundDuration, v))
}

// RoundDurationGTE applies the GTE predicate on the "round_duration" field.
func RoundDurationGTE(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldRoundDuration, v))
}

// RoundDurationLT applies the LT predicate on the "round_duration" field.
func RoundDurationLT(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldRoundDuration, v))
}

// RoundDurationLTE applies the LTE predicate on the "round_duration" field.
func RoundDurationLTE(v float64) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldRoundDuration, v))
}

// ToBeStartedAtEQ applies the EQ predicate on the "to_be_started_at" field.
func ToBeStartedAtEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldToBeStartedAt, v))
}

// ToBeStartedAtNEQ applies the NEQ predicate on the "to_be_started_at" field.
func ToBeStartedAtNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldToBeStartedAt, v))
}

// ToBeStartedAtIn applies the In predicate on the "to_be_started_at" field.
func ToBeStartedAtIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldToBeStartedAt, vs...))
}

// ToBeStartedAtNotIn applies the NotIn predicate on the "to_be_started_at" field.
func ToBeStartedAtNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldToBeStartedAt, vs...))
}

// ToBeStartedAtGT applies the GT predicate on the "to_be_started_at" field.
func ToBeStartedAtGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldToBeStartedAt, v))
}

// ToBeStartedAtGTE applies the GTE predicate on the "to_be_started_at" field.
func ToBeStartedAtGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldToBeStartedAt, v))
}

// ToBeStartedAtLT applies the LT predicate on the "to_be_started_at" field.
func ToBeStartedAtLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldToBeStartedAt, v))
}

// ToBeStartedAtLTE applies the LTE predicate on the "to_be_started_at" field.
func ToBeStartedAtLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldToBeStartedAt, v))
}

// ToBeStartedAtIsNil applies the IsNil predicate on the "to_be_started_at" field.
func ToBeStartedAtIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldToBeStartedAt))
}

// ToBeStartedAtNotNil applies the NotNil predicate on the "to_be_started_at" field.
func ToBeStartedAtNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldToBeStartedAt))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldStartedAt))
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldFinishedAt, v))
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldFinishedAt, v))
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldFinishedAt, vs...))
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldFinishedAt, vs...))
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldFinishedAt, v))
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldFinishedAt, v))
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldFinishedAt, v))
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldFinishedAt, v))
}

// FinishedAtIsNil applies the IsNil predicate on the "finished_at" field.
func FinishedAtIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldFinishedAt))
}

// FinishedAtNotNil applies the NotNil predicate on the "finished_at" field.
func FinishedAtNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldFinishedAt))
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeamsTable, TeamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Competition) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Competition) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
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
func Not(p predicate.Competition) predicate.Competition {
	return predicate.Competition(func(s *sql.Selector) {
		p(s.Not())
	})
}
