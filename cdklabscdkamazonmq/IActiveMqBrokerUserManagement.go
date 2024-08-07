package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IActiveMqBrokerUserManagement interface {
	// Experimental.
	Render() *ActiveMqBrokerDeploymentUserManagementDefinition
}

// The jsii proxy for IActiveMqBrokerUserManagement
type jsiiProxy_IActiveMqBrokerUserManagement struct {
	_ byte // padding
}

func (i *jsiiProxy_IActiveMqBrokerUserManagement) Render() *ActiveMqBrokerDeploymentUserManagementDefinition {
	var returns *ActiveMqBrokerDeploymentUserManagementDefinition

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

