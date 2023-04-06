package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invitations", Invitation.Type).StorageKey(
			edge.Column("user_uuid"),
		),
	}
}
