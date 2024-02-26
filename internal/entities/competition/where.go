// Code generated by ent, DO NOT EDIT.

package competition

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scoretrak/scoretrak/internal/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Competition {
	return predicate.Competition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Competition {
	return predicate.Competition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Competition {
	return predicate.Competition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Competition {
	return predicate.Competition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Competition {
	return predicate.Competition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Competition {
	return predicate.Competition(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Competition {
	return predicate.Competition(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Competition {
	return predicate.Competition(sql.FieldContainsFold(FieldID, id))
}

// Pause applies equality check predicate on the "pause" field. It's identical to PauseEQ.
func Pause(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPause, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldDisplayName, v))
}

// ViewableToPublic applies equality check predicate on the "viewable_to_public" field. It's identical to ViewableToPublicEQ.
func ViewableToPublic(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldViewableToPublic, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldStartedAt, v))
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldFinishedAt, v))
}

// PauseEQ applies the EQ predicate on the "pause" field.
func PauseEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldPause, v))
}

// PauseNEQ applies the NEQ predicate on the "pause" field.
func PauseNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldPause, v))
}

// PauseIsNil applies the IsNil predicate on the "pause" field.
func PauseIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldPause))
}

// PauseNotNil applies the NotNil predicate on the "pause" field.
func PauseNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldPause))
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

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldCreateTime))
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

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldUpdateTime))
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

// ViewableToPublicEQ applies the EQ predicate on the "viewable_to_public" field.
func ViewableToPublicEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldEQ(FieldViewableToPublic, v))
}

// ViewableToPublicNEQ applies the NEQ predicate on the "viewable_to_public" field.
func ViewableToPublicNEQ(v bool) predicate.Competition {
	return predicate.Competition(sql.FieldNEQ(FieldViewableToPublic, v))
}

// ViewableToPublicIsNil applies the IsNil predicate on the "viewable_to_public" field.
func ViewableToPublicIsNil() predicate.Competition {
	return predicate.Competition(sql.FieldIsNull(FieldViewableToPublic))
}

// ViewableToPublicNotNil applies the NotNil predicate on the "viewable_to_public" field.
func ViewableToPublicNotNil() predicate.Competition {
	return predicate.Competition(sql.FieldNotNull(FieldViewableToPublic))
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
