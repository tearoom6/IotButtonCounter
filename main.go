package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/tearoom6/IotButtonCounter/slack"
)

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

func HandleRequest(ctx context.Context) {
	log.Print("Lambda received new request.")

	lctx, _ := lambdacontext.FromContext(ctx)
	log.Print(lctx)

	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	slackChannel := os.Getenv("SLACK_CHANNEL")

	sendSlackMessage(slackWebhookUrl, slackChannel, "test")
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda.
	lambda.Start(HandleRequest)
}
