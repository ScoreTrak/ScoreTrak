// Code generated by ent, DO NOT EDIT.

package hostservice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldDisplayName, v))
}

// Pause applies equality check predicate on the "pause" field. It's identical to PauseEQ.
func Pause(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldPause, v))
}

// Hidden applies equality check predicate on the "hidden" field. It's identical to HiddenEQ.
func Hidden(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldHidden, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldWeight, v))
}

// PointBoost applies equality check predicate on the "point_boost" field. It's identical to PointBoostEQ.
func PointBoost(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldPointBoost, v))
}

// RoundUnits applies equality check predicate on the "round_units" field. It's identical to RoundUnitsEQ.
func RoundUnits(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldRoundUnits, v))
}

// RoundDelay applies equality check predicate on the "round_delay" field. It's identical to RoundDelayEQ.
func RoundDelay(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldRoundDelay, v))
}

// ServiceID applies equality check predicate on the "service_id" field. It's identical to ServiceIDEQ.
func ServiceID(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldServiceID, v))
}

// HostID applies equality check predicate on the "host_id" field. It's identical to HostIDEQ.
func HostID(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldHostID, v))
}

// TeamID applies equality check predicate on the "team_id" field. It's identical to TeamIDEQ.
func TeamID(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldTeamID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldName, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldDisplayName, v))
}

// PauseEQ applies the EQ predicate on the "pause" field.
func PauseEQ(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldPause, v))
}

// PauseNEQ applies the NEQ predicate on the "pause" field.
func PauseNEQ(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldPause, v))
}

// PauseIsNil applies the IsNil predicate on the "pause" field.
func PauseIsNil() predicate.HostService {
	return predicate.HostService(sql.FieldIsNull(FieldPause))
}

// PauseNotNil applies the NotNil predicate on the "pause" field.
func PauseNotNil() predicate.HostService {
	return predicate.HostService(sql.FieldNotNull(FieldPause))
}

// HiddenEQ applies the EQ predicate on the "hidden" field.
func HiddenEQ(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldHidden, v))
}

// HiddenNEQ applies the NEQ predicate on the "hidden" field.
func HiddenNEQ(v bool) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldHidden, v))
}

// HiddenIsNil applies the IsNil predicate on the "hidden" field.
func HiddenIsNil() predicate.HostService {
	return predicate.HostService(sql.FieldIsNull(FieldHidden))
}

// HiddenNotNil applies the NotNil predicate on the "hidden" field.
func HiddenNotNil() predicate.HostService {
	return predicate.HostService(sql.FieldNotNull(FieldHidden))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldWeight, v))
}

// PointBoostEQ applies the EQ predicate on the "point_boost" field.
func PointBoostEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldPointBoost, v))
}

// PointBoostNEQ applies the NEQ predicate on the "point_boost" field.
func PointBoostNEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldPointBoost, v))
}

// PointBoostIn applies the In predicate on the "point_boost" field.
func PointBoostIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldPointBoost, vs...))
}

// PointBoostNotIn applies the NotIn predicate on the "point_boost" field.
func PointBoostNotIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldPointBoost, vs...))
}

// PointBoostGT applies the GT predicate on the "point_boost" field.
func PointBoostGT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldPointBoost, v))
}

// PointBoostGTE applies the GTE predicate on the "point_boost" field.
func PointBoostGTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldPointBoost, v))
}

// PointBoostLT applies the LT predicate on the "point_boost" field.
func PointBoostLT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldPointBoost, v))
}

// PointBoostLTE applies the LTE predicate on the "point_boost" field.
func PointBoostLTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldPointBoost, v))
}

// RoundUnitsEQ applies the EQ predicate on the "round_units" field.
func RoundUnitsEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldRoundUnits, v))
}

// RoundUnitsNEQ applies the NEQ predicate on the "round_units" field.
func RoundUnitsNEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldRoundUnits, v))
}

// RoundUnitsIn applies the In predicate on the "round_units" field.
func RoundUnitsIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldRoundUnits, vs...))
}

// RoundUnitsNotIn applies the NotIn predicate on the "round_units" field.
func RoundUnitsNotIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldRoundUnits, vs...))
}

