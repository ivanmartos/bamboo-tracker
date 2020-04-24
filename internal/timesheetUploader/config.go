package timesheetUploader

import "os"

func getEnvVariable(envVarName string) string {
	val, exists := os.LookupEnv(envVarName)
	if !exists {
		panic("Missing required env variable " + envVarName)
	}

	return val
}
