package tools

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func createSession() *session.Session {
	return session.Must(session.NewSession())
}

func (*Local) GetDynamoDB() dynamo.Table {
	var db *dynamo.DB

	ses := createSession()

	db = dynamo.New(ses, &aws.Config{
		Region:      aws.String("ap-north-east-1"),
		Endpoint:    aws.String("http://dynamodb-local:8000"),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	return db.Table("UBIC-FOOD")
}

func (*Prod) GetDynamoDB() dynamo.Table {
	var db *dynamo.DB

	ses := createSession()

	db = dynamo.New(ses)

	return db.Table("UBIC-FOOD")
}
