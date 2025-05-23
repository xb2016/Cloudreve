// Code generated by ent, DO NOT EDIT.

package node

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cloudreve/Cloudreve/v4/ent/predicate"
	"github.com/cloudreve/Cloudreve/v4/pkg/boolset"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// Server applies equality check predicate on the "server" field. It's identical to ServerEQ.
func Server(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldServer, v))
}

// SlaveKey applies equality check predicate on the "slave_key" field. It's identical to SlaveKeyEQ.
func SlaveKey(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldSlaveKey, v))
}

// Capabilities applies equality check predicate on the "capabilities" field. It's identical to CapabilitiesEQ.
func Capabilities(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCapabilities, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldWeight, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldDeletedAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldType, vs...))
}

// ServerEQ applies the EQ predicate on the "server" field.
func ServerEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldServer, v))
}

// ServerNEQ applies the NEQ predicate on the "server" field.
func ServerNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldServer, v))
}

// ServerIn applies the In predicate on the "server" field.
func ServerIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldServer, vs...))
}

// ServerNotIn applies the NotIn predicate on the "server" field.
func ServerNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldServer, vs...))
}

// ServerGT applies the GT predicate on the "server" field.
func ServerGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldServer, v))
}

// ServerGTE applies the GTE predicate on the "server" field.
func ServerGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldServer, v))
}

// ServerLT applies the LT predicate on the "server" field.
func ServerLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldServer, v))
}

// ServerLTE applies the LTE predicate on the "server" field.
func ServerLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldServer, v))
}

// ServerContains applies the Contains predicate on the "server" field.
func ServerContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldServer, v))
}

// ServerHasPrefix applies the HasPrefix predicate on the "server" field.
func ServerHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldServer, v))
}

// ServerHasSuffix applies the HasSuffix predicate on the "server" field.
func ServerHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldServer, v))
}

// ServerIsNil applies the IsNil predicate on the "server" field.
func ServerIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldServer))
}

// ServerNotNil applies the NotNil predicate on the "server" field.
func ServerNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldServer))
}

// ServerEqualFold applies the EqualFold predicate on the "server" field.
func ServerEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldServer, v))
}

// ServerContainsFold applies the ContainsFold predicate on the "server" field.
func ServerContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldServer, v))
}

// SlaveKeyEQ applies the EQ predicate on the "slave_key" field.
func SlaveKeyEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldSlaveKey, v))
}

// SlaveKeyNEQ applies the NEQ predicate on the "slave_key" field.
func SlaveKeyNEQ(v string) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldSlaveKey, v))
}

// SlaveKeyIn applies the In predicate on the "slave_key" field.
func SlaveKeyIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldSlaveKey, vs...))
}

// SlaveKeyNotIn applies the NotIn predicate on the "slave_key" field.
func SlaveKeyNotIn(vs ...string) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldSlaveKey, vs...))
}

// SlaveKeyGT applies the GT predicate on the "slave_key" field.
func SlaveKeyGT(v string) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldSlaveKey, v))
}

// SlaveKeyGTE applies the GTE predicate on the "slave_key" field.
func SlaveKeyGTE(v string) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldSlaveKey, v))
}

// SlaveKeyLT applies the LT predicate on the "slave_key" field.
func SlaveKeyLT(v string) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldSlaveKey, v))
}

// SlaveKeyLTE applies the LTE predicate on the "slave_key" field.
func SlaveKeyLTE(v string) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldSlaveKey, v))
}

// SlaveKeyContains applies the Contains predicate on the "slave_key" field.
func SlaveKeyContains(v string) predicate.Node {
	return predicate.Node(sql.FieldContains(FieldSlaveKey, v))
}

// SlaveKeyHasPrefix applies the HasPrefix predicate on the "slave_key" field.
func SlaveKeyHasPrefix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasPrefix(FieldSlaveKey, v))
}

// SlaveKeyHasSuffix applies the HasSuffix predicate on the "slave_key" field.
func SlaveKeyHasSuffix(v string) predicate.Node {
	return predicate.Node(sql.FieldHasSuffix(FieldSlaveKey, v))
}

// SlaveKeyIsNil applies the IsNil predicate on the "slave_key" field.
func SlaveKeyIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldSlaveKey))
}

// SlaveKeyNotNil applies the NotNil predicate on the "slave_key" field.
func SlaveKeyNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldSlaveKey))
}

// SlaveKeyEqualFold applies the EqualFold predicate on the "slave_key" field.
func SlaveKeyEqualFold(v string) predicate.Node {
	return predicate.Node(sql.FieldEqualFold(FieldSlaveKey, v))
}

// SlaveKeyContainsFold applies the ContainsFold predicate on the "slave_key" field.
func SlaveKeyContainsFold(v string) predicate.Node {
	return predicate.Node(sql.FieldContainsFold(FieldSlaveKey, v))
}

// CapabilitiesEQ applies the EQ predicate on the "capabilities" field.
func CapabilitiesEQ(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldCapabilities, v))
}

// CapabilitiesNEQ applies the NEQ predicate on the "capabilities" field.
func CapabilitiesNEQ(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldCapabilities, v))
}

// CapabilitiesIn applies the In predicate on the "capabilities" field.
func CapabilitiesIn(vs ...*boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldCapabilities, vs...))
}

// CapabilitiesNotIn applies the NotIn predicate on the "capabilities" field.
func CapabilitiesNotIn(vs ...*boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldCapabilities, vs...))
}

// CapabilitiesGT applies the GT predicate on the "capabilities" field.
func CapabilitiesGT(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldCapabilities, v))
}

// CapabilitiesGTE applies the GTE predicate on the "capabilities" field.
func CapabilitiesGTE(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldCapabilities, v))
}

// CapabilitiesLT applies the LT predicate on the "capabilities" field.
func CapabilitiesLT(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldCapabilities, v))
}

// CapabilitiesLTE applies the LTE predicate on the "capabilities" field.
func CapabilitiesLTE(v *boolset.BooleanSet) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldCapabilities, v))
}

// SettingsIsNil applies the IsNil predicate on the "settings" field.
func SettingsIsNil() predicate.Node {
	return predicate.Node(sql.FieldIsNull(FieldSettings))
}

// SettingsNotNil applies the NotNil predicate on the "settings" field.
func SettingsNotNil() predicate.Node {
	return predicate.Node(sql.FieldNotNull(FieldSettings))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.Node {
	return predicate.Node(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.Node {
	return predicate.Node(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.Node {
	return predicate.Node(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.Node {
	return predicate.Node(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.Node {
	return predicate.Node(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.Node {
	return predicate.Node(sql.FieldLTE(FieldWeight, v))
}

// HasStoragePolicy applies the HasEdge predicate on the "storage_policy" edge.
func HasStoragePolicy() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StoragePolicyTable, StoragePolicyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStoragePolicyWith applies the HasEdge predicate on the "storage_policy" edge with a given conditions (other predicates).
func HasStoragePolicyWith(preds ...predicate.StoragePolicy) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := newStoragePolicyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(sql.NotPredicates(p))
}
