package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Represents an AWS Lambda Event Source Mapping for RabbitMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
type RabbitMqEventSource interface {
	EventSourceBase
	awslambda.IEventSource
	// Experimental.
	MqType() *string
	// properties of the RabbitMQ event source.
	// Experimental.
	Props() *EventSourceBaseProps
	// Experimental.
	AddToSourceAccessConfigurations(config *awslambda.SourceAccessConfiguration)
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	// Experimental.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for RabbitMqEventSource
type jsiiProxy_RabbitMqEventSource struct {
	jsiiProxy_EventSourceBase
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_RabbitMqEventSource) MqType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mqType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqEventSource) Props() *EventSourceBaseProps {
	var returns *EventSourceBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Instantiates an AWS Lambda Event Source Mapping for RabbitMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
func NewRabbitMqEventSource(props *RabbitMqEventSourceProps) RabbitMqEventSource {
	_init_.Initialize()

	if err := validateNewRabbitMqEventSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitMqEventSource{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Instantiates an AWS Lambda Event Source Mapping for RabbitMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
func NewRabbitMqEventSource_Override(r RabbitMqEventSource, props *RabbitMqEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqEventSource",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RabbitMqEventSource) AddToSourceAccessConfigurations(config *awslambda.SourceAccessConfiguration) {
	if err := r.validateAddToSourceAccessConfigurationsParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addToSourceAccessConfigurations",
		[]interface{}{config},
	)
}

func (r *jsiiProxy_RabbitMqEventSource) Bind(target awslambda.IFunction) {
	if err := r.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"bind",
		[]interface{}{target},
	)
}

