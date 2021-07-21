// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lixin9311/vac-bot/ent/schema"
	"github.com/lixin9311/vac-bot/ent/vacuser"
	"github.com/lixin9311/vac-bot/tokyovacapi"
)

// VacUserCreate is the builder for creating a VacUser entity.
type VacUserCreate struct {
	config
	mutation *VacUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vuc *VacUserCreate) SetCreatedAt(t time.Time) *VacUserCreate {
	vuc.mutation.SetCreatedAt(t)
	return vuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableCreatedAt(t *time.Time) *VacUserCreate {
	if t != nil {
		vuc.SetCreatedAt(*t)
	}
	return vuc
}

// SetUpdatedAt sets the "updated_at" field.
func (vuc *VacUserCreate) SetUpdatedAt(t time.Time) *VacUserCreate {
	vuc.mutation.SetUpdatedAt(t)
	return vuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableUpdatedAt(t *time.Time) *VacUserCreate {
	if t != nil {
		vuc.SetUpdatedAt(*t)
	}
	return vuc
}

// SetSnsID sets the "sns_id" field.
func (vuc *VacUserCreate) SetSnsID(s string) *VacUserCreate {
	vuc.mutation.SetSnsID(s)
	return vuc
}

// SetPartition sets the "partition" field.
func (vuc *VacUserCreate) SetPartition(s string) *VacUserCreate {
	vuc.mutation.SetPartition(s)
	return vuc
}

// SetNillablePartition sets the "partition" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillablePartition(s *string) *VacUserCreate {
	if s != nil {
		vuc.SetPartition(*s)
	}
	return vuc
}

// SetRangeKey sets the "range_key" field.
func (vuc *VacUserCreate) SetRangeKey(s string) *VacUserCreate {
	vuc.mutation.SetRangeKey(s)
	return vuc
}

// SetNillableRangeKey sets the "range_key" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableRangeKey(s *string) *VacUserCreate {
	if s != nil {
		vuc.SetRangeKey(*s)
	}
	return vuc
}

// SetPassword sets the "password" field.
func (vuc *VacUserCreate) SetPassword(s string) *VacUserCreate {
	vuc.mutation.SetPassword(s)
	return vuc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillablePassword(s *string) *VacUserCreate {
	if s != nil {
		vuc.SetPassword(*s)
	}
	return vuc
}

// SetToken sets the "token" field.
func (vuc *VacUserCreate) SetToken(s string) *VacUserCreate {
	vuc.mutation.SetToken(s)
	return vuc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableToken(s *string) *VacUserCreate {
	if s != nil {
		vuc.SetToken(*s)
	}
	return vuc
}

// SetWatcherEnabled sets the "watcher_enabled" field.
func (vuc *VacUserCreate) SetWatcherEnabled(b bool) *VacUserCreate {
	vuc.mutation.SetWatcherEnabled(b)
	return vuc
}

// SetNillableWatcherEnabled sets the "watcher_enabled" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableWatcherEnabled(b *bool) *VacUserCreate {
	if b != nil {
		vuc.SetWatcherEnabled(*b)
	}
	return vuc
}

// SetReserveConfig sets the "reserve_config" field.
func (vuc *VacUserCreate) SetReserveConfig(sc *schema.ReserveConfig) *VacUserCreate {
	vuc.mutation.SetReserveConfig(sc)
	return vuc
}

// SetReserveEnabled sets the "reserve_enabled" field.
func (vuc *VacUserCreate) SetReserveEnabled(b bool) *VacUserCreate {
	vuc.mutation.SetReserveEnabled(b)
	return vuc
}

// SetNillableReserveEnabled sets the "reserve_enabled" field if the given value is not nil.
func (vuc *VacUserCreate) SetNillableReserveEnabled(b *bool) *VacUserCreate {
	if b != nil {
		vuc.SetReserveEnabled(*b)
	}
	return vuc
}

// SetReservations sets the "reservations" field.
func (vuc *VacUserCreate) SetReservations(tl *tokyovacapi.ReservationList) *VacUserCreate {
	vuc.mutation.SetReservations(tl)
	return vuc
}

// Mutation returns the VacUserMutation object of the builder.
func (vuc *VacUserCreate) Mutation() *VacUserMutation {
	return vuc.mutation
}

// Save creates the VacUser in the database.
func (vuc *VacUserCreate) Save(ctx context.Context) (*VacUser, error) {
	var (
		err  error
		node *VacUser
	)
	vuc.defaults()
	if len(vuc.hooks) == 0 {
		if err = vuc.check(); err != nil {
			return nil, err
		}
		node, err = vuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VacUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuc.check(); err != nil {
				return nil, err
			}
			vuc.mutation = mutation
			node, err = vuc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuc.hooks) - 1; i >= 0; i-- {
			mut = vuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vuc *VacUserCreate) SaveX(ctx context.Context) *VacUser {
	v, err := vuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (vuc *VacUserCreate) defaults() {
	if _, ok := vuc.mutation.CreatedAt(); !ok {
		v := vacuser.DefaultCreatedAt()
		vuc.mutation.SetCreatedAt(v)
	}
	if _, ok := vuc.mutation.UpdatedAt(); !ok {
		v := vacuser.DefaultUpdatedAt()
		vuc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuc *VacUserCreate) check() error {
	if _, ok := vuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := vuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := vuc.mutation.SnsID(); !ok {
		return &ValidationError{Name: "sns_id", err: errors.New("ent: missing required field \"sns_id\"")}
	}
	return nil
}

func (vuc *VacUserCreate) sqlSave(ctx context.Context) (*VacUser, error) {
	_node, _spec := vuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vuc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vuc *VacUserCreate) createSpec() (*VacUser, *sqlgraph.CreateSpec) {
	var (
		_node = &VacUser{config: vuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vacuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vacuser.FieldID,
			},
		}
	)
	if value, ok := vuc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vacuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vuc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vacuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vuc.mutation.SnsID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vacuser.FieldSnsID,
		})
		_node.SnsID = value
	}
	if value, ok := vuc.mutation.Partition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vacuser.FieldPartition,
		})
		_node.Partition = value
	}
	if value, ok := vuc.mutation.RangeKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vacuser.FieldRangeKey,
		})
		_node.RangeKey = value
	}
	if value, ok := vuc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vacuser.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := vuc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vacuser.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := vuc.mutation.WatcherEnabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vacuser.FieldWatcherEnabled,
		})
		_node.WatcherEnabled = value
	}
	if value, ok := vuc.mutation.ReserveConfig(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vacuser.FieldReserveConfig,
		})
		_node.ReserveConfig = value
	}
	if value, ok := vuc.mutation.ReserveEnabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vacuser.FieldReserveEnabled,
		})
		_node.ReserveEnabled = value
	}
	if value, ok := vuc.mutation.Reservations(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vacuser.FieldReservations,
		})
		_node.Reservations = value
	}
	return _node, _spec
}

// VacUserCreateBulk is the builder for creating many VacUser entities in bulk.
type VacUserCreateBulk struct {
	config
	builders []*VacUserCreate
}

// Save creates the VacUser entities in the database.
func (vucb *VacUserCreateBulk) Save(ctx context.Context) ([]*VacUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vucb.builders))
	nodes := make([]*VacUser, len(vucb.builders))
	mutators := make([]Mutator, len(vucb.builders))
	for i := range vucb.builders {
		func(i int, root context.Context) {
			builder := vucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VacUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vucb *VacUserCreateBulk) SaveX(ctx context.Context) []*VacUser {
	v, err := vucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
