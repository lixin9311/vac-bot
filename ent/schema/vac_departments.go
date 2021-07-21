package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/lixin9311/vac-bot/tokyovacapi"
)

// VacDepartment holds the schema definition for the VacDepartment entity.
type VacDepartment struct {
	ent.Schema
}

// Fields of the VacDepartment.
func (VacDepartment) Fields() []ent.Field {
	return []ent.Field{
		field.String("partition").Immutable(),
		field.Int("department_id").Immutable(),
		field.JSON("data", &tokyovacapi.Department{}),
	}
}

// Indexes of the VacDepartment.
func (VacDepartment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("partition", "department_id"),
	}
}

// Edges of the VacDepartment.
func (VacDepartment) Edges() []ent.Edge {
	return nil
}

// Mixin of the VacDepartment.
func (VacDepartment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
