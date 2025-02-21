// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/contrib/entproto/internal/todo/ent/schema"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Joined applies equality check predicate on the "joined" field. It's identical to JoinedEQ.
func Joined(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldJoined, v))
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v uint) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPoints, v))
}

// Exp applies equality check predicate on the "exp" field. It's identical to ExpEQ.
func Exp(v uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExp, v))
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExternalID, v))
}

// CrmID applies equality check predicate on the "crm_id" field. It's identical to CrmIDEQ.
func CrmID(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCrmID, v))
}

// Banned applies equality check predicate on the "banned" field. It's identical to BannedEQ.
func Banned(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBanned, v))
}

// CustomPb applies equality check predicate on the "custom_pb" field. It's identical to CustomPbEQ.
func CustomPb(v uint8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCustomPb, v))
}

// OptNum applies equality check predicate on the "opt_num" field. It's identical to OptNumEQ.
func OptNum(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptNum, v))
}

// OptStr applies equality check predicate on the "opt_str" field. It's identical to OptStrEQ.
func OptStr(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptStr, v))
}

// OptBool applies equality check predicate on the "opt_bool" field. It's identical to OptBoolEQ.
func OptBool(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptBool, v))
}

// BigInt applies equality check predicate on the "big_int" field. It's identical to BigIntEQ.
func BigInt(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBigInt, v))
}

// BUser1 applies equality check predicate on the "b_user_1" field. It's identical to BUser1EQ.
func BUser1(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBUser1, v))
}

// HeightInCm applies equality check predicate on the "height_in_cm" field. It's identical to HeightInCmEQ.
func HeightInCm(v float32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHeightInCm, v))
}

// AccountBalance applies equality check predicate on the "account_balance" field. It's identical to AccountBalanceEQ.
func AccountBalance(v float64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccountBalance, v))
}

// Unnecessary applies equality check predicate on the "unnecessary" field. It's identical to UnnecessaryEQ.
func Unnecessary(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUnnecessary, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldType, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// JoinedEQ applies the EQ predicate on the "joined" field.
func JoinedEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldJoined, v))
}

// JoinedNEQ applies the NEQ predicate on the "joined" field.
func JoinedNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldJoined, v))
}

// JoinedIn applies the In predicate on the "joined" field.
func JoinedIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldJoined, vs...))
}

// JoinedNotIn applies the NotIn predicate on the "joined" field.
func JoinedNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldJoined, vs...))
}

// JoinedGT applies the GT predicate on the "joined" field.
func JoinedGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldJoined, v))
}

// JoinedGTE applies the GTE predicate on the "joined" field.
func JoinedGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldJoined, v))
}

// JoinedLT applies the LT predicate on the "joined" field.
func JoinedLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldJoined, v))
}

// JoinedLTE applies the LTE predicate on the "joined" field.
func JoinedLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldJoined, v))
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v uint) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPoints, v))
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v uint) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPoints, v))
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...uint) predicate.User {
	return predicate.User(sql.FieldIn(FieldPoints, vs...))
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...uint) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPoints, vs...))
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v uint) predicate.User {
	return predicate.User(sql.FieldGT(FieldPoints, v))
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v uint) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPoints, v))
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v uint) predicate.User {
	return predicate.User(sql.FieldLT(FieldPoints, v))
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v uint) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPoints, v))
}

// ExpEQ applies the EQ predicate on the "exp" field.
func ExpEQ(v uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExp, v))
}

// ExpNEQ applies the NEQ predicate on the "exp" field.
func ExpNEQ(v uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldExp, v))
}

// ExpIn applies the In predicate on the "exp" field.
func ExpIn(vs ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldExp, vs...))
}

// ExpNotIn applies the NotIn predicate on the "exp" field.
func ExpNotIn(vs ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldExp, vs...))
}

// ExpGT applies the GT predicate on the "exp" field.
func ExpGT(v uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldExp, v))
}

// ExpGTE applies the GTE predicate on the "exp" field.
func ExpGTE(v uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldExp, v))
}

// ExpLT applies the LT predicate on the "exp" field.
func ExpLT(v uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldExp, v))
}

