//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"
)

func validateActiveMqBrokerEngineVersion_OfParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

func validateNewActiveMqBrokerEngineVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

