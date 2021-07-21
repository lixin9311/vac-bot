package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/lixin9311/vac-bot/tokyovacapi"
)

type ReserveConfig struct {
	StartOffset time.Duration
	EndOffset   time.Duration
	StartDate   time.Time
}

// VacUser holds the schema definition for the VacUser entity.
type VacUser struct {
	ent.Schema
}

// Fields of the VacUser.
func (VacUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("sns_id").Immutable().Unique(),
		field.String("partition").Optional(),
		field.String("range_key").Optional(),
		field.String("password").Optional(),
		field.String("token").Optional(),
		field.Bool("watcher_enabled").Optional(),
		field.JSON("reserve_config", &ReserveConfig{}).Optional(),
		field.Bool("reserve_enabled").Optional(),
		field.Bool("reserved").Optional(),
		field.JSON("reservations", &tokyovacapi.ReservationList{}).Optional(),
	}
}

// Indexes of the VacUser.
func (VacUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sns_id", "partition"),
	}
}

// Edges of the VacUser.
func (VacUser) Edges() []ent.Edge {
	return nil
}

// Mixin of the VacUser.
func (VacUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
