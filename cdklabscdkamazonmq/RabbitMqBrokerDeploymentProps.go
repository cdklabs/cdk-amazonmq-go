package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type RabbitMqBrokerDeploymentProps struct {
	// Determines whether the broker will undergo a minor version upgrade during the maintenance window.
	// Default: - false. No minor version upgrade happens.
	//
	// Experimental.
	AutoMinorVersionUpgrade *bool `field:"required" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// An instance type to use for the broker.
	//
	// Only a subset of available instance types is allowed.
	// See: https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-instance-types.html
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Specifies whether the broker is open to public Internet or deployed with endpoints in own VPC.
	// Experimental.
	PubliclyAccessible *bool `field:"required" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Experimental.
	BrokerName *string `field:"optional" json:"brokerName" yaml:"brokerName"`
	// Sets the retention days for the broker's CloudWatch LogGroups.
	// Default: - undefined; CloudWatch Log Groups retention is set to never expire.
	//
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	MaintenanceWindowStartTime *MaintenanceWindowStartTime `field:"optional" json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// The Security Groups to apply for a non publicly accessible broker.
	//
	// NOTE: This needs to be set only if `publiclyAccessible` is true.
	// Default: - undefined. If no VPC is selected then a default VPC's default SG will be used.
	//              Otherwise - a security group will be created.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC in which create the communication endpoints for a private broker.
	// Default: - undefined. A default VPC will be used
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// vpcSubnets and vpc are optional.
	//
	// But when present - publiclyAccessible attribute must equal false.
	// Default: - undefined. If vpc is present - this attribute must be present.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Sets the credentials of the broker administrative user.
	// Experimental.
	Admin *Admin `field:"required" json:"admin" yaml:"admin"`
	// Sets the version of the broker engine.
	// Experimental.
	Version RabbitMqBrokerEngineVersion `field:"required" json:"version" yaml:"version"`
	// Sets the CloudWatch logs exports for the broker.
	// Experimental.
	CloudwatchLogsExports *RabbitMqCloudwatchLogsExports `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// Sets the configuration of the broker.
	// Experimental.
	Configuration IRabbitMqBrokerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
}

