//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBrokerDeployment) validateMetricParameters(metricName *string, options *awscloudwatch.MetricOptions) error {
	return nil
}

