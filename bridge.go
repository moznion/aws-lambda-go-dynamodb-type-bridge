package bridge

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ConvertLambdaDynamoDBAttributeValue[T any](attributeValue map[string]events.DynamoDBAttributeValue) (*T, error) {
	avMap := make(map[string]*dynamodb.AttributeValue)
	for key, val := range attributeValue {
		bs, err := val.MarshalJSON()
		if err != nil {
			return nil, err
		}

		var av dynamodb.AttributeValue
		err = json.Unmarshal(bs, &av)
		if err != nil {
			return nil, err
		}
		avMap[key] = &av
	}

	var target T
	err := dynamodbattribute.UnmarshalMap(avMap, &target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
