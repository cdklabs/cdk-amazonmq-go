//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IActiveMqBrokerConfiguration) validateAssociateWithParameters(broker IActiveMqBrokerDeployment) error {
	return nil
}

func (i *jsiiProxy_IActiveMqBrokerConfiguration) validateCreateRevisionParameters(options *ActiveMqBrokerConfigurationOptions) error {
	return nil
}

