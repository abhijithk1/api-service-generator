package models

// Inputs from the CLI for DB
type DBInputs struct {
	WrkDir        string
	GoModule      string
	ContainerName string
	ContainerPort int
	DBMS          string
	DBName        string
	DriverPackage string
	Postgres      PostgresDriver
	MySQL         MySQLDriver
	TableName     string
}

// Postgres
type PostgresDriver struct {
	PsqlUser     string
	PsqlPassword string
}

//MySQL
type MySQLDriver struct {
	MysqlRootPassword string
	MysqlUser         string
	MysqlPassword     string
}

// Inputs from the CLI for API
type APIInputs struct {
	WrkDir         string
	GoModule       string
	APIGroup       string
	APIGroupTitle  string
	TableName      string
	TableNameTitle string
}

// Table details
type InitSchema struct {
	TableName string
	WrkDir    string
}


// SQLC YAML File
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

type UnitTestData struct {
	Package     string
	WrkDir      string
	APIName     string
	TableName   string
	TableObject string
}
