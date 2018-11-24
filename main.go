package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/tearoom6/IotButtonCounter/aws"
	"github.com/tearoom6/IotButtonCounter/slack"
)

func getClickType(clickType string, deviceId string) string {
	switch clickType {
	case "SINGLE":
		return "single"
	case "DOUBLE":
		return "double"
	case "LONG":
		return "long"
	}
	return "none"
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
	clickType := getClickType(event.DeviceEvent.ButtonClicked.ClickType, event.DeviceInfo.DeviceId)

	sendSlackMessage(slackWebhookUrl, slackChannel, clickType)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda.
	lambda.Start(handleRequest)
}
