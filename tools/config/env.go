package config

import "os"

func (*Local) GetEnv(key string) string {
	return "testEnvVar"
}

func (*Prod) GetEnv(key string) string {
	return os.Getenv(key)
}
