//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BrokerConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BrokerConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BrokerConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBrokerConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBrokerConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBrokerConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBrokerConfigurationParameters(scope constructs.Construct, id *string, props *ConfigurationProps) error {
	return nil
}

