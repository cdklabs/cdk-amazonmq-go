package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type Admin struct {
	// Sets the administrative user password.
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Sets the administrative user name.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

