//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"
)

func validateRabbitMqBrokerEngineVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateNewRabbitMqBrokerEngineVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

