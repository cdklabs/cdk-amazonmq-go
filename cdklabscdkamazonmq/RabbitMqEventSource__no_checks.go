//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RabbitMqEventSource) validateAddToSourceAccessConfigurationsParameters(config *awslambda.SourceAccessConfiguration) error {
	return nil
}

func (r *jsiiProxy_RabbitMqEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewRabbitMqEventSourceParameters(props *RabbitMqEventSourceProps) error {
	return nil
}

