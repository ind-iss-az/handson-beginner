package main

import (
	"os"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/google/uuid"
)

type Response struct {
    Date		string	`json:"Date"`
	Hostname	string	`json:"Hostname"`
	MessageStr	string	`json:"MessageStr"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// DynamoDBへの認証情報取得する
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	uuidObj, _ := uuid.NewRandom()
	messageId := uuidObj.String()
	date := GetDatetime()
	hostname := GetHostname()
	messageStr := ""
	tablename := os.Getenv("TABLE_NAME")

	// パラメータを作成する
	putParams := &dynamodb.PutItemInput{
        TableName: aws.String(tablename),
        Item: map[string]*dynamodb.AttributeValue{
            "MessageID": {
                S : aws.String(messageId),
            },
            "Date": {
                S : aws.String(date),
			},
			"Hostname": {
                S : aws.String(hostname),
			},
        },
    }

	// Messagesテーブルに格納する
    _, putErr := ddb.PutItem(putParams)
    if putErr != nil {
		messageStr = "Not connected DynamoDB"
	} else {
		messageStr = "Connected DynamoDB"
	}

	res := Response{
		Date:		date,
		Hostname:	hostname,
		MessageStr:	messageStr,
	}

    bytes, _ := json.Marshal(res)
    outputString := string(bytes)

	Header := map[string]string {
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
		"Content-Type": "application/json",
	}

	return events.APIGatewayProxyResponse{
		Headers: 	Header,
		Body:       outputString,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}