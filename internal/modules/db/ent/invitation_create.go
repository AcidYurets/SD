// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/modules/db/ent/accessright"
	"calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/user"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvitationCreate is the builder for creating a Invitation entity.
type InvitationCreate struct {
	config
	mutation *InvitationMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (ic *InvitationCreate) SetID(s string) *InvitationCreate {
	ic.mutation.SetID(s)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableID(s *string) *InvitationCreate {
	if s != nil {
		ic.SetID(*s)
	}
	return ic
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (ic *InvitationCreate) SetEventID(id string) *InvitationCreate {
	ic.mutation.SetEventID(id)
	return ic
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (ic *InvitationCreate) SetNillableEventID(id *string) *InvitationCreate {
	if id != nil {
		ic = ic.SetEventID(*id)
	}
	return ic
}

// SetEvent sets the "event" edge to the Event entity.
func (ic *InvitationCreate) SetEvent(e *Event) *InvitationCreate {
	return ic.SetEventID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ic *InvitationCreate) SetUserID(id string) *InvitationCreate {
	ic.mutation.SetUserID(id)
	return ic
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ic *InvitationCreate) SetNillableUserID(id *string) *InvitationCreate {
	if id != nil {
		ic = ic.SetUserID(*id)
	}
	return ic
}

// SetUser sets the "user" edge to the User entity.
func (ic *InvitationCreate) SetUser(u *User) *InvitationCreate {
	return ic.SetUserID(u.ID)
}

// SetAccessRightID sets the "access_right" edge to the AccessRight entity by ID.
func (ic *InvitationCreate) SetAccessRightID(id string) *InvitationCreate {
	ic.mutation.SetAccessRightID(id)
	return ic
}

// SetNillableAccessRightID sets the "access_right" edge to the AccessRight entity by ID if the given value is not nil.
func (ic *InvitationCreate) SetNillableAccessRightID(id *string) *InvitationCreate {
	if id != nil {
		ic = ic.SetAccessRightID(*id)
	}
	return ic
}

// SetAccessRight sets the "access_right" edge to the AccessRight entity.
func (ic *InvitationCreate) SetAccessRight(a *AccessRight) *InvitationCreate {
	return ic.SetAccessRightID(a.ID)
}

// Mutation returns the InvitationMutation object of the builder.
func (ic *InvitationCreate) Mutation() *InvitationMutation {
	return ic.mutation
}

// Save creates the Invitation in the database.
func (ic *InvitationCreate) Save(ctx context.Context) (*Invitation, error) {
	ic.defaults()
	return withHooks[*Invitation, InvitationMutation](ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvitationCreate) SaveX(ctx context.Context) *Invitation {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvitationCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvitationCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvitationCreate) defaults() {
	if _, ok := ic.mutation.ID(); !ok {
		v := invitation.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvitationCreate) check() error {
	return nil
}

func (ic *InvitationCreate) sqlSave(ctx context.Context) (*Invitation, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Invitation.ID type: %T", _spec.ID.Value)
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvitationCreate) createSpec() (*Invitation, *sqlgraph.CreateSpec) {
	var (
		_node = &Invitation{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invitation.Table, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := ic.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.EventTable,
			Columns: []string{invitation.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_uuid = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.UserTable,
			Columns: []string{invitation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_uuid = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AccessRightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.AccessRightTable,
			Columns: []string{invitation.AccessRightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessright.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.access_right_code = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvitationCreateBulk is the builder for creating many Invitation entities in bulk.
type InvitationCreateBulk struct {
	config
	builders []*InvitationCreate
}

// Save creates the Invitation entities in the database.
func (icb *InvitationCreateBulk) Save(ctx context.Context) ([]*Invitation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invitation, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvitationCreateBulk) SaveX(ctx context.Context) []*Invitation {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvitationCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}