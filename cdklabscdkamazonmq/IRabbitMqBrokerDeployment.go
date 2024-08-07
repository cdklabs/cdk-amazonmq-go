package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Experimental.
type IRabbitMqBrokerDeployment interface {
	IBrokerDeployment
	awscdk.IResource
	// Experimental.
	MetricAckRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricChannelCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricConfirmRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricExchangeCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricMessageReadyCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricMessageUnacknowledgedCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricPublishRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricQueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQDiskFree(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQDiskFreeLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQFdUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQIOReadAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQIOWriteAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQMemLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricRabbitMQMemUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricSystemCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy for IRabbitMqBrokerDeployment
type jsiiProxy_IRabbitMqBrokerDeployment struct {
	jsiiProxy_IBrokerDeployment
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricAckRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricChannelCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricConfirmRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricExchangeCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricMessageReadyCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricMessageUnacknowledgedCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricPublishRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricQueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQDiskFree(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQDiskFreeLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQFdUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQIOReadAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQIOWriteAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQMemLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricRabbitMQMemUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) MetricSystemCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRabbitMqBrokerDeployment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRabbitMqBrokerDeployment) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRabbitMqBrokerDeployment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

