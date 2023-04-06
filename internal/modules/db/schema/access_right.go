package schema

import "entgo.io/ent"

// AccessRight holds the schema definition for the AccessRight entity.
type AccessRight struct {
	ent.Schema
}

// Fields of the AccessRight.
func (AccessRight) Fields() []ent.Field {
	return nil
}

// Edges of the AccessRight.
func (AccessRight) Edges() []ent.Edge {
	return nil
}
