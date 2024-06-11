package common

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/abhijithk1/api-service-generator/models"
	"gopkg.in/yaml.v3"
)

var (
	InitialDirectories = []string{"/api", "/api/v1", "/api/v1/mw", "/api/v1/mw/cors", "/api/v1/mw/auth", "/pkg", "/pkg/db", "/pkg/db/migrations", "/pkg/db/query", "/utils"}
	DependentPackages  = []string{"github.com/gin-gonic/gin", "github.com/IBM/alchemy-logging/src/go/alog", "github.com/golang-migrate/migrate/v4", "github.com/gin-contrib/cors", "github.com/spf13/viper", "github.com/stretchr/testify/mock"}
	MarshalYAML        = yaml.Marshal
)

// CommandExecutor defines the interface for executing commands
type CommandExecutor interface {
	ExecuteCmds(cmdStr string, cmdArgs []string, workDir string) ([]byte, error)
	CreateDirectory(path string) error
	CreateFileAndItsContent(fileName string, fileData interface{}, content string) error
}

// DefaultExecutor is the default implementation of CommandExecutor
var DefaultExecutor CommandExecutor = &realCommandExecutor{}

type realCommandExecutor struct{}

func (r *realCommandExecutor) ExecuteCmds(cmdStr string, cmdArgs []string, workDir string) ([]byte, error) {
	cmd := exec.Command(cmdStr, cmdArgs...)
	cmd.Dir = workDir
	return cmd.CombinedOutput()
}

func (r *realCommandExecutor) CreateDirectory(path string) error {
	return os.MkdirAll(path, 0777)
}

func ExecuteGoMod(path, name string) error {
	modName := fmt.Sprintf("%s/%s", path, name)

	_, err := ExecuteCmds("go", []string{"mod", "init", modName}, name)
	if err != nil {
		return fmt.Errorf("error in initialising go module: %w", err)
	}

	return nil
}

func ExecuteGoGets(wrkDir string) error {
	for _, pkg := range DependentPackages {
		_, err := ExecuteCmds("go", []string{"get", pkg}, wrkDir)
		if err != nil {
			return fmt.Errorf("error in downloading go package %s : %w", pkg, err)
		}
	}

	return nil
}

func ExecuteCmds(cmdStr string, cmdArgs []string, workDir string) ([]byte, error) {
	return DefaultExecutor.ExecuteCmds(cmdStr, cmdArgs, workDir)
}

// Function that Creates Directory of specific path
func CreateDirectory(path string) error {
	return DefaultExecutor.CreateDirectory(path)
}

// CreateFileAndItsContent creates a file with 0777 permissions and writes content to it using a template
func (r *realCommandExecutor) CreateFileAndItsContent(fileName string, fileData interface{}, content string) error {

	// Create a new file with 0777 permissions
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", fileName, err)
	}
	defer file.Close()

	// Parse the template
	tmpl, err := template.New("file").Parse(content)
	if err != nil {
		return fmt.Errorf("error parsing template : %w", err)

	}

	// Execute the template and write to the file
	err = tmpl.Execute(file, fileData)
	if err != nil {
		return fmt.Errorf("error executing template : %w", err)
	}

	fmt.Printf("File %s created successfully\n", fileName)

	return nil
}

func CreateFileAndItsContent(fileName string, fileData interface{}, content string) error {
	return DefaultExecutor.CreateFileAndItsContent(fileName, fileData, content)
}

func Initialise(path string, dbInputs *models.DBInputs) (err error){

	fmt.Printf("\n\n*** Creating the Service Directory %s ***\n", dbInputs.WrkDir)
	err = CreateDirectory(dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Printf("\n*** Successfully created the Service Directory %s ***\n", dbInputs.WrkDir)

	fmt.Println("*** Creating go.mod ***")
	err = ExecuteGoMod(path, dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n*** Successfully created go.mod ***")

	appendDriverPackage(dbInputs)
	fmt.Println("\n*** Updating go packages ***")
	err = ExecuteGoGets(dbInputs.WrkDir)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("\n*** Successfully updated go packages ***")

	fmt.Println("\n*** Creating the initial Directories ***")
	for _, dir := range InitialDirectories {
		dir = dbInputs.WrkDir + dir
		err = CreateDirectory(dir)
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
	}
	fmt.Println("\n*** Successfully Created the initial Directories ***")

	return nil
}

// ToCamelCase converts a snake_case string to CamelCase
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = capitalize(part)
	}
	return strings.Join(parts, "")
}

// capitalize capitalizes the first letter of a string
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func IsValidString(s string) bool {
	// Define the regular expression pattern to match only alphanumeric characters and underscores.
	var validStringPattern = `^[a-zA-Z0-9_]*$`

	// Compile the regular expression.
	re := regexp.MustCompile(validStringPattern)

	// Check if the string matches the pattern.
	return re.MatchString(s)
}

func appendDriverPackage(dbInputs *models.DBInputs) {
	switch dbInputs.DBMS {
	case "postgres":
		DependentPackages = append(DependentPackages, "github.com/lib/pq")
		dbInputs.DriverPackage = "github.com/lib/pq"
	case "mysql":
		DependentPackages = append(DependentPackages, "github.com/go-sql-driver/mysql")
		dbInputs.DriverPackage = "github.com/go-sql-driver/mysql"
	}
}
