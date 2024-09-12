package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type ActiveMqUser struct {
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Experimental.
	HasConsoleAccess *bool `field:"optional" json:"hasConsoleAccess" yaml:"hasConsoleAccess"`
}

