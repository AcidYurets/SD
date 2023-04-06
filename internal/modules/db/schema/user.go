package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone"),
		field.String("login"),
		field.String("password_hash"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invitations", Invitation.Type).StorageKey(
			edge.Column("user_uuid"),
		),
		edge.To("created_events", Event.Type).StorageKey(
			edge.Column("creator_uuid"),
		),
	}
}
