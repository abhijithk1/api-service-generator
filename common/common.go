package common

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

var (
	InitialDirectories = []string{"/api", "/api/v1", "/pkg", "/pkg/db", "/pkg/db/migrations", "/pkg/db/query"}
	DependentPackages  = []string{"github.com/gin-gonic/gin", "github.com/IBM/alchemy-logging/src/go/alog", "github.com/spf13/viper"}
)

// CommandExecutor defines the interface for executing commands
type CommandExecutor interface {
	ExecuteCmds(cmdStr string, cmdArgs []string) ([]byte, error)
	CreateDirectory(path string) error
	CreateFileAndItsContent(fileName string, fileData interface{}, content string) error
}

// DefaultExecutor is the default implementation of CommandExecutor
var DefaultExecutor CommandExecutor = &realCommandExecutor{}

type realCommandExecutor struct{}

func (r *realCommandExecutor) ExecuteCmds(cmdStr string, cmdArgs []string) ([]byte, error) {
	return exec.Command(cmdStr, cmdArgs...).Output()
}

func (r *realCommandExecutor) CreateDirectory(path string) error {
	return os.MkdirAll(path, 0777)
}

func ExecuteGoMod(path, name string) error {
	modName := fmt.Sprintf("%s/%s", path, name)

	_, err := ExecuteCmds("go", []string{"mod", "init", modName})
	if err != nil {
		return fmt.Errorf("error in initialising go module: %w", err)
	}

	return nil
}

func ExecuteGoGets() error {
	for _, pkg := range DependentPackages {
		_, err := ExecuteCmds("go", []string{"get", pkg})
		if err != nil {
			return fmt.Errorf("error in downloading go package %s : %w", pkg, err)
		}
	}

	return nil
}

func ExecuteCmds(cmdStr string, cmdArgs []string) ([]byte, error) {
	return DefaultExecutor.ExecuteCmds(cmdStr, cmdArgs)
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

func Initialise(path, name string) {

	fmt.Printf("*** Creating the Service Directory %s ***\n", name)
	err := CreateDirectory(name)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Printf("*** Successfully created the Service Directory %s ***\n", name)

	_, err = ExecuteCmds("cd", []string{name})
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("*** Creating go.mod ***")
	err = ExecuteGoMod(path, name)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("*** Successfully created go.mod ***")

	fmt.Println("*** Updating go packages ***")
	err = ExecuteGoGets()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("*** Successfully updated go packages ***")

	fmt.Println("*** Creating the initial Directories ***")
	for _, dir := range InitialDirectories {
		err = CreateDirectory(dir)
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
	}
	fmt.Println("*** Successfully Created the initial Directories ***")
}
