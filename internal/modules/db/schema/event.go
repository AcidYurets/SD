package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp"),
		field.String("name"),
		field.String("description").Optional().Nillable(),
		field.String("type"),
		field.Bool("is_whole_day"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).StorageKey(
			edge.Table("events_tags"), edge.Columns("event_uuid", "tag_uuid"),
		),
		edge.To("invitations", Invitation.Type).StorageKey(
			edge.Column("event_uuid"),
		),
		edge.From("creator", User.Type).
			Ref("created_events").
			Unique(),
	}
}
