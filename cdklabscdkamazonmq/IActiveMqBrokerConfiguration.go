package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IActiveMqBrokerConfiguration interface {
	IBrokerConfiguration
	// Experimental.
	AssociateWith(broker IActiveMqBrokerDeployment) ConfigurationAssociation
	// Experimental.
	CreateRevision(options *ActiveMqBrokerConfigurationOptions) IActiveMqBrokerConfiguration
}

// The jsii proxy for IActiveMqBrokerConfiguration
type jsiiProxy_IActiveMqBrokerConfiguration struct {
	jsiiProxy_IBrokerConfiguration
}

func (i *jsiiProxy_IActiveMqBrokerConfiguration) AssociateWith(broker IActiveMqBrokerDeployment) ConfigurationAssociation {
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

func (i *jsiiProxy_IActiveMqBrokerConfiguration) CreateRevision(options *ActiveMqBrokerConfigurationOptions) IActiveMqBrokerConfiguration {
	if err := i.validateCreateRevisionParameters(options); err != nil {
		panic(err)
	}
	var returns IActiveMqBrokerConfiguration

	_jsii_.Invoke(
		i,
		"createRevision",
		[]interface{}{options},
		&returns,
	)

	return returns
}

