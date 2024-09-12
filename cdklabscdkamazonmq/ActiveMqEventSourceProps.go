package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Experimental.
type ActiveMqEventSourceProps struct {
	// A secret with credentials of the user to use when receiving messages.
	//
	// The credentials in the secret have fields required:
	//  * username
	// * password.
	// Experimental.
	Credentials awssecretsmanager.ISecret `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the queue that the function will receive messages from.
	// Experimental.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// If the default permissions should be added to the Lambda function's execution role.
	// Default: true.
	//
	// Experimental.
	AddPermissions *bool `field:"optional" json:"addPermissions" yaml:"addPermissions"`
	// source at the time of invoking your function.
	//
	// Your function receives an
	// The largest number of records that AWS Lambda will retrieve from your event
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of: 10000.
	// Default: 100.
	//
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5).
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-batching
	//
	// Default: - Duration.millis(500) for Amazon MQ.
	//
	// Experimental.
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The ActiveMQ broker deployment to receive messages from.
	// Experimental.
	Broker IActiveMqBrokerDeployment `field:"required" json:"broker" yaml:"broker"`
}

