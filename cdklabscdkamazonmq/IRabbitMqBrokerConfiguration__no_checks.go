//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) validateAssociateWithParameters(broker IRabbitMqBrokerDeployment) error {
	return nil
}

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) validateCreateRevisionParameters(options *RabbitMqBrokerConfigurationOptions) error {
	return nil
}

