package util

import "os"

func GetEnvOrDefault(varName string, defaultVal string) string {
	val := os.Getenv(varName)
	if val == "" {
		return defaultVal
	}
	return val
}
