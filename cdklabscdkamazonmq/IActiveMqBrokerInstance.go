package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type IActiveMqBrokerInstance interface {
	IActiveMqBroker
	IActiveMqBrokerDeployment
}

// The jsii proxy for IActiveMqBrokerInstance
type jsiiProxy_IActiveMqBrokerInstance struct {
	jsiiProxy_IActiveMqBroker
	jsiiProxy_IActiveMqBrokerDeployment
}

func (i *jsiiProxy_IActiveMqBrokerInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IActiveMqBrokerInstance) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricAmqpMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAmqpMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAmqpMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricBurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBurstBalanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBurstBalance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricCpuCreditBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCpuCreditBalanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCpuCreditBalance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricCurrentConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCurrentConnectionsCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCurrentConnectionsCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDequeueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDequeueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricDispatchCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDispatchCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDispatchCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEnqueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEnqueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricEnqueueTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEnqueueTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEnqueueTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricEstablishedConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEstablishedConnectionsCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEstablishedConnectionsCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricExpiredCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricExpiredCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricExpiredCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricHeapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricHeapUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricHeapUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricInactiveDurableTopicSubscribersCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInactiveDurableTopicSubscribersCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInactiveDurableTopicSubscribersCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricInFlightCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInFlightCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInFlightCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricJobSchedulerStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricJobSchedulerStorePercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricJobSchedulerStorePercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricJournalFilesForFastRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricJournalFilesForFastRecoveryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricJournalFilesForFastRecovery",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricJournalFilesForFullRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricJournalFilesForFullRecoveryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricJournalFilesForFullRecovery",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricMemoryUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMemoryUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMemoryUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricMqttMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMqttMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMqttMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricNetworkConnectorConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkConnectorConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkConnectorConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricNetworkIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricNetworkOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricOpenTransactionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricOpenTransactionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricOpenTransactionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricOpenwireMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricOpenwireMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricOpenwireMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricProducerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricProducerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricQueueSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricQueueSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricQueueSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricReceiveCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricReceiveCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricReceiveCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricStompMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricStompMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricStompMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricStorePercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricStorePercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTempPercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTempPercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTempPercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTotalConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalConsumerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalConsumerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTotalDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalDequeueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalDequeueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTotalEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalEnqueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalEnqueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTotalMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalMessageCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalMessageCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricTotalProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalProducerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalProducerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricVolumeReadOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricVolumeReadOpsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeReadOps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricVolumeWriteOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricVolumeWriteOpsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeWriteOps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IActiveMqBrokerInstance) MetricWsMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricWsMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricWsMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Endpoints() *ActiveMqBrokerEndpoints {
	var returns *ActiveMqBrokerEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

