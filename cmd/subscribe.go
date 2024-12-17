/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/neilmcgibbon/sns-example/internal/helpers/outputhelper"
	"github.com/spf13/cobra"
)

type gensqs struct {
	Message string `json:"Message"`
}

var subscribeCmdFlag_QeueURL string

// broadcastCmd represents the broadcast command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		awscfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion("eu-west-1"))
		if err != nil {
			outputhelper.Fail(err.Error(), true)
		}

		client := sqs.NewFromConfig(awscfg)

		for {
			spinner := outputhelper.SpinnerMessage("polling for new messages")
			messages, err := client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(subscribeCmdFlag_QeueURL),
				MaxNumberOfMessages: 1,
				WaitTimeSeconds:     20,
			})
			spinner.Stop()

			if err != nil {
				outputhelper.Fail(fmt.Sprintf("failed to fetch sqs message %v", err), true)
			}

			for _, message := range messages.Messages {
				var msg gensqs
				if err := json.Unmarshal([]byte(*message.Body), &msg); err != nil {
					outputhelper.Fail(fmt.Sprintf("failed to decode JSONe %v", err), true)
				}

				outputhelper.Success(fmt.Sprintf("received SQS message: %s", msg.Message), false)

				// ACK
				client.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(subscribeCmdFlag_QeueURL),
					ReceiptHandle: message.ReceiptHandle,
				})
			}

		}
	},
}

func init() {

	subscribeCmd.PersistentFlags().StringVar(&subscribeCmdFlag_QeueURL, "queueurl", "", "The SQS URL for receieving messages")

	rootCmd.AddCommand(subscribeCmd)
}
