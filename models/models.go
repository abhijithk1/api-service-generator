package models

// Inputs from the CLI for DB
type DBInputs struct {
	WrkDir        string
	ContainerName string
	DBMS          string
	DBName        string
	PsqlUser      string
	PsqlPassword  string
	TableName     string
}

// Inputs from the CLI for API
type APIInputs struct {
	WrkDir   string
	APIGroup string
}

// Table details
type InitSchema struct {
	TableName string
	WrkDir    string
}
