package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IActiveMqBrokerRedundantPair interface {
	IActiveMqBrokerDeployment
	// Experimental.
	First() IActiveMqBroker
	// Experimental.
	Second() IActiveMqBroker
}

// The jsii proxy for IActiveMqBrokerRedundantPair
type jsiiProxy_IActiveMqBrokerRedundantPair struct {
	jsiiProxy_IActiveMqBrokerDeployment
}

func (j *jsiiProxy_IActiveMqBrokerRedundantPair) First() IActiveMqBroker {
	var returns IActiveMqBroker
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerRedundantPair) Second() IActiveMqBroker {
	var returns IActiveMqBroker
	_jsii_.Get(
		j,
		"second",
		&returns,
	)
	return returns
}

