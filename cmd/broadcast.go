/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"math/rand/v2"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/neilmcgibbon/sns-example/internal/helpers/outputhelper"
	"github.com/spf13/cobra"
)

var broadcastCmdFlag_TopicARN string

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		awscfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion("eu-west-1"))
		if err != nil {
			outputhelper.Fail(err.Error(), true)
		}

		client := sns.NewFromConfig(awscfg)

		wait := make(chan struct{})
		for {
			go func(ch chan struct{}) {
				i := rand.IntN(1000)
				msg := fmt.Sprintf("emitting new event (number %d)", i)
				outputhelper.Success(msg, false)

				res, err := client.Publish(context.Background(), &sns.PublishInput{
					Message:        aws.String(fmt.Sprintf("%d", i)),
					TopicArn:       aws.String(broadcastCmdFlag_TopicARN),
					MessageGroupId: aws.String("1"),
				})
				if err != nil {
					outputhelper.Fail(err.Error(), true)
				}
				outputhelper.Success(*res.MessageId, false)
				time.Sleep(10 * time.Second)
				ch <- struct{}{}
			}(wait)
			<-wait
		}

	},
}

func init() {

	broadcastCmd.PersistentFlags().StringVar(&broadcastCmdFlag_TopicARN, "topicarn", "", "The SNS Topic ARN for broadcasting")

	rootCmd.AddCommand(broadcastCmd)
}
