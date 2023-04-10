package schema

import (
	"calend/internal/models/access"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Invitation holds the schema definition for the Invitation entity.
type Invitation struct {
	ent.Schema
}

func (Invitation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
	}
}

// Fields of the Invitation.
func (Invitation) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("event_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("access_right_code").GoType(access.Type("")),
	}
}

// Edges of the Invitation.
func (Invitation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("invitations").
			Field("event_uuid").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("invitations").
			Field("user_uuid").
			Unique().
			Required(),
		edge.From("access_right", AccessRight.Type).
			Ref("invitations").
			Field("access_right_code").
			Unique().
			Required(),
	}
}

func (Invitation) Indexes() []ent.Index {
	return []ent.Index{
		// Создаем уникальный индекс на пару атрибутов
		index.Fields("event_uuid", "user_uuid").
			Unique(),
	}
}
