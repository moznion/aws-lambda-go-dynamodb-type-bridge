# aws-lambda-go-dynamodb-type-bridge [![.github/workflows/check.yml](https://github.com/moznion/aws-lambda-go-dynamodb-type-bridge/actions/workflows/check.yml/badge.svg)](https://github.com/moznion/aws-lambda-go-dynamodb-type-bridge/actions/workflows/check.yml)

A library that makes it easier to unmarshal the `map[string]events.DynamoDBAttributeValue` which comes from `aws/aws-lambda-go/events` into a target struct that uses `dynamodbav` custom tags.

This library uses generics, so it requires Go runtimve version 1.18 or later.

## Motivation

There is no official one-stop solution to convert the `events.DynamoDBAttributeValue` value which comes from `aws/aws-lambda-go` to the attribute value of aws-sdk-go and/or the struct which uses `dynamodbav` custom tags with `aws/aws-sdk-go/service/dynamodb/dynamodbattribute` so far.

Accordingly, this library aims to provide such feature to make it easier with Go generics feature.

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

