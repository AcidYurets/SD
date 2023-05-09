// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessRightsColumns holds the columns for the "access_rights" table.
	AccessRightsColumns = []*schema.Column{
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// AccessRightsTable holds the schema information for the "access_rights" table.
	AccessRightsTable = &schema.Table{
		Name:       "access_rights",
		Columns:    AccessRightsColumns,
		PrimaryKey: []*schema.Column{AccessRightsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeString, Default: schema.Expr("uuid_generate_v4()"), SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString},
		{Name: "is_whole_day", Type: field.TypeBool},
		{Name: "creator_uuid", Type: field.TypeString, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_users_created_events",
				Columns:    []*schema.Column{EventsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InvitationsColumns holds the columns for the "invitations" table.
	InvitationsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeString, Default: schema.Expr("uuid_generate_v4()"), SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "access_right_code", Type: field.TypeString},
		{Name: "event_uuid", Type: field.TypeString, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "user_uuid", Type: field.TypeString, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// InvitationsTable holds the schema information for the "invitations" table.
	InvitationsTable = &schema.Table{
		Name:       "invitations",
		Columns:    InvitationsColumns,
		PrimaryKey: []*schema.Column{InvitationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invitations_access_rights_invitations",
				Columns:    []*schema.Column{InvitationsColumns[1]},
				RefColumns: []*schema.Column{AccessRightsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invitations_events_invitations",
				Columns:    []*schema.Column{InvitationsColumns[2]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "invitations_users_invitations",
				Columns:    []*schema.Column{InvitationsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "invitation_event_uuid_user_uuid",
				Unique:  true,
				Columns: []*schema.Column{InvitationsColumns[2], InvitationsColumns[3]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeString, Default: schema.Expr("uuid_generate_v4()"), SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeString, Default: schema.Expr("uuid_generate_v4()"), SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// EventsTagsColumns holds the columns for the "events_tags" table.
	EventsTagsColumns = []*schema.Column{
		{Name: "event_uuid", Type: field.TypeString, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "tag_uuid", Type: field.TypeString, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// EventsTagsTable holds the schema information for the "events_tags" table.
	EventsTagsTable = &schema.Table{
		Name:       "events_tags",
		Columns:    EventsTagsColumns,
		PrimaryKey: []*schema.Column{EventsTagsColumns[0], EventsTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_tags_event_uuid",
				Columns:    []*schema.Column{EventsTagsColumns[0]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "events_tags_tag_uuid",
				Columns:    []*schema.Column{EventsTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessRightsTable,
		EventsTable,
		InvitationsTable,
		TagsTable,
		UsersTable,
		EventsTagsTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = UsersTable
	InvitationsTable.ForeignKeys[0].RefTable = AccessRightsTable
	InvitationsTable.ForeignKeys[1].RefTable = EventsTable
	InvitationsTable.ForeignKeys[2].RefTable = UsersTable
	EventsTagsTable.ForeignKeys[0].RefTable = EventsTable
	EventsTagsTable.ForeignKeys[1].RefTable = TagsTable
}
