//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateRabbitMqBrokerConfigurationDefinition_DataParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateRabbitMqBrokerConfigurationDefinition_ParametersParameters(parameters *RabbitMqBrokerConfigurationParameters) error {
	if parameters == nil {
		return fmt.Errorf("parameter parameters is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}

func validateNewRabbitMqBrokerConfigurationDefinitionParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

