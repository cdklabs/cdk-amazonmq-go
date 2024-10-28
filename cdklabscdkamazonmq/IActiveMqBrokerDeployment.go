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
type IActiveMqBrokerDeployment interface {
	IBrokerDeployment
	awscdk.IResource
	// Experimental.
	MetricAmqpMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricBurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricCpuCreditBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricCurrentConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricDispatchCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricEnqueueTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricEstablishedConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricExpiredCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricHeapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricInactiveDurableTopicSubscribersCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricInFlightCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricJobSchedulerStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricJournalFilesForFastRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricJournalFilesForFullRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricMemoryUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricMqttMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricNetworkConnectorConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricNetworkIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricNetworkOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricOpenTransactionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricOpenwireMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricQueueSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricReceiveCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricStompMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTempPercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTotalConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTotalDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTotalEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTotalMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricTotalProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricVolumeReadOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricVolumeWriteOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	MetricWsMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy for IActiveMqBrokerDeployment
type jsiiProxy_IActiveMqBrokerDeployment struct {
	jsiiProxy_IBrokerDeployment
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricAmqpMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricBurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricCpuCreditBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricCurrentConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricDispatchCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricEnqueueTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricEstablishedConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricExpiredCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricHeapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricInactiveDurableTopicSubscribersCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricInFlightCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricJobSchedulerStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricJournalFilesForFastRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricJournalFilesForFullRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricMemoryUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricMqttMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricNetworkConnectorConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricNetworkIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricNetworkOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricOpenTransactionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricOpenwireMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricQueueSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricReceiveCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricStompMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTempPercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTotalConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTotalDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTotalEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTotalMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricTotalProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricVolumeReadOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricVolumeWriteOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) MetricWsMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IActiveMqBrokerDeployment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IActiveMqBrokerDeployment) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IActiveMqBrokerDeployment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveMqBrokerDeployment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

