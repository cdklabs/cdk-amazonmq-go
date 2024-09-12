//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventSourceBase) validateAddToSourceAccessConfigurationsParameters(config *awslambda.SourceAccessConfiguration) error {
	return nil
}

func (e *jsiiProxy_EventSourceBase) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewEventSourceBaseParameters(props *EventSourceBaseProps, mqType *string) error {
	return nil
}

