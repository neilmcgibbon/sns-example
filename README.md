# Example

### AWS SETUP

> Do the the following in eu-west-1 region.

Create AWS SNS Topic with following settings:

- Type: FIFO queue
- Content-based message deduplication: enabled

Once created, copy the topic ARN - this will be the "topicarn" used for broadcast

Create two SQS queus with the following settings

- Type: FIFO
- Receive Message Wait time: 20 seconds (IMPORTANT)
- Encryption disabled

Once created, copy the Queue URLs - these will be used for "queueurl" in subscriptions. Then subscribe each queue to the SNS topic, by clicking into the created queue, and pressing the "Subscribe to Amaazon SNS Topic" button at the bottom. Select the topic you created above from the available list, and press "Save". Make sure to do this from both queues.

### Running the app

Make sure your terminal environment is set to the same AWS environment as your resource setup in the previous step. The app uses the current environment.

Build the app:

```
go build
```

Open 3 terminal windows. In the first one, run the broadcast:

```
# This will send a random number to the topic every 10 seconds
./sns-example broadcast --topicarn <your topicarn you copied earlier>
```

In another terminal window, subscribe to your first queue - this represents a microsercvice consuming messages off its queue:

```
./sns-example subscribe --queueurl <your queueurl you copied earlier, for first queue>
```

In last terminal window, subscribe to your second queue - this represents a microsercvice consuming messages off its queue:

```
./sns-example subscribe --queueurl <your queueurl you copied earlier, for second queue>
```