// ExpLTE applies the LTE predicate on the "exp" field.
func ExpLTE(v uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldExp, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExternalID, v))
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldExternalID, v))
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldExternalID, vs...))
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldExternalID, vs...))
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldExternalID, v))
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldExternalID, v))
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldExternalID, v))
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldExternalID, v))
}

// CrmIDEQ applies the EQ predicate on the "crm_id" field.
func CrmIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCrmID, v))
}

// CrmIDNEQ applies the NEQ predicate on the "crm_id" field.
func CrmIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCrmID, v))
}

// CrmIDIn applies the In predicate on the "crm_id" field.
func CrmIDIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldCrmID, vs...))
}

// CrmIDNotIn applies the NotIn predicate on the "crm_id" field.
func CrmIDNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCrmID, vs...))
}

// CrmIDGT applies the GT predicate on the "crm_id" field.
func CrmIDGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldCrmID, v))
}

// CrmIDGTE applies the GTE predicate on the "crm_id" field.
func CrmIDGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCrmID, v))
}

// CrmIDLT applies the LT predicate on the "crm_id" field.
func CrmIDLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldCrmID, v))
}

// CrmIDLTE applies the LTE predicate on the "crm_id" field.
func CrmIDLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCrmID, v))
}

// BannedEQ applies the EQ predicate on the "banned" field.
func BannedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBanned, v))
}

// BannedNEQ applies the NEQ predicate on the "banned" field.
func BannedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBanned, v))
}

// CustomPbEQ applies the EQ predicate on the "custom_pb" field.
func CustomPbEQ(v uint8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCustomPb, v))
}

// CustomPbNEQ applies the NEQ predicate on the "custom_pb" field.
func CustomPbNEQ(v uint8) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCustomPb, v))
}

// CustomPbIn applies the In predicate on the "custom_pb" field.
func CustomPbIn(vs ...uint8) predicate.User {
	return predicate.User(sql.FieldIn(FieldCustomPb, vs...))
}

// CustomPbNotIn applies the NotIn predicate on the "custom_pb" field.
func CustomPbNotIn(vs ...uint8) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCustomPb, vs...))
}

// CustomPbGT applies the GT predicate on the "custom_pb" field.
func CustomPbGT(v uint8) predicate.User {
	return predicate.User(sql.FieldGT(FieldCustomPb, v))
}

// CustomPbGTE applies the GTE predicate on the "custom_pb" field.
func CustomPbGTE(v uint8) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCustomPb, v))
}

// CustomPbLT applies the LT predicate on the "custom_pb" field.
func CustomPbLT(v uint8) predicate.User {
	return predicate.User(sql.FieldLT(FieldCustomPb, v))
}

// CustomPbLTE applies the LTE predicate on the "custom_pb" field.
func CustomPbLTE(v uint8) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCustomPb, v))
}

// OptNumEQ applies the EQ predicate on the "opt_num" field.
func OptNumEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptNum, v))
}

// OptNumNEQ applies the NEQ predicate on the "opt_num" field.
func OptNumNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOptNum, v))
}

// OptNumIn applies the In predicate on the "opt_num" field.
func OptNumIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldOptNum, vs...))
}

// OptNumNotIn applies the NotIn predicate on the "opt_num" field.
func OptNumNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOptNum, vs...))
}

// OptNumGT applies the GT predicate on the "opt_num" field.
func OptNumGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldOptNum, v))
}

// OptNumGTE applies the GTE predicate on the "opt_num" field.
func OptNumGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOptNum, v))
}

// OptNumLT applies the LT predicate on the "opt_num" field.
func OptNumLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldOptNum, v))
}

// OptNumLTE applies the LTE predicate on the "opt_num" field.
func OptNumLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOptNum, v))
}

// OptNumIsNil applies the IsNil predicate on the "opt_num" field.
func OptNumIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldOptNum))
}

// OptNumNotNil applies the NotNil predicate on the "opt_num" field.
func OptNumNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldOptNum))
}

// OptStrEQ applies the EQ predicate on the "opt_str" field.
func OptStrEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptStr, v))
}

// OptStrNEQ applies the NEQ predicate on the "opt_str" field.
func OptStrNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOptStr, v))
}

