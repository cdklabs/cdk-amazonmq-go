//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BrokerDeploymentBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BrokerDeploymentBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BrokerDeploymentBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BrokerDeploymentBase) validateMetricParameters(metricName *string, options *awscloudwatch.MetricOptions) error {
	return nil
}

func validateBrokerDeploymentBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBrokerDeploymentBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBrokerDeploymentBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBrokerDeploymentBaseParameters(scope constructs.Construct, id *string, props *BrokerDeploymentBaseProps) error {
	return nil
}

