package timesheetUploader

import (
	"os"
	"testing"
)

func Test_getEnvVariable(t *testing.T) {
	envVarName := "foo"
	envVarValue := "bar"

	recovered := func() (r bool) {
		defer func() {
			if r := recover(); r != nil {
				r = true
			}
		}()
		getEnvVariable(envVarName)
		return false
	}
	if recovered() {
		t.Errorf("Retrieving not set env variable should fail, but did not")
	}

	_ = os.Setenv(envVarName, envVarValue)

	res := getEnvVariable(envVarName)
	if res != envVarValue {
		t.Errorf("Expected %v as env variable, received %v", envVarValue, res)
	}
}
