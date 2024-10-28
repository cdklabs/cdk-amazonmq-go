//go:build !no_runtime_type_checking

package cdklabscdkamazonmq

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (e *jsiiProxy_EventSourceBase) validateAddToSourceAccessConfigurationsParameters(config *awslambda.SourceAccessConfiguration) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_EventSourceBase) validateBindParameters(target awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func validateNewEventSourceBaseParameters(props *EventSourceBaseProps, mqType *string) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if mqType == nil {
		return fmt.Errorf("parameter mqType is required, but nil was provided")
	}

	return nil
}

