package bridge

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_ConvertLambdaDynamoDBAttributeValue(t *testing.T) {
	type Entry struct {
		ID   uint64 `dynamodbav:"id"`
		Name string `dynamodbav:"name"`
	}

	lambdaAVMap := map[string]events.DynamoDBAttributeValue{
		"id":   events.NewNumberAttribute("65535"),
		"name": events.NewStringAttribute("John Doe"),
	}

	converted, err := ConvertLambdaDynamoDBAttributeValue[Entry](lambdaAVMap)
	assert.NoError(t, err)
	assert.EqualValues(t, &Entry{
		ID:   65535,
		Name: "John Doe",
	}, converted)
}
