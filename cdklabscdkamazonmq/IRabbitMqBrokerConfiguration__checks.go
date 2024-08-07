//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) validateAssociateWithParameters(broker IRabbitMqBrokerDeployment) error {
	if broker == nil {
		return fmt.Errorf("parameter broker is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IRabbitMqBrokerConfiguration) validateCreateRevisionParameters(options *RabbitMqBrokerConfigurationOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

