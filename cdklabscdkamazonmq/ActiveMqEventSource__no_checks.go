//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActiveMqEventSource) validateAddToSourceAccessConfigurationsParameters(config *awslambda.SourceAccessConfiguration) error {
	return nil
}

func (a *jsiiProxy_ActiveMqEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewActiveMqEventSourceParameters(props *ActiveMqEventSourceProps) error {
	return nil
}

