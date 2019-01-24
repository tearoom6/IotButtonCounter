package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/tearoom6/IotButtonCounter/aws"
	"github.com/tearoom6/IotButtonCounter/slack"
)

func processCounter(tableName string, clickType string, deviceId string) int {
	dynamodb, _ := aws.InitDynamoDbClient()
	currentCount, _ := dynamodb.GetNumberItem(tableName, deviceId)

	switch clickType {
	case "SINGLE":
		// Return incremented value.
		newCount := currentCount + 1
		dynamodb.PutNumberItem(tableName, deviceId, newCount)
		return newCount
	case "DOUBLE":
		// Not increment, just return current value.
		return currentCount
	case "LONG":
		// Reset counter.
		newCount := 0
		dynamodb.PutNumberItem(tableName, deviceId, newCount)
		return newCount
	}
	return -1
}

func sendSlackMessage(webhookUrl string, channel string, text string) {
	webhook := slack.Webhook{
		Url: webhookUrl,
	}
	message := slack.Message{
		Text:    text,
		Channel: channel,
	}
	webhook.Send(message)
}

func handleRequest(ctx context.Context, event aws.IotEnterpriseButtonEvent) {
	log.Print("Lambda received new request.")

	lctx, _ := lambdacontext.FromContext(ctx)
	log.Print(lctx)
	log.Print(event)

	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	slackChannel := os.Getenv("SLACK_CHANNEL")
	tableName := os.Getenv("TABLE_NAME")
	currentCount := processCounter(tableName, event.DeviceEvent.ButtonClicked.ClickType, event.DeviceInfo.DeviceId)

	sendSlackMessage(slackWebhookUrl, slackChannel, strconv.Itoa(currentCount))
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda.
	lambda.Start(handleRequest)
}
