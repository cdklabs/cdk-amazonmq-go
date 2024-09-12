package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"
)

// Experimental.
type RabbitMqBrokerConfigurationDefinition interface {
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RabbitMqBrokerConfigurationDefinition
type jsiiProxy_RabbitMqBrokerConfigurationDefinition struct {
	_ byte // padding
}

// Experimental.
func NewRabbitMqBrokerConfigurationDefinition(data *string) RabbitMqBrokerConfigurationDefinition {
	_init_.Initialize()

	if err := validateNewRabbitMqBrokerConfigurationDefinitionParameters(data); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitMqBrokerConfigurationDefinition{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationDefinition",
		[]interface{}{data},
		&j,
	)

	return &j
}

// Experimental.
func NewRabbitMqBrokerConfigurationDefinition_Override(r RabbitMqBrokerConfigurationDefinition, data *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationDefinition",
		[]interface{}{data},
		r,
	)
}

// Experimental.
func RabbitMqBrokerConfigurationDefinition_Data(data *string) RabbitMqBrokerConfigurationDefinition {
	_init_.Initialize()

	if err := validateRabbitMqBrokerConfigurationDefinition_DataParameters(data); err != nil {
		panic(err)
	}
	var returns RabbitMqBrokerConfigurationDefinition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationDefinition",
		"data",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Experimental.
func RabbitMqBrokerConfigurationDefinition_Parameters(parameters *RabbitMqBrokerConfigurationParameters) RabbitMqBrokerConfigurationDefinition {
	_init_.Initialize()

	if err := validateRabbitMqBrokerConfigurationDefinition_ParametersParameters(parameters); err != nil {
		panic(err)
	}
	var returns RabbitMqBrokerConfigurationDefinition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationDefinition",
		"parameters",
		[]interface{}{parameters},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerConfigurationDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