// OptStrIn applies the In predicate on the "opt_str" field.
func OptStrIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldOptStr, vs...))
}

// OptStrNotIn applies the NotIn predicate on the "opt_str" field.
func OptStrNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOptStr, vs...))
}

// OptStrGT applies the GT predicate on the "opt_str" field.
func OptStrGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldOptStr, v))
}

// OptStrGTE applies the GTE predicate on the "opt_str" field.
func OptStrGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOptStr, v))
}

// OptStrLT applies the LT predicate on the "opt_str" field.
func OptStrLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldOptStr, v))
}

// OptStrLTE applies the LTE predicate on the "opt_str" field.
func OptStrLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOptStr, v))
}

// OptStrContains applies the Contains predicate on the "opt_str" field.
func OptStrContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldOptStr, v))
}

// OptStrHasPrefix applies the HasPrefix predicate on the "opt_str" field.
func OptStrHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldOptStr, v))
}

// OptStrHasSuffix applies the HasSuffix predicate on the "opt_str" field.
func OptStrHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldOptStr, v))
}

// OptStrIsNil applies the IsNil predicate on the "opt_str" field.
func OptStrIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldOptStr))
}

// OptStrNotNil applies the NotNil predicate on the "opt_str" field.
func OptStrNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldOptStr))
}

// OptStrEqualFold applies the EqualFold predicate on the "opt_str" field.
func OptStrEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldOptStr, v))
}

// OptStrContainsFold applies the ContainsFold predicate on the "opt_str" field.
func OptStrContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldOptStr, v))
}

// OptBoolEQ applies the EQ predicate on the "opt_bool" field.
func OptBoolEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptBool, v))
}

// OptBoolNEQ applies the NEQ predicate on the "opt_bool" field.
func OptBoolNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOptBool, v))
}

// OptBoolIsNil applies the IsNil predicate on the "opt_bool" field.
func OptBoolIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldOptBool))
}

// OptBoolNotNil applies the NotNil predicate on the "opt_bool" field.
func OptBoolNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldOptBool))
}

// BigIntEQ applies the EQ predicate on the "big_int" field.
func BigIntEQ(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBigInt, v))
}

// BigIntNEQ applies the NEQ predicate on the "big_int" field.
func BigIntNEQ(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBigInt, v))
}

// BigIntIn applies the In predicate on the "big_int" field.
func BigIntIn(vs ...schema.BigInt) predicate.User {
	return predicate.User(sql.FieldIn(FieldBigInt, vs...))
}

// BigIntNotIn applies the NotIn predicate on the "big_int" field.
func BigIntNotIn(vs ...schema.BigInt) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBigInt, vs...))
}

// BigIntGT applies the GT predicate on the "big_int" field.
func BigIntGT(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldGT(FieldBigInt, v))
}

// BigIntGTE applies the GTE predicate on the "big_int" field.
func BigIntGTE(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBigInt, v))
}

// BigIntLT applies the LT predicate on the "big_int" field.
func BigIntLT(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldLT(FieldBigInt, v))
}

// BigIntLTE applies the LTE predicate on the "big_int" field.
func BigIntLTE(v schema.BigInt) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBigInt, v))
}

// BigIntIsNil applies the IsNil predicate on the "big_int" field.
func BigIntIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBigInt))
}

// BigIntNotNil applies the NotNil predicate on the "big_int" field.
func BigIntNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBigInt))
}

// BUser1EQ applies the EQ predicate on the "b_user_1" field.
func BUser1EQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBUser1, v))
}

// BUser1NEQ applies the NEQ predicate on the "b_user_1" field.
func BUser1NEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBUser1, v))
}

// BUser1In applies the In predicate on the "b_user_1" field.
func BUser1In(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldBUser1, vs...))
}

// BUser1NotIn applies the NotIn predicate on the "b_user_1" field.
func BUser1NotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBUser1, vs...))
}

// BUser1GT applies the GT predicate on the "b_user_1" field.
func BUser1GT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldBUser1, v))
}

// BUser1GTE applies the GTE predicate on the "b_user_1" field.
func BUser1GTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBUser1, v))
}

// BUser1LT applies the LT predicate on the "b_user_1" field.
func BUser1LT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldBUser1, v))
}

