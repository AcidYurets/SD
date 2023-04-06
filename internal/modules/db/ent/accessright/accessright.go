// Code generated by ent, DO NOT EDIT.

package accessright

const (
	// Label holds the string label denoting the accessright type in the database.
	Label = "access_right"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// InvitationFieldID holds the string denoting the ID field of the Invitation.
	InvitationFieldID = "uuid"
	// Table holds the table name of the accessright in the database.
	Table = "access_rights"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "access_right_code"
)

// Columns holds all SQL columns for accessright fields.
var Columns = []string{
	FieldID,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}