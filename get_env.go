package utilslibs

import (
	"log"
	"os"
)

// getEnvOrFail is a helper that crashes if an env variable is missing or empty.
func GetEnvOrFail(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return val
}
