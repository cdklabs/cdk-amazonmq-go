package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type ActiveMqBrokerDeploymentProps struct {
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
	// Determines whether the broker will undergo a patch version upgrade during the maintenance window.
	//
	// NOTE: Contrary to the name this  setting does not upgrade the minor versions, but patch versions (i.e. in the X.Y.Z notation - only the Z numbers are upgraded)
	// Default: - for versions with the patch version number the default is not to upgrade the patch versions; for versions withouth the patch version number patch versions are updated and this setting takes no effect.
	//
	// Experimental.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
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
	// Sets the User Management option for the Amazon MQ for ActiveMQ broker.
	// Experimental.
	UserManagement IActiveMqBrokerUserManagement `field:"required" json:"userManagement" yaml:"userManagement"`
	// Sets the version of the Amazon MQ for ActiveMQ broker engine.
	// Experimental.
	Version ActiveMqBrokerEngineVersion `field:"required" json:"version" yaml:"version"`
	// Sets the CloudWatch Logs exports for the Amazon MQ for ActiveMQ broker.
	// Default: - undefined; No logs are exported to CloudWatch.
	//
	// Experimental.
	CloudwatchLogsExports *ActiveMqCloudwatchLogsExports `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// Sets the configuration of the Amazon MQ for ActiveMQ broker.
	// Experimental.
	Configuration IActiveMqBrokerConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Sets the number of days to retain logs for the Amazon MQ for ActiveMQ broker.
	// Experimental.
	LogsRetentionDays *float64 `field:"optional" json:"logsRetentionDays" yaml:"logsRetentionDays"`
}