// BUser1LTE applies the LTE predicate on the "b_user_1" field.
func BUser1LTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBUser1, v))
}

// BUser1IsNil applies the IsNil predicate on the "b_user_1" field.
func BUser1IsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBUser1))
}

// BUser1NotNil applies the NotNil predicate on the "b_user_1" field.
func BUser1NotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBUser1))
}

// HeightInCmEQ applies the EQ predicate on the "height_in_cm" field.
func HeightInCmEQ(v float32) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHeightInCm, v))
}

// HeightInCmNEQ applies the NEQ predicate on the "height_in_cm" field.
func HeightInCmNEQ(v float32) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHeightInCm, v))
}

// HeightInCmIn applies the In predicate on the "height_in_cm" field.
func HeightInCmIn(vs ...float32) predicate.User {
	return predicate.User(sql.FieldIn(FieldHeightInCm, vs...))
}

// HeightInCmNotIn applies the NotIn predicate on the "height_in_cm" field.
func HeightInCmNotIn(vs ...float32) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHeightInCm, vs...))
}

// HeightInCmGT applies the GT predicate on the "height_in_cm" field.
func HeightInCmGT(v float32) predicate.User {
	return predicate.User(sql.FieldGT(FieldHeightInCm, v))
}

// HeightInCmGTE applies the GTE predicate on the "height_in_cm" field.
func HeightInCmGTE(v float32) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHeightInCm, v))
}

// HeightInCmLT applies the LT predicate on the "height_in_cm" field.
func HeightInCmLT(v float32) predicate.User {
	return predicate.User(sql.FieldLT(FieldHeightInCm, v))
}

// HeightInCmLTE applies the LTE predicate on the "height_in_cm" field.
func HeightInCmLTE(v float32) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHeightInCm, v))
}

// AccountBalanceEQ applies the EQ predicate on the "account_balance" field.
func AccountBalanceEQ(v float64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAccountBalance, v))
}

// AccountBalanceNEQ applies the NEQ predicate on the "account_balance" field.
func AccountBalanceNEQ(v float64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAccountBalance, v))
}

// AccountBalanceIn applies the In predicate on the "account_balance" field.
func AccountBalanceIn(vs ...float64) predicate.User {
	return predicate.User(sql.FieldIn(FieldAccountBalance, vs...))
}

// AccountBalanceNotIn applies the NotIn predicate on the "account_balance" field.
func AccountBalanceNotIn(vs ...float64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAccountBalance, vs...))
}

// AccountBalanceGT applies the GT predicate on the "account_balance" field.
func AccountBalanceGT(v float64) predicate.User {
	return predicate.User(sql.FieldGT(FieldAccountBalance, v))
}

// AccountBalanceGTE applies the GTE predicate on the "account_balance" field.
func AccountBalanceGTE(v float64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAccountBalance, v))
}

// AccountBalanceLT applies the LT predicate on the "account_balance" field.
func AccountBalanceLT(v float64) predicate.User {
	return predicate.User(sql.FieldLT(FieldAccountBalance, v))
}

// AccountBalanceLTE applies the LTE predicate on the "account_balance" field.
func AccountBalanceLTE(v float64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAccountBalance, v))
}

// UnnecessaryEQ applies the EQ predicate on the "unnecessary" field.
func UnnecessaryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUnnecessary, v))
}

// UnnecessaryNEQ applies the NEQ predicate on the "unnecessary" field.
func UnnecessaryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUnnecessary, v))
}

// UnnecessaryIn applies the In predicate on the "unnecessary" field.
func UnnecessaryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUnnecessary, vs...))
}

// UnnecessaryNotIn applies the NotIn predicate on the "unnecessary" field.
func UnnecessaryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUnnecessary, vs...))
}

// UnnecessaryGT applies the GT predicate on the "unnecessary" field.
func UnnecessaryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUnnecessary, v))
}

// UnnecessaryGTE applies the GTE predicate on the "unnecessary" field.
func UnnecessaryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUnnecessary, v))
}

// UnnecessaryLT applies the LT predicate on the "unnecessary" field.
func UnnecessaryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUnnecessary, v))
}

// UnnecessaryLTE applies the LTE predicate on the "unnecessary" field.
func UnnecessaryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUnnecessary, v))
}

