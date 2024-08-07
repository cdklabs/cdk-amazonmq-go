package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRabbitMqBroker interface {
	// Experimental.
	Endpoints() *RabbitMqBrokerEndpoints
}

// The jsii proxy for IRabbitMqBroker
type jsiiProxy_IRabbitMqBroker struct {
	_ byte // padding
}

func (j *jsiiProxy_IRabbitMqBroker) Endpoints() *RabbitMqBrokerEndpoints {
	var returns *RabbitMqBrokerEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

