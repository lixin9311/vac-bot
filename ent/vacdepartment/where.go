// Code generated by entc, DO NOT EDIT.

package vacdepartment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lixin9311/vac-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Partition applies equality check predicate on the "partition" field. It's identical to PartitionEQ.
func Partition(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPartition), v))
	})
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// PartitionEQ applies the EQ predicate on the "partition" field.
func PartitionEQ(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPartition), v))
	})
}

// PartitionNEQ applies the NEQ predicate on the "partition" field.
func PartitionNEQ(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPartition), v))
	})
}

// PartitionIn applies the In predicate on the "partition" field.
func PartitionIn(vs ...string) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPartition), v...))
	})
}

// PartitionNotIn applies the NotIn predicate on the "partition" field.
func PartitionNotIn(vs ...string) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPartition), v...))
	})
}

// PartitionGT applies the GT predicate on the "partition" field.
func PartitionGT(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPartition), v))
	})
}

// PartitionGTE applies the GTE predicate on the "partition" field.
func PartitionGTE(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPartition), v))
	})
}

// PartitionLT applies the LT predicate on the "partition" field.
func PartitionLT(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPartition), v))
	})
}

// PartitionLTE applies the LTE predicate on the "partition" field.
func PartitionLTE(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPartition), v))
	})
}

// PartitionContains applies the Contains predicate on the "partition" field.
func PartitionContains(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPartition), v))
	})
}

// PartitionHasPrefix applies the HasPrefix predicate on the "partition" field.
func PartitionHasPrefix(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPartition), v))
	})
}

// PartitionHasSuffix applies the HasSuffix predicate on the "partition" field.
func PartitionHasSuffix(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPartition), v))
	})
}

// PartitionEqualFold applies the EqualFold predicate on the "partition" field.
func PartitionEqualFold(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPartition), v))
	})
}

// PartitionContainsFold applies the ContainsFold predicate on the "partition" field.
func PartitionContainsFold(v string) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPartition), v))
	})
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartmentID), v))
	})
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepartmentID), v))
	})
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...int) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepartmentID), v...))
	})
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...int) predicate.VacDepartment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VacDepartment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepartmentID), v...))
	})
}

// DepartmentIDGT applies the GT predicate on the "department_id" field.
func DepartmentIDGT(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepartmentID), v))
	})
}

// DepartmentIDGTE applies the GTE predicate on the "department_id" field.
func DepartmentIDGTE(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepartmentID), v))
	})
}

// DepartmentIDLT applies the LT predicate on the "department_id" field.
func DepartmentIDLT(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepartmentID), v))
	})
}

// DepartmentIDLTE applies the LTE predicate on the "department_id" field.
func DepartmentIDLTE(v int) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepartmentID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VacDepartment) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VacDepartment) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
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
func Not(p predicate.VacDepartment) predicate.VacDepartment {
	return predicate.VacDepartment(func(s *sql.Selector) {
		p(s.Not())
	})
}
