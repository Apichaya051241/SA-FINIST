// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the user in the database.
	Table = "users"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "registerstores"
	// UsersInverseTable is the table name for the Registerstore entity.
	// It exists in this package in order to avoid circular dependency with the "registerstore" package.
	UsersInverseTable = "registerstores"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldName,
	FieldPassword,
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)