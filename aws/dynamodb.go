package aws

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	idKey    = "id"
	valueKey = "value"
)

type client struct {
	instance *dynamodb.DynamoDB
}

func InitDynamoDbClient() (*client, error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	instance := dynamodb.New(sess)
	return &client{instance: instance}, nil
}

func (cli *client) GetNumberItem(tableName string, id string) (int, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			idKey: {
				S: aws.String(id),
			},
		},
	}
	item, err := cli.instance.GetItem(params)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	if item == nil || item.Item == nil {
		return 0, nil
	}
	value, parseErr := strconv.Atoi(*item.Item[valueKey].N)
	if parseErr != nil {
		log.Print(parseErr)
		return 0, parseErr
	}
	return value, nil
}

func (cli *client) PutNumberItem(tableName string, id string, value int) error {
	params := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			idKey: {
				S: aws.String(id),
			},
			valueKey: {
				N: aws.String(strconv.Itoa(value)),
			},
		},
	}
	_, err := cli.instance.PutItem(params)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
