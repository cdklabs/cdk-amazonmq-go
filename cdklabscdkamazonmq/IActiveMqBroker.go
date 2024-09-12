package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IActiveMqBroker interface {
	// A set of endpoints for the broker.
	// Experimental.
	Endpoints() *ActiveMqBrokerEndpoints
	// The IP address of the broker.
	// Experimental.
	IpAddress() *string
}

// The jsii proxy for IActiveMqBroker
type jsiiProxy_IActiveMqBroker struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveMqBroker) Endpoints() *ActiveMqBrokerEndpoints {
	var returns *ActiveMqBrokerEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBroker) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

