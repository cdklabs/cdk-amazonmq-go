package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for RabbitMqCustomResource.
//
// Note that at least onCreate, onUpdate or onDelete must be specified.
// Experimental.
type RabbitMqCustomResourceProps struct {
	// The broker to send requests to.
	// Experimental.
	Broker IRabbitMqBroker `field:"required" json:"broker" yaml:"broker"`
	// The secret containing the broker login credentials.
	// Experimental.
	Credentials awssecretsmanager.ISecret `field:"required" json:"credentials" yaml:"credentials"`
	// The logGroup to use for the function.
	// Experimental.
	LogGroup awslogs.LogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// LogGroup retention to use for the function.
	// Default: RetentionDays.INFINITE
	//
	// Deprecated: use logGroup instead.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The RabbitMQ Management HTTP API call to make when the resource is created.
	// Default: - the call when the resource is updated.
	//
	// Experimental.
	OnCreate *RabbitMqApiCall `field:"optional" json:"onCreate" yaml:"onCreate"`
	// The RabbitMQ Management HTTP API call to make when the resource is updated.
	// Default: - no call.
	//
	// Experimental.
	OnDelete *RabbitMqApiCall `field:"optional" json:"onDelete" yaml:"onDelete"`
	// The RabbitMQ Management HTTP API call to make when the resource is updated.
	// Default: - no call.
	//
	// Experimental.
	OnUpdate *RabbitMqApiCall `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// The policies to attach to the function's role.
	// Experimental.
	Policy RabbitMqCustomResourcePolicy `field:"optional" json:"policy" yaml:"policy"`
	// The execution role for the function.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The security groups to assign to the function.
	// Experimental.
	SecurityGroups *[]awsec2.SecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The timeout for the custom resource.
	// Default: Duration.minutes(1)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The vpc to connect to.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The vpc subnets to connect to.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

