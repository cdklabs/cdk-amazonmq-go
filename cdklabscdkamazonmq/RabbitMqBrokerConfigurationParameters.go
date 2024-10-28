package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type RabbitMqBrokerConfigurationParameters struct {
	// Experimental.
	ConsumerTimeout awscdk.Duration `field:"required" json:"consumerTimeout" yaml:"consumerTimeout"`
}

