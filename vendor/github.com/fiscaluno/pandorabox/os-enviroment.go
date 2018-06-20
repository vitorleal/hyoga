package pandorabox

import (
	"os"
)

// GetOSEnvironment get a env var or set to default
func GetOSEnvironment(env string, defaultEnv string) string {
	if e := os.Getenv(env); e != "" {
		return os.Getenv(env)
	}
	return defaultEnv
}
