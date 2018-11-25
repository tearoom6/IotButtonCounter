# IotButtonCounter

## Features

Use AWS IoT Enterprise Button as counter which send current value to Slack.

- Single click -> count up
- Double click -> check current value
- Long press   -> reset counter


## Setup

First, claim IoT Button device by following [Claiming Devices - AWS IoT 1-Click](https://docs.aws.amazon.com/iot-1-click/latest/developerguide/1click-claiming.html).

Then, setup resouces using CloudFormation:

```sh
S3_BUCKET="<BUCKET>"
FUNCTION_NAME="IotButtonCounter"
STACK_NAME="IotButtonCounter"
SLACK_WEBHOOK_URL="https://hooks.slack.com/services/<YOUR_SLACK_SPECIFIC_TOKENS>"
DSN="<IOT_BUTTON_DSN>"

# Build Go source code.
GOOS=linux GOARCH=amd64 go build -o bin/main main.go

# Archive binary to zip file.
zip build/handler.zip bin/main

# Package Lambda binary and upload it to S3 bucket.
aws cloudformation package \
  --template-file cloudformation-template.yml \
  --s3-bucket $S3_BUCKET \
  --s3-prefix "Lambda/$FUNCTION_NAME" \
  --output-template-file packaged-template.yml

# Deploy all resources in AWS. (Also create CloudFormation stack)
aws cloudformation deploy \
  --template-file packaged-template.yml \
  --stack-name $STACK_NAME \
  --capabilities CAPABILITY_IAM \
  --parameter-overrides "SlackWebhookUrl=$SLACK_WEBHOOK_URL" "DeviceDsn=$DSN"

# Destroy all resources in AWS. (Also delete CloudFormation stack)
aws cloudformation delete-stack --stack-name $STACK_NAME
```

