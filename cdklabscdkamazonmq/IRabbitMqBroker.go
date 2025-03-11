package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Experimental.
type IRabbitMqBroker interface {
	awscdk.IResource
	// Experimental.
	Endpoints() *RabbitMqBrokerEndpoints
}

// The jsii proxy for IRabbitMqBroker
type jsiiProxy_IRabbitMqBroker struct {
	internal.Type__awscdkIResource
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

