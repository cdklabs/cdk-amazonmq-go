//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RabbitMqCustomResource) validateGetResponseFieldParameters(key *string) error {
	return nil
}

func (r *jsiiProxy_RabbitMqCustomResource) validateGetResponseFieldReferenceParameters(key *string) error {
	return nil
}

func validateRabbitMqCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRabbitMqCustomResourceParameters(scope constructs.Construct, id *string, props *RabbitMqCustomResourceProps) error {
	return nil
}

