//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func validateRabbitMqCustomResourcePolicy_FromStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	return nil
}

