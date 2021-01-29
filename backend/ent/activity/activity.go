// Code generated by entc, DO NOT EDIT.

package activity

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldACTIVITYNAME holds the string denoting the activityname field in the database.
	FieldACTIVITYNAME = "activityname"
	// FieldAdded holds the string denoting the added field in the database.
	FieldAdded = "added"
	// FieldHours holds the string denoting the hours field in the database.
	FieldHours = "hours"

	// EdgeActiStud holds the string denoting the acti_stud edge name in mutations.
	EdgeActiStud = "acti_stud"
	// EdgeActiPlace holds the string denoting the acti_place edge name in mutations.
	EdgeActiPlace = "acti_place"
	// EdgeActiAgen holds the string denoting the acti_agen edge name in mutations.
	EdgeActiAgen = "acti_agen"
	// EdgeActiYear holds the string denoting the acti_year edge name in mutations.
	EdgeActiYear = "acti_year"
	// EdgeActiTerm holds the string denoting the acti_term edge name in mutations.
	EdgeActiTerm = "acti_term"

	// Table holds the table name of the activity in the database.
	Table = "activities"
	// ActiStudTable is the table the holds the acti_stud relation/edge.
	ActiStudTable = "activities"
	// ActiStudInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	ActiStudInverseTable = "students"
	// ActiStudColumn is the table column denoting the acti_stud relation/edge.
	ActiStudColumn = "student_stud_acti"
	// ActiPlaceTable is the table the holds the acti_place relation/edge.
	ActiPlaceTable = "activities"
	// ActiPlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	ActiPlaceInverseTable = "places"
	// ActiPlaceColumn is the table column denoting the acti_place relation/edge.
	ActiPlaceColumn = "place_place_acti"
	// ActiAgenTable is the table the holds the acti_agen relation/edge.
	ActiAgenTable = "activities"
	// ActiAgenInverseTable is the table name for the Agency entity.
	// It exists in this package in order to avoid circular dependency with the "agency" package.
	ActiAgenInverseTable = "agencies"
	// ActiAgenColumn is the table column denoting the acti_agen relation/edge.
	ActiAgenColumn = "agency_agen_acti"
	// ActiYearTable is the table the holds the acti_year relation/edge.
	ActiYearTable = "activities"
	// ActiYearInverseTable is the table name for the Year entity.
	// It exists in this package in order to avoid circular dependency with the "year" package.
	ActiYearInverseTable = "years"
	// ActiYearColumn is the table column denoting the acti_year relation/edge.
	ActiYearColumn = "year_year_acti"
	// ActiTermTable is the table the holds the acti_term relation/edge.
	ActiTermTable = "activities"
	// ActiTermInverseTable is the table name for the Term entity.
	// It exists in this package in order to avoid circular dependency with the "term" package.
	ActiTermInverseTable = "terms"
	// ActiTermColumn is the table column denoting the acti_term relation/edge.
	ActiTermColumn = "term_term_acti"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldACTIVITYNAME,
	FieldAdded,
	FieldHours,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Activity type.
var ForeignKeys = []string{
	"agency_agen_acti",
	"place_place_acti",
	"student_stud_acti",
	"term_term_acti",
	"year_year_acti",
}

var (
	// ACTIVITYNAMEValidator is a validator for the "ACTIVITYNAME" field. It is called by the builders before save.
	ACTIVITYNAMEValidator func(string) error
	// HoursValidator is a validator for the "hours" field. It is called by the builders before save.
	HoursValidator func(int) error
)
