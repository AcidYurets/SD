// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsWholeDay holds the string denoting the is_whole_day field in the database.
	FieldIsWholeDay = "is_whole_day"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// Table holds the table name of the event in the database.
	Table = "events"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "events_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "event_uuid"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "events"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_uuid"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTimestamp,
	FieldName,
	FieldDescription,
	FieldType,
	FieldIsWholeDay,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"creator_uuid",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"event_uuid", "tag_uuid"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "calend/internal/modules/db/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)