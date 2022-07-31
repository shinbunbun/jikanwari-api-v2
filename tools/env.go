package tools

import "os"

type Config interface {
	GetEnv(key string) string
}

type ProdEnv struct{}
type TestEnv struct{}

func (p *ProdEnv) GetEnv(key string) string {
	return os.Getenv(key)
}

func (d *TestEnv) GetEnv(key string) string {
	return "testEnvVar"
}
