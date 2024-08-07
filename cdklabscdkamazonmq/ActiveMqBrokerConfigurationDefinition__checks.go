//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"
)

func validateActiveMqBrokerConfigurationDefinition_DataParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateNewActiveMqBrokerConfigurationDefinitionParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

