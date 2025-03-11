package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type IRabbitMqBrokerInstance interface {
	IRabbitMqBroker
	IRabbitMqBrokerDeployment
}

// The jsii proxy for IRabbitMqBrokerInstance
type jsiiProxy_IRabbitMqBrokerInstance struct {
	jsiiProxy_IRabbitMqBroker
	jsiiProxy_IRabbitMqBrokerDeployment
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricAckRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAckRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAckRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricChannelCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricChannelCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricChannelCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricConfirmRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConfirmRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConfirmRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConsumerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConsumerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricExchangeCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricExchangeCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricExchangeCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMessageCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMessageCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricMessageReadyCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMessageReadyCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMessageReadyCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricMessageUnacknowledgedCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMessageUnacknowledgedCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMessageUnacknowledgedCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricPublishRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPublishRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPublishRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricQueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricQueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricQueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQDiskFree(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQDiskFreeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQDiskFree",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQDiskFreeLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQDiskFreeLimitParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQDiskFreeLimit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQFdUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQFdUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQFdUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQIOReadAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQIOReadAverageTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQIOReadAverageTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQIOWriteAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQIOWriteAverageTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQIOWriteAverageTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQMemLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQMemLimitParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQMemLimit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricRabbitMQMemUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricRabbitMQMemUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricRabbitMQMemUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRabbitMqBrokerInstance) MetricSystemCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSystemCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSystemCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Endpoints() *RabbitMqBrokerEndpoints {
	var returns *RabbitMqBrokerEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