// RoundUnitsGT applies the GT predicate on the "round_units" field.
func RoundUnitsGT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldRoundUnits, v))
}

// RoundUnitsGTE applies the GTE predicate on the "round_units" field.
func RoundUnitsGTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldRoundUnits, v))
}

// RoundUnitsLT applies the LT predicate on the "round_units" field.
func RoundUnitsLT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldRoundUnits, v))
}

// RoundUnitsLTE applies the LTE predicate on the "round_units" field.
func RoundUnitsLTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldRoundUnits, v))
}

// RoundDelayEQ applies the EQ predicate on the "round_delay" field.
func RoundDelayEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldRoundDelay, v))
}

// RoundDelayNEQ applies the NEQ predicate on the "round_delay" field.
func RoundDelayNEQ(v int) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldRoundDelay, v))
}

// RoundDelayIn applies the In predicate on the "round_delay" field.
func RoundDelayIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldRoundDelay, vs...))
}

// RoundDelayNotIn applies the NotIn predicate on the "round_delay" field.
func RoundDelayNotIn(vs ...int) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldRoundDelay, vs...))
}

// RoundDelayGT applies the GT predicate on the "round_delay" field.
func RoundDelayGT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldRoundDelay, v))
}

// RoundDelayGTE applies the GTE predicate on the "round_delay" field.
func RoundDelayGTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldRoundDelay, v))
}

// RoundDelayLT applies the LT predicate on the "round_delay" field.
func RoundDelayLT(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldRoundDelay, v))
}

// RoundDelayLTE applies the LTE predicate on the "round_delay" field.
func RoundDelayLTE(v int) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldRoundDelay, v))
}

// ServiceIDEQ applies the EQ predicate on the "service_id" field.
func ServiceIDEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldServiceID, v))
}

// ServiceIDNEQ applies the NEQ predicate on the "service_id" field.
func ServiceIDNEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldServiceID, v))
}

// ServiceIDIn applies the In predicate on the "service_id" field.
func ServiceIDIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldServiceID, vs...))
}

// ServiceIDNotIn applies the NotIn predicate on the "service_id" field.
func ServiceIDNotIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldServiceID, vs...))
}

// ServiceIDGT applies the GT predicate on the "service_id" field.
func ServiceIDGT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldServiceID, v))
}

// ServiceIDGTE applies the GTE predicate on the "service_id" field.
func ServiceIDGTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldServiceID, v))
}

// ServiceIDLT applies the LT predicate on the "service_id" field.
func ServiceIDLT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldServiceID, v))
}

// ServiceIDLTE applies the LTE predicate on the "service_id" field.
func ServiceIDLTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldServiceID, v))
}

// ServiceIDContains applies the Contains predicate on the "service_id" field.
func ServiceIDContains(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContains(FieldServiceID, v))
}

// ServiceIDHasPrefix applies the HasPrefix predicate on the "service_id" field.
func ServiceIDHasPrefix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasPrefix(FieldServiceID, v))
}

// ServiceIDHasSuffix applies the HasSuffix predicate on the "service_id" field.
func ServiceIDHasSuffix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasSuffix(FieldServiceID, v))
}

// ServiceIDEqualFold applies the EqualFold predicate on the "service_id" field.
func ServiceIDEqualFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldServiceID, v))
}

// ServiceIDContainsFold applies the ContainsFold predicate on the "service_id" field.
func ServiceIDContainsFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldServiceID, v))
}

// HostIDEQ applies the EQ predicate on the "host_id" field.
func HostIDEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldHostID, v))
}

// HostIDNEQ applies the NEQ predicate on the "host_id" field.
func HostIDNEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldHostID, v))
}

// HostIDIn applies the In predicate on the "host_id" field.
func HostIDIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldHostID, vs...))
}

// HostIDNotIn applies the NotIn predicate on the "host_id" field.
func HostIDNotIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldHostID, vs...))
}

// HostIDGT applies the GT predicate on the "host_id" field.
func HostIDGT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldHostID, v))
}

// HostIDGTE applies the GTE predicate on the "host_id" field.
func HostIDGTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldHostID, v))
}

// HostIDLT applies the LT predicate on the "host_id" field.
func HostIDLT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldHostID, v))
}

