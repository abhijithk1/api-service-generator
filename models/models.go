package models

// Inputs from the CLI for DB
type DBInputs struct {
	WrkDir        string
	ContainerName string
	ContainerPort int
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

// Connecting Database
type DBConnection struct {
	Driver        string
	User          string
	Password      string
	DBName        string
}

//SQLC YAML File
type SQLCYAML struct {
	Version  string     `yaml:"version"`
	Packages []Packages `yaml:"packages"`
}

type Packages struct {
	Name          string `yaml:"name"`
	Path          string `yaml:"path"`
	Queries       string `yaml:"queries"`
	Schema        string `yaml:"schema"`
	Engine        string `yaml:"engine"`
	EmitInterface bool   `yaml:"emit_interface"`
}

//Migration Strcut
type Migration struct {
	DatabaseURL string
}
