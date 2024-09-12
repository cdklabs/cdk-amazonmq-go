package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"
)

// Experimental.
type ActiveMqBrokerConfigurationDefinition interface {
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveMqBrokerConfigurationDefinition
type jsiiProxy_ActiveMqBrokerConfigurationDefinition struct {
	_ byte // padding
}

// Experimental.
func NewActiveMqBrokerConfigurationDefinition(data *string) ActiveMqBrokerConfigurationDefinition {
	_init_.Initialize()

	if err := validateNewActiveMqBrokerConfigurationDefinitionParameters(data); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveMqBrokerConfigurationDefinition{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationDefinition",
		[]interface{}{data},
		&j,
	)

	return &j
}

// Experimental.
func NewActiveMqBrokerConfigurationDefinition_Override(a ActiveMqBrokerConfigurationDefinition, data *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationDefinition",
		[]interface{}{data},
		a,
	)
}

// Experimental.
func ActiveMqBrokerConfigurationDefinition_Data(data *string) ActiveMqBrokerConfigurationDefinition {
	_init_.Initialize()

	if err := validateActiveMqBrokerConfigurationDefinition_DataParameters(data); err != nil {
		panic(err)
	}
	var returns ActiveMqBrokerConfigurationDefinition

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationDefinition",
		"data",
		[]interface{}{data},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerConfigurationDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

