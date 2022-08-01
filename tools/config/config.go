package config

import (
	"github.com/guregu/dynamo"
)

type Config interface {
	GetEnv(key string) string
	GetDynamoDB() dynamo.Table
}

type Local struct{}
type Prod struct{}
