package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"
)

// The Amazon RabbitMQ Broker Engine version.
// See: https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/rabbitmq-version-management.html
//
// Experimental.
type RabbitMqBrokerEngineVersion interface {
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RabbitMqBrokerEngineVersion
type jsiiProxy_RabbitMqBrokerEngineVersion struct {
	_ byte // padding
}

// Experimental.
func NewRabbitMqBrokerEngineVersion(version *string) RabbitMqBrokerEngineVersion {
	_init_.Initialize()

	if err := validateNewRabbitMqBrokerEngineVersionParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitMqBrokerEngineVersion{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Experimental.
func NewRabbitMqBrokerEngineVersion_Override(r RabbitMqBrokerEngineVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		[]interface{}{version},
		r,
	)
}

// Experimental.
func RabbitMqBrokerEngineVersion_Of(version *string) RabbitMqBrokerEngineVersion {
	_init_.Initialize()

	if err := validateRabbitMqBrokerEngineVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns RabbitMqBrokerEngineVersion

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func RabbitMqBrokerEngineVersion_V3_11_16() RabbitMqBrokerEngineVersion {
	_init_.Initialize()
	var returns RabbitMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"V3_11_16",
		&returns,
	)
	return returns
}

func RabbitMqBrokerEngineVersion_V3_11_20() RabbitMqBrokerEngineVersion {
	_init_.Initialize()
	var returns RabbitMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"V3_11_20",
		&returns,
	)
	return returns
}

func RabbitMqBrokerEngineVersion_V3_11_28() RabbitMqBrokerEngineVersion {
	_init_.Initialize()
	var returns RabbitMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"V3_11_28",
		&returns,
	)
	return returns
}

func RabbitMqBrokerEngineVersion_V3_12_13() RabbitMqBrokerEngineVersion {
	_init_.Initialize()
	var returns RabbitMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"V3_12_13",
		&returns,
	)
	return returns
}

func RabbitMqBrokerEngineVersion_V3_13() RabbitMqBrokerEngineVersion {
	_init_.Initialize()
	var returns RabbitMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		"V3_13",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RabbitMqBrokerEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

