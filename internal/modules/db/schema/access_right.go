package schema

import (
	"calend/internal/models/access"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessRight holds the schema definition for the AccessRight entity.
type AccessRight struct {
	ent.Schema
}

// Fields of the AccessRight.
func (AccessRight) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable().
			GoType(access.Type("")).
			StorageKey("code"),
		field.String("description"),
	}
}

// Edges of the AccessRight.
func (AccessRight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invitations", Invitation.Type),
	}
}
