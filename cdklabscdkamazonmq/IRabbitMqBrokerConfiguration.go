package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRabbitMqBrokerConfiguration interface {
	IBrokerConfiguration
	// Experimental.
	AssociateWith(broker IRabbitMqBrokerDeployment) ConfigurationAssociation
	// Experimental.
	CreateRevision(options *RabbitMqBrokerConfigurationOptions) IRabbitMqBrokerConfiguration
}

// The jsii proxy for IRabbitMqBrokerConfiguration
type jsiiProxy_IRabbitMqBrokerConfiguration struct {
	jsiiProxy_IBrokerConfiguration
}

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) AssociateWith(broker IRabbitMqBrokerDeployment) ConfigurationAssociation {
	if err := i.validateAssociateWithParameters(broker); err != nil {
		panic(err)
	}
	var returns ConfigurationAssociation

	_jsii_.Invoke(
		i,
		"associateWith",
		[]interface{}{broker},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) CreateRevision(options *RabbitMqBrokerConfigurationOptions) IRabbitMqBrokerConfiguration {
	if err := i.validateCreateRevisionParameters(options); err != nil {
		panic(err)
	}
	var returns IRabbitMqBrokerConfiguration

	_jsii_.Invoke(
		i,
		"createRevision",
		[]interface{}{options},
		&returns,
	)

	return returns
}

