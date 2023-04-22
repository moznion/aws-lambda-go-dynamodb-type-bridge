# aws-lambda-go-dynamodb-type-bridge [![.github/workflows/check.yml](https://github.com/moznion/aws-lambda-go-dynamodb-type-bridge/actions/workflows/check.yml/badge.svg)](https://github.com/moznion/aws-lambda-go-dynamodb-type-bridge/actions/workflows/check.yml)

A library that makes it easier to unmarshal the `map[string]events.DynamoDBAttributeValue` which comes from `aws/aws-lambda-go/events` into a target struct that uses `dynamodbav` custom tags.

This library uses generics, so it requires Go runtimve version 1.18 or later.

## Synopsis

```go
type Entry struct {
    ID   uint64 `dynamodbav:"id"`
    Name string `dynamodbav:"name"`
}

lambdaAVMap := map[string]events.DynamoDBAttributeValue{
    "id":   events.NewNumberAttribute("65535"),
    "name": events.NewStringAttribute("John Doe"),
}

converted, err := ConvertLambdaDynamoDBAttributeValue[Entry](lambdaAVMap)
fmt.Println(err) // => <nil>
fmt.Printf("%#v\n", converted) // => &bridge.Entry{ID:0xffff, Name:"John Doe"}
```

## Author

moznion (<moznion@mail.moznion.net>)

