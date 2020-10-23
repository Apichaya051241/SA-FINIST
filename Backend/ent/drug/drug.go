// Code generated by entc, DO NOT EDIT.

package drug

const (
	// Label holds the string label denoting the drug type in the database.
	Label = "drug"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeDrugs holds the string denoting the drugs edge name in mutations.
	EdgeDrugs = "drugs"

	// Table holds the table name of the drug in the database.
	Table = "drugs"
	// DrugsTable is the table the holds the drugs relation/edge.
	DrugsTable = "registerstores"
	// DrugsInverseTable is the table name for the Registerstore entity.
	// It exists in this package in order to avoid circular dependency with the "registerstore" package.
	DrugsInverseTable = "registerstores"
	// DrugsColumn is the table column denoting the drugs relation/edge.
	DrugsColumn = "drug_drugs"
)

// Columns holds all SQL columns for drug fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)