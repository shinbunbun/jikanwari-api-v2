package tools

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoItem struct {
	ID          string
	Date        string `dynamo:"date"`
	Flag        string `dynamo:"flag,omitempty"`
	Fri         string `dynamo:"fri"`
	Mon         string `dynamo:"mon"`
	NotifyToken string `dynamo:"notifyToken,omitempty"`
	Property    string `dynamo:"property,omitempty"`
	Sat         string `dynamo:"sat"`
	Thu         string `dynamo:"thu"`
	Tue         string `dynamo:"tue"`
	Uuid        string `dynamo:"uuid,omitempty"`
	Wed         string `dynamo:"wed"`
}

const (
	DBName   = "TimeTable"
	Region   = "ap-north-east-1"
	Endpoint = "http://dynamodb-local:8000"
)

func createSession() *session.Session {
	return session.Must(session.NewSession())
}

func (*Local) GetDynamoDB() dynamo.Table {
	var db *dynamo.DB

	ses := createSession()

	db = dynamo.New(ses, &aws.Config{
		Region:      aws.String(Region),
		Endpoint:    aws.String(Endpoint),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	return db.Table(DBName)
}

func (*Prod) GetDynamoDB() dynamo.Table {
	var db *dynamo.DB

	ses := createSession()

	db = dynamo.New(ses)

	return db.Table(DBName)
}

func GetByID(id string, table *dynamo.Table) (*DynamoItem, error) {
	var readResult DynamoItem
	err := table.Get("ID", id).One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return &DynamoItem{}, err
	}
	return &readResult, nil
}

func Put(item *DynamoItem, table *dynamo.Table) error {
	return table.Put(item).Run()
}

func DeleteByID(id string, table *dynamo.Table) error {
	return table.Delete("ID", id).Run()
}