// UnnecessaryContains applies the Contains predicate on the "unnecessary" field.
func UnnecessaryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUnnecessary, v))
}

// UnnecessaryHasPrefix applies the HasPrefix predicate on the "unnecessary" field.
func UnnecessaryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUnnecessary, v))
}

// UnnecessaryHasSuffix applies the HasSuffix predicate on the "unnecessary" field.
func UnnecessaryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUnnecessary, v))
}

// UnnecessaryIsNil applies the IsNil predicate on the "unnecessary" field.
func UnnecessaryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUnnecessary))
}

// UnnecessaryNotNil applies the NotNil predicate on the "unnecessary" field.
func UnnecessaryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUnnecessary))
}

// UnnecessaryEqualFold applies the EqualFold predicate on the "unnecessary" field.
func UnnecessaryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUnnecessary, v))
}

// UnnecessaryContainsFold applies the ContainsFold predicate on the "unnecessary" field.
func UnnecessaryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUnnecessary, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldType, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldLabels))
}

// DeviceTypeEQ applies the EQ predicate on the "device_type" field.
func DeviceTypeEQ(v DeviceType) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeviceType, v))
}

// DeviceTypeNEQ applies the NEQ predicate on the "device_type" field.
func DeviceTypeNEQ(v DeviceType) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeviceType, v))
}

// DeviceTypeIn applies the In predicate on the "device_type" field.
func DeviceTypeIn(vs ...DeviceType) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeviceType, vs...))
}

// DeviceTypeNotIn applies the NotIn predicate on the "device_type" field.
func DeviceTypeNotIn(vs ...DeviceType) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeviceType, vs...))
}

// OmitPrefixEQ applies the EQ predicate on the "omit_prefix" field.
func OmitPrefixEQ(v OmitPrefix) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOmitPrefix, v))
}

// OmitPrefixNEQ applies the NEQ predicate on the "omit_prefix" field.
func OmitPrefixNEQ(v OmitPrefix) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOmitPrefix, v))
}

// OmitPrefixIn applies the In predicate on the "omit_prefix" field.
func OmitPrefixIn(vs ...OmitPrefix) predicate.User {
	return predicate.User(sql.FieldIn(FieldOmitPrefix, vs...))
}

// OmitPrefixNotIn applies the NotIn predicate on the "omit_prefix" field.
func OmitPrefixNotIn(vs ...OmitPrefix) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOmitPrefix, vs...))
}

// MimeTypeEQ applies the EQ predicate on the "mime_type" field.
func MimeTypeEQ(v MimeType) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMimeType, v))
}

// MimeTypeNEQ applies the NEQ predicate on the "mime_type" field.
func MimeTypeNEQ(v MimeType) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMimeType, v))
}

// MimeTypeIn applies the In predicate on the "mime_type" field.
func MimeTypeIn(vs ...MimeType) predicate.User {
	return predicate.User(sql.FieldIn(FieldMimeType, vs...))
}

// MimeTypeNotIn applies the NotIn predicate on the "mime_type" field.
func MimeTypeNotIn(vs ...MimeType) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMimeType, vs...))
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttachment applies the HasEdge predicate on the "attachment" edge.
func HasAttachment() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AttachmentTable, AttachmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentWith applies the HasEdge predicate on the "attachment" edge with a given conditions (other predicates).
func HasAttachmentWith(preds ...predicate.Attachment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAttachmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceived1 applies the HasEdge predicate on the "received_1" edge.
func HasReceived1() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, Received1Table, Received1PrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceived1With applies the HasEdge predicate on the "received_1" edge with a given conditions (other predicates).
func HasReceived1With(preds ...predicate.Attachment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReceived1Step()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPet applies the HasEdge predicate on the "pet" edge.
func HasPet() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PetTable, PetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetWith applies the HasEdge predicate on the "pet" edge with a given conditions (other predicates).
func HasPetWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkipEdge applies the HasEdge predicate on the "skip_edge" edge.
func HasSkipEdge() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SkipEdgeTable, SkipEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkipEdgeWith applies the HasEdge predicate on the "skip_edge" edge with a given conditions (other predicates).
func HasSkipEdgeWith(preds ...predicate.SkipEdgeExample) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSkipEdgeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