// HostIDLTE applies the LTE predicate on the "host_id" field.
func HostIDLTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldHostID, v))
}

// HostIDContains applies the Contains predicate on the "host_id" field.
func HostIDContains(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContains(FieldHostID, v))
}

// HostIDHasPrefix applies the HasPrefix predicate on the "host_id" field.
func HostIDHasPrefix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasPrefix(FieldHostID, v))
}

// HostIDHasSuffix applies the HasSuffix predicate on the "host_id" field.
func HostIDHasSuffix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasSuffix(FieldHostID, v))
}

// HostIDEqualFold applies the EqualFold predicate on the "host_id" field.
func HostIDEqualFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldHostID, v))
}

// HostIDContainsFold applies the ContainsFold predicate on the "host_id" field.
func HostIDContainsFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldHostID, v))
}

// TeamIDEQ applies the EQ predicate on the "team_id" field.
func TeamIDEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEQ(FieldTeamID, v))
}

// TeamIDNEQ applies the NEQ predicate on the "team_id" field.
func TeamIDNEQ(v string) predicate.HostService {
	return predicate.HostService(sql.FieldNEQ(FieldTeamID, v))
}

// TeamIDIn applies the In predicate on the "team_id" field.
func TeamIDIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldIn(FieldTeamID, vs...))
}

// TeamIDNotIn applies the NotIn predicate on the "team_id" field.
func TeamIDNotIn(vs ...string) predicate.HostService {
	return predicate.HostService(sql.FieldNotIn(FieldTeamID, vs...))
}

// TeamIDGT applies the GT predicate on the "team_id" field.
func TeamIDGT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGT(FieldTeamID, v))
}

// TeamIDGTE applies the GTE predicate on the "team_id" field.
func TeamIDGTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldGTE(FieldTeamID, v))
}

// TeamIDLT applies the LT predicate on the "team_id" field.
func TeamIDLT(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLT(FieldTeamID, v))
}

// TeamIDLTE applies the LTE predicate on the "team_id" field.
func TeamIDLTE(v string) predicate.HostService {
	return predicate.HostService(sql.FieldLTE(FieldTeamID, v))
}

// TeamIDContains applies the Contains predicate on the "team_id" field.
func TeamIDContains(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContains(FieldTeamID, v))
}

// TeamIDHasPrefix applies the HasPrefix predicate on the "team_id" field.
func TeamIDHasPrefix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasPrefix(FieldTeamID, v))
}

// TeamIDHasSuffix applies the HasSuffix predicate on the "team_id" field.
func TeamIDHasSuffix(v string) predicate.HostService {
	return predicate.HostService(sql.FieldHasSuffix(FieldTeamID, v))
}

// TeamIDEqualFold applies the EqualFold predicate on the "team_id" field.
func TeamIDEqualFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldEqualFold(FieldTeamID, v))
}

// TeamIDContainsFold applies the ContainsFold predicate on the "team_id" field.
func TeamIDContainsFold(v string) predicate.HostService {
	return predicate.HostService(sql.FieldContainsFold(FieldTeamID, v))
}

// HasChecks applies the HasEdge predicate on the "checks" edge.
func HasChecks() predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChecksTable, ChecksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChecksWith applies the HasEdge predicate on the "checks" edge with a given conditions (other predicates).
func HasChecksWith(preds ...predicate.Check) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := newChecksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProperties applies the HasEdge predicate on the "properties" edge.
func HasProperties() predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PropertiesTable, PropertiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPropertiesWith applies the HasEdge predicate on the "properties" edge with a given conditions (other predicates).
func HasPropertiesWith(preds ...predicate.Property) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := newPropertiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasService applies the HasEdge predicate on the "service" edge.
func HasService() predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServiceWith applies the HasEdge predicate on the "service" edge with a given conditions (other predicates).
func HasServiceWith(preds ...predicate.Service) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := newServiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHost applies the HasEdge predicate on the "host" edge.
func HasHost() predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostTable, HostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostWith applies the HasEdge predicate on the "host" edge with a given conditions (other predicates).
func HasHostWith(preds ...predicate.Host) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := newHostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HostService) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HostService) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
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
func Not(p predicate.HostService) predicate.HostService {
	return predicate.HostService(func(s *sql.Selector) {
		p(s.Not())
	})
}
