//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigurationAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConfigurationAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConfigurationAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateConfigurationAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigurationAssociation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConfigurationAssociation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationAssociationParameters(scope constructs.Construct, id *string, props *ConfigurationAssociationProps) error {
	return nil
}

