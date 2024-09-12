package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Represents an AWS Lambda Event Source Mapping for ActiveMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
type ActiveMqEventSource interface {
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

// The jsii proxy struct for ActiveMqEventSource
type jsiiProxy_ActiveMqEventSource struct {
	jsiiProxy_EventSourceBase
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_ActiveMqEventSource) MqType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mqType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqEventSource) Props() *EventSourceBaseProps {
	var returns *EventSourceBaseProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Instantiates an AWS Lambda Event Source Mapping for ActiveMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
func NewActiveMqEventSource(props *ActiveMqEventSourceProps) ActiveMqEventSource {
	_init_.Initialize()

	if err := validateNewActiveMqEventSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveMqEventSource{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Instantiates an AWS Lambda Event Source Mapping for ActiveMQ.
//
// This event source will add additional permissions to
// the AWS Lambda function's IAM Role following https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions
// Experimental.
func NewActiveMqEventSource_Override(a ActiveMqEventSource, props *ActiveMqEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqEventSource",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_ActiveMqEventSource) AddToSourceAccessConfigurations(config *awslambda.SourceAccessConfiguration) {
	if err := a.validateAddToSourceAccessConfigurationsParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToSourceAccessConfigurations",
		[]interface{}{config},
	)
}

func (a *jsiiProxy_ActiveMqEventSource) Bind(target awslambda.IFunction) {
	if err := a.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{target},
	)
}

