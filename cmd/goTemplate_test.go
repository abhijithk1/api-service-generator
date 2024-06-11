package cmd

import (
	"bufio"
	"strings"
	"testing"

	"github.com/abhijithk1/api-service-generator/models"
	"github.com/stretchr/testify/assert"
)

// Mock validation function that always returns true
func alwaysValid(input string) bool {
	return true
}

// Mock validation function that returns false for specific invalid inputs
func mockValidation(input string) bool {
	return input != "invalid"
}

func TestPromptForInput(t *testing.T) {
	tests := []struct{
		name         string
		input        string
		prompt       string
		defaultValue string
		validation   func(string) bool
		expected     string
	}{
		{
			name:         "Valid Input",
			input:        "validinput\n",
			prompt:       "Enter something: ",
			defaultValue: "default",
			validation:   alwaysValid,
			expected:     "validinput",
		},
		{
			name:         "Empty Input",
			input:        "\n",
			prompt:       "Enter something: ",
			defaultValue: "default",
			validation:   alwaysValid,
			expected:     "default",
		},
		{
			name:         "Invalid Input Once Then Valid",
			input:        "invalid\nvalidinput\n",
			prompt:       "Enter something: ",
			defaultValue: "default",
			validation:   mockValidation,
			expected:     "validinput",
		},
		{
			name:         "Invalid Input Always",
			input:        "invalid\ninvalid\n",
			prompt:       "Enter something: ",
			defaultValue: "default",
			validation:   mockValidation,
			expected:     "default", // Will use default after loop (simulated)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			result := promptForInput(reader, tt.prompt, tt.defaultValue, tt.validation)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPromptForInt(t *testing.T) {
	tests := []struct{
		name         string
		input        string
		prompt       string
		defaultValue int
		expected     int
	}{
		{
			name:         "Valid Input",
			input:        "12\n",
			prompt:       "Enter something: ",
			defaultValue: 1,
			expected:     12,
		},
		{
			name:         "Empty Input",
			input:        "\n",
			prompt:       "Enter something: ",
			defaultValue: 1,
			expected:     1,
		},
		{
			name:         "Invalid Input Once Then Valid",
			input:        "invalid\nabc\n",
			prompt:       "Enter something: ",
			defaultValue: 1,
			expected:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			result := promptForInt(reader, tt.prompt, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDriverInputs(t *testing.T) {
	tests := []struct {
		name       string
		dbms       string
		expected   *models.DBInputs
		expectErr  bool
	}{
		{
			name: "Postgres Input",
			dbms: "postgres",
			expected: &models.DBInputs{
				DBMS: "postgres",
				Postgres: models.PostgresDriver{
					PsqlUser:     "postgres",
					PsqlPassword: "password",
				},
			},
			expectErr: false,
		},
		{
			name: "MySQL Input",
			dbms: "mysql",
			expected: &models.DBInputs{
				DBMS: "mysql",
				MySQL: models.MySQLDriver{
					MysqlRootPassword: "my-root-secret",
					MysqlUser:         "mysql",
					MysqlPassword:     "password",
				},
			},
			expectErr: false,
		},
		{
			name: "Unsupported DBMS",
			dbms: "unsupported",
			expected: &models.DBInputs{
				DBMS: "unsupported",
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbInputs := &models.DBInputs{
				DBMS: tt.dbms,
			}
			reader := bufio.NewReader(strings.NewReader(""))
			err := driverInputs(reader, dbInputs)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, dbInputs)
			}
		})
	}
}
