package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq"
)

// Experimental.
type ActiveMqBrokerDeploymentUserManagementDefinition struct {
	// Experimental.
	Users *[]*awsamazonmq.CfnBroker_UserProperty `field:"required" json:"users" yaml:"users"`
	// Experimental.
	AuthenticationStrategy ActiveMqAuthenticationStrategy `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Experimental.
	LdapServerMetadata *awsamazonmq.CfnBroker_LdapServerMetadataProperty `field:"optional" json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
}

