//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func validateRabbitMqCustomResourcePolicy_FromStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	if statements == nil {
		return fmt.Errorf("parameter statements is required, but nil was provided")
	}

	return nil
}

