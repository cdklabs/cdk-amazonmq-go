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
type EventSourceBase interface {
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

// The jsii proxy struct for EventSourceBase
type jsiiProxy_EventSourceBase struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_EventSourceBase) MqType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mqType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventSourceBase) Props() *EventSourceBaseProps {
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
func NewEventSourceBase_Override(e EventSourceBase, props *EventSourceBaseProps, mqType *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.EventSourceBase",
		[]interface{}{props, mqType},
		e,
	)
}

func (e *jsiiProxy_EventSourceBase) AddToSourceAccessConfigurations(config *awslambda.SourceAccessConfiguration) {
	if err := e.validateAddToSourceAccessConfigurationsParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addToSourceAccessConfigurations",
		[]interface{}{config},
	)
}

func (e *jsiiProxy_EventSourceBase) Bind(target awslambda.IFunction) {
	if err := e.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"bind",
		[]interface{}{target},
	)
}

