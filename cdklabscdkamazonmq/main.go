// @cdklabs/cdk-amazonmq
package cdklabscdkamazonmq

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.ActiveMqAuthenticationStrategy",
		reflect.TypeOf((*ActiveMqAuthenticationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"SIMPLE": ActiveMqAuthenticationStrategy_SIMPLE,
			"LDAP": ActiveMqAuthenticationStrategy_LDAP,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		reflect.TypeOf((*ActiveMqBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "associateWith", GoMethod: "AssociateWith"},
			_jsii_.MemberMethod{JsiiMethod: "createRevision", GoMethod: "CreateRevision"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveMqBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BrokerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationDefinition",
		reflect.TypeOf((*ActiveMqBrokerConfigurationDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ActiveMqBrokerConfigurationDefinition{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationOptions",
		reflect.TypeOf((*ActiveMqBrokerConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfigurationProps",
		reflect.TypeOf((*ActiveMqBrokerConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerDeploymentBase",
		reflect.TypeOf((*ActiveMqBrokerDeploymentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAmqpMaximumConnections", GoMethod: "MetricAmqpMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricBurstBalance", GoMethod: "MetricBurstBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuCreditBalance", GoMethod: "MetricCpuCreditBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentConnectionsCount", GoMethod: "MetricCurrentConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDequeueCount", GoMethod: "MetricDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDispatchCount", GoMethod: "MetricDispatchCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueCount", GoMethod: "MetricEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueTime", GoMethod: "MetricEnqueueTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricEstablishedConnectionsCount", GoMethod: "MetricEstablishedConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExpiredCount", GoMethod: "MetricExpiredCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapUsage", GoMethod: "MetricHeapUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricInactiveDurableTopicSubscribersCount", GoMethod: "MetricInactiveDurableTopicSubscribersCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricInFlightCount", GoMethod: "MetricInFlightCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricJobSchedulerStorePercentUsage", GoMethod: "MetricJobSchedulerStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFastRecovery", GoMethod: "MetricJournalFilesForFastRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFullRecovery", GoMethod: "MetricJournalFilesForFullRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryUsage", GoMethod: "MetricMemoryUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricMqttMaximumConnections", GoMethod: "MetricMqttMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkConnectorConnectionCount", GoMethod: "MetricNetworkConnectorConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkIn", GoMethod: "MetricNetworkIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkOut", GoMethod: "MetricNetworkOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenTransactionCount", GoMethod: "MetricOpenTransactionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenwireMaximumConnections", GoMethod: "MetricOpenwireMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricProducerCount", GoMethod: "MetricProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueSize", GoMethod: "MetricQueueSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricReceiveCount", GoMethod: "MetricReceiveCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricStompMaximumConnections", GoMethod: "MetricStompMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricStorePercentUsage", GoMethod: "MetricStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTempPercentUsage", GoMethod: "MetricTempPercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalConsumerCount", GoMethod: "MetricTotalConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalDequeueCount", GoMethod: "MetricTotalDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalEnqueueCount", GoMethod: "MetricTotalEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalMessageCount", GoMethod: "MetricTotalMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalProducerCount", GoMethod: "MetricTotalProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadOps", GoMethod: "MetricVolumeReadOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteOps", GoMethod: "MetricVolumeWriteOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricWsMaximumConnections", GoMethod: "MetricWsMaximumConnections"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveMqBrokerDeploymentBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BrokerDeploymentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveMqBrokerDeployment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerDeploymentBaseProps",
		reflect.TypeOf((*ActiveMqBrokerDeploymentBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerDeploymentProps",
		reflect.TypeOf((*ActiveMqBrokerDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerDeploymentUserManagementDefinition",
		reflect.TypeOf((*ActiveMqBrokerDeploymentUserManagementDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEndpoints",
		reflect.TypeOf((*ActiveMqBrokerEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		reflect.TypeOf((*ActiveMqBrokerEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ActiveMqBrokerEngineVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerInstance",
		reflect.TypeOf((*ActiveMqBrokerInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAmqpMaximumConnections", GoMethod: "MetricAmqpMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricBurstBalance", GoMethod: "MetricBurstBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuCreditBalance", GoMethod: "MetricCpuCreditBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentConnectionsCount", GoMethod: "MetricCurrentConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDequeueCount", GoMethod: "MetricDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDispatchCount", GoMethod: "MetricDispatchCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueCount", GoMethod: "MetricEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueTime", GoMethod: "MetricEnqueueTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricEstablishedConnectionsCount", GoMethod: "MetricEstablishedConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExpiredCount", GoMethod: "MetricExpiredCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapUsage", GoMethod: "MetricHeapUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricInactiveDurableTopicSubscribersCount", GoMethod: "MetricInactiveDurableTopicSubscribersCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricInFlightCount", GoMethod: "MetricInFlightCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricJobSchedulerStorePercentUsage", GoMethod: "MetricJobSchedulerStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFastRecovery", GoMethod: "MetricJournalFilesForFastRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFullRecovery", GoMethod: "MetricJournalFilesForFullRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryUsage", GoMethod: "MetricMemoryUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricMqttMaximumConnections", GoMethod: "MetricMqttMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkConnectorConnectionCount", GoMethod: "MetricNetworkConnectorConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkIn", GoMethod: "MetricNetworkIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkOut", GoMethod: "MetricNetworkOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenTransactionCount", GoMethod: "MetricOpenTransactionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenwireMaximumConnections", GoMethod: "MetricOpenwireMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricProducerCount", GoMethod: "MetricProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueSize", GoMethod: "MetricQueueSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricReceiveCount", GoMethod: "MetricReceiveCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricStompMaximumConnections", GoMethod: "MetricStompMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricStorePercentUsage", GoMethod: "MetricStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTempPercentUsage", GoMethod: "MetricTempPercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalConsumerCount", GoMethod: "MetricTotalConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalDequeueCount", GoMethod: "MetricTotalDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalEnqueueCount", GoMethod: "MetricTotalEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalMessageCount", GoMethod: "MetricTotalMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalProducerCount", GoMethod: "MetricTotalProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadOps", GoMethod: "MetricVolumeReadOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteOps", GoMethod: "MetricVolumeWriteOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricWsMaximumConnections", GoMethod: "MetricWsMaximumConnections"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveMqBrokerInstance{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ActiveMqBrokerDeploymentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveMqBroker)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerInstanceProps",
		reflect.TypeOf((*ActiveMqBrokerInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		reflect.TypeOf((*ActiveMqBrokerRedundantPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "first", GoGetter: "First"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAmqpMaximumConnections", GoMethod: "MetricAmqpMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricBurstBalance", GoMethod: "MetricBurstBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuCreditBalance", GoMethod: "MetricCpuCreditBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentConnectionsCount", GoMethod: "MetricCurrentConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDequeueCount", GoMethod: "MetricDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDispatchCount", GoMethod: "MetricDispatchCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueCount", GoMethod: "MetricEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueTime", GoMethod: "MetricEnqueueTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricEstablishedConnectionsCount", GoMethod: "MetricEstablishedConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExpiredCount", GoMethod: "MetricExpiredCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapUsage", GoMethod: "MetricHeapUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricInactiveDurableTopicSubscribersCount", GoMethod: "MetricInactiveDurableTopicSubscribersCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricInFlightCount", GoMethod: "MetricInFlightCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricJobSchedulerStorePercentUsage", GoMethod: "MetricJobSchedulerStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFastRecovery", GoMethod: "MetricJournalFilesForFastRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFullRecovery", GoMethod: "MetricJournalFilesForFullRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryUsage", GoMethod: "MetricMemoryUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricMqttMaximumConnections", GoMethod: "MetricMqttMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkConnectorConnectionCount", GoMethod: "MetricNetworkConnectorConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkIn", GoMethod: "MetricNetworkIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkOut", GoMethod: "MetricNetworkOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenTransactionCount", GoMethod: "MetricOpenTransactionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenwireMaximumConnections", GoMethod: "MetricOpenwireMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricProducerCount", GoMethod: "MetricProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueSize", GoMethod: "MetricQueueSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricReceiveCount", GoMethod: "MetricReceiveCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricStompMaximumConnections", GoMethod: "MetricStompMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricStorePercentUsage", GoMethod: "MetricStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTempPercentUsage", GoMethod: "MetricTempPercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalConsumerCount", GoMethod: "MetricTotalConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalDequeueCount", GoMethod: "MetricTotalDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalEnqueueCount", GoMethod: "MetricTotalEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalMessageCount", GoMethod: "MetricTotalMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalProducerCount", GoMethod: "MetricTotalProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadOps", GoMethod: "MetricVolumeReadOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteOps", GoMethod: "MetricVolumeWriteOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricWsMaximumConnections", GoMethod: "MetricWsMaximumConnections"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "second", GoGetter: "Second"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveMqBrokerRedundantPair{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ActiveMqBrokerDeploymentBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPairProps",
		reflect.TypeOf((*ActiveMqBrokerRedundantPairProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerUserManagement",
		reflect.TypeOf((*ActiveMqBrokerUserManagement)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ActiveMqBrokerUserManagement{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqCloudwatchLogsExports",
		reflect.TypeOf((*ActiveMqCloudwatchLogsExports)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ActiveMqEventSource",
		reflect.TypeOf((*ActiveMqEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToSourceAccessConfigurations", GoMethod: "AddToSourceAccessConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "mqType", GoGetter: "MqType"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveMqEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EventSourceBase)
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqEventSourceProps",
		reflect.TypeOf((*ActiveMqEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqLdapAuthorization",
		reflect.TypeOf((*ActiveMqLdapAuthorization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ActiveMqUser",
		reflect.TypeOf((*ActiveMqUser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.Admin",
		reflect.TypeOf((*Admin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.BrokerCloudwatchLogsExports",
		reflect.TypeOf((*BrokerCloudwatchLogsExports)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.BrokerConfiguration",
		reflect.TypeOf((*BrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.BrokerConfigurationAttributes",
		reflect.TypeOf((*BrokerConfigurationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.BrokerDeploymentBase",
		reflect.TypeOf((*BrokerDeploymentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BrokerDeploymentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerDeployment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.BrokerDeploymentBaseProps",
		reflect.TypeOf((*BrokerDeploymentBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.BrokerDeploymentMode",
		reflect.TypeOf((*BrokerDeploymentMode)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_MULTI_AZ": BrokerDeploymentMode_CLUSTER_MULTI_AZ,
			"SINGLE_INSTANCE": BrokerDeploymentMode_SINGLE_INSTANCE,
			"ACTIVE_STANDBY_MULTI_AZ": BrokerDeploymentMode_ACTIVE_STANDBY_MULTI_AZ,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.BrokerDeploymentProps",
		reflect.TypeOf((*BrokerDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.BrokerEndpoint",
		reflect.TypeOf((*BrokerEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.BrokerEngine",
		reflect.TypeOf((*BrokerEngine)(nil)).Elem(),
		map[string]interface{}{
			"RABBITMQ": BrokerEngine_RABBITMQ,
			"ACTIVEMQ": BrokerEngine_ACTIVEMQ,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.ConfigurationAssociation",
		reflect.TypeOf((*ConfigurationAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfigurationAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ConfigurationAssociationProps",
		reflect.TypeOf((*ConfigurationAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.DayOfWeek",
		reflect.TypeOf((*DayOfWeek)(nil)).Elem(),
		map[string]interface{}{
			"MONDAY": DayOfWeek_MONDAY,
			"TUESDAY": DayOfWeek_TUESDAY,
			"WEDNESDAY": DayOfWeek_WEDNESDAY,
			"THURSDAY": DayOfWeek_THURSDAY,
			"FRIDAY": DayOfWeek_FRIDAY,
			"SATURDAY": DayOfWeek_SATURDAY,
			"SUNDAY": DayOfWeek_SUNDAY,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.EventSourceBase",
		reflect.TypeOf((*EventSourceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToSourceAccessConfigurations", GoMethod: "AddToSourceAccessConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "mqType", GoGetter: "MqType"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_EventSourceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.EventSourceBaseProps",
		reflect.TypeOf((*EventSourceBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.EventSourceProps",
		reflect.TypeOf((*EventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.HttpMethods",
		reflect.TypeOf((*HttpMethods)(nil)).Elem(),
		map[string]interface{}{
			"GET": HttpMethods_GET,
			"POST": HttpMethods_POST,
			"PUT": HttpMethods_PUT,
			"DELETE": HttpMethods_DELETE,
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IActiveMqBroker",
		reflect.TypeOf((*IActiveMqBroker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveMqBroker{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IActiveMqBrokerConfiguration",
		reflect.TypeOf((*IActiveMqBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "associateWith", GoMethod: "AssociateWith"},
			_jsii_.MemberMethod{JsiiMethod: "createRevision", GoMethod: "CreateRevision"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IActiveMqBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IActiveMqBrokerDeployment",
		reflect.TypeOf((*IActiveMqBrokerDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAmqpMaximumConnections", GoMethod: "MetricAmqpMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricBurstBalance", GoMethod: "MetricBurstBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuCreditBalance", GoMethod: "MetricCpuCreditBalance"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricCurrentConnectionsCount", GoMethod: "MetricCurrentConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDequeueCount", GoMethod: "MetricDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricDispatchCount", GoMethod: "MetricDispatchCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueCount", GoMethod: "MetricEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEnqueueTime", GoMethod: "MetricEnqueueTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricEstablishedConnectionsCount", GoMethod: "MetricEstablishedConnectionsCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExpiredCount", GoMethod: "MetricExpiredCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeapUsage", GoMethod: "MetricHeapUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricInactiveDurableTopicSubscribersCount", GoMethod: "MetricInactiveDurableTopicSubscribersCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricInFlightCount", GoMethod: "MetricInFlightCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricJobSchedulerStorePercentUsage", GoMethod: "MetricJobSchedulerStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFastRecovery", GoMethod: "MetricJournalFilesForFastRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricJournalFilesForFullRecovery", GoMethod: "MetricJournalFilesForFullRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryUsage", GoMethod: "MetricMemoryUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricMqttMaximumConnections", GoMethod: "MetricMqttMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkConnectorConnectionCount", GoMethod: "MetricNetworkConnectorConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkIn", GoMethod: "MetricNetworkIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricNetworkOut", GoMethod: "MetricNetworkOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenTransactionCount", GoMethod: "MetricOpenTransactionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricOpenwireMaximumConnections", GoMethod: "MetricOpenwireMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricProducerCount", GoMethod: "MetricProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueSize", GoMethod: "MetricQueueSize"},
			_jsii_.MemberMethod{JsiiMethod: "metricReceiveCount", GoMethod: "MetricReceiveCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricStompMaximumConnections", GoMethod: "MetricStompMaximumConnections"},
			_jsii_.MemberMethod{JsiiMethod: "metricStorePercentUsage", GoMethod: "MetricStorePercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTempPercentUsage", GoMethod: "MetricTempPercentUsage"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalConsumerCount", GoMethod: "MetricTotalConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalDequeueCount", GoMethod: "MetricTotalDequeueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalEnqueueCount", GoMethod: "MetricTotalEnqueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalMessageCount", GoMethod: "MetricTotalMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTotalProducerCount", GoMethod: "MetricTotalProducerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeReadOps", GoMethod: "MetricVolumeReadOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricVolumeWriteOps", GoMethod: "MetricVolumeWriteOps"},
			_jsii_.MemberMethod{JsiiMethod: "metricWsMaximumConnections", GoMethod: "MetricWsMaximumConnections"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IActiveMqBrokerDeployment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerDeployment)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IActiveMqBrokerUserManagement",
		reflect.TypeOf((*IActiveMqBrokerUserManagement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveMqBrokerUserManagement{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IBrokerConfiguration",
		reflect.TypeOf((*IBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IBrokerDeployment",
		reflect.TypeOf((*IBrokerDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBrokerDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IRabbitMqBroker",
		reflect.TypeOf((*IRabbitMqBroker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
		},
		func() interface{} {
			return &jsiiProxy_IRabbitMqBroker{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IRabbitMqBrokerConfiguration",
		reflect.TypeOf((*IRabbitMqBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "associateWith", GoMethod: "AssociateWith"},
			_jsii_.MemberMethod{JsiiMethod: "createRevision", GoMethod: "CreateRevision"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRabbitMqBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-amazonmq.IRabbitMqBrokerDeployment",
		reflect.TypeOf((*IRabbitMqBrokerDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAckRate", GoMethod: "MetricAckRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricChannelCount", GoMethod: "MetricChannelCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConfirmRate", GoMethod: "MetricConfirmRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnectionCount", GoMethod: "MetricConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExchangeCount", GoMethod: "MetricExchangeCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageCount", GoMethod: "MetricMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageReadyCount", GoMethod: "MetricMessageReadyCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageUnacknowledgedCount", GoMethod: "MetricMessageUnacknowledgedCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishRate", GoMethod: "MetricPublishRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueCount", GoMethod: "MetricQueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFree", GoMethod: "MetricRabbitMQDiskFree"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFreeLimit", GoMethod: "MetricRabbitMQDiskFreeLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQFdUsed", GoMethod: "MetricRabbitMQFdUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOReadAverageTime", GoMethod: "MetricRabbitMQIOReadAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOWriteAverageTime", GoMethod: "MetricRabbitMQIOWriteAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemLimit", GoMethod: "MetricRabbitMQMemLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemUsed", GoMethod: "MetricRabbitMQMemUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemCpuUtilization", GoMethod: "MetricSystemCpuUtilization"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRabbitMqBrokerDeployment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBrokerDeployment)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.LdapUserStoreOptions",
		reflect.TypeOf((*LdapUserStoreOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.MaintenanceWindowStartTime",
		reflect.TypeOf((*MaintenanceWindowStartTime)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqApiCall",
		reflect.TypeOf((*RabbitMqApiCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerCluster",
		reflect.TypeOf((*RabbitMqBrokerCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAckRate", GoMethod: "MetricAckRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricChannelCount", GoMethod: "MetricChannelCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConfirmRate", GoMethod: "MetricConfirmRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnectionCount", GoMethod: "MetricConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExchangeCount", GoMethod: "MetricExchangeCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageCount", GoMethod: "MetricMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageReadyCount", GoMethod: "MetricMessageReadyCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageUnacknowledgedCount", GoMethod: "MetricMessageUnacknowledgedCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishRate", GoMethod: "MetricPublishRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueCount", GoMethod: "MetricQueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFree", GoMethod: "MetricRabbitMQDiskFree"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFreeLimit", GoMethod: "MetricRabbitMQDiskFreeLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQFdUsed", GoMethod: "MetricRabbitMQFdUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOReadAverageTime", GoMethod: "MetricRabbitMQIOReadAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOWriteAverageTime", GoMethod: "MetricRabbitMQIOWriteAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemLimit", GoMethod: "MetricRabbitMQMemLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemUsed", GoMethod: "MetricRabbitMQMemUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemCpuUtilization", GoMethod: "MetricSystemCpuUtilization"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqBrokerCluster{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RabbitMqBrokerDeploymentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRabbitMqBroker)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerClusterProps",
		reflect.TypeOf((*RabbitMqBrokerClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfiguration",
		reflect.TypeOf((*RabbitMqBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "associateWith", GoMethod: "AssociateWith"},
			_jsii_.MemberMethod{JsiiMethod: "createRevision", GoMethod: "CreateRevision"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BrokerConfiguration)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationDefinition",
		reflect.TypeOf((*RabbitMqBrokerConfigurationDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_RabbitMqBrokerConfigurationDefinition{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationOptions",
		reflect.TypeOf((*RabbitMqBrokerConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationParameters",
		reflect.TypeOf((*RabbitMqBrokerConfigurationParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerConfigurationProps",
		reflect.TypeOf((*RabbitMqBrokerConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerDeploymentBase",
		reflect.TypeOf((*RabbitMqBrokerDeploymentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAckRate", GoMethod: "MetricAckRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricChannelCount", GoMethod: "MetricChannelCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConfirmRate", GoMethod: "MetricConfirmRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnectionCount", GoMethod: "MetricConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExchangeCount", GoMethod: "MetricExchangeCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageCount", GoMethod: "MetricMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageReadyCount", GoMethod: "MetricMessageReadyCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageUnacknowledgedCount", GoMethod: "MetricMessageUnacknowledgedCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishRate", GoMethod: "MetricPublishRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueCount", GoMethod: "MetricQueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFree", GoMethod: "MetricRabbitMQDiskFree"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFreeLimit", GoMethod: "MetricRabbitMQDiskFreeLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQFdUsed", GoMethod: "MetricRabbitMQFdUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOReadAverageTime", GoMethod: "MetricRabbitMQIOReadAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOWriteAverageTime", GoMethod: "MetricRabbitMQIOWriteAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemLimit", GoMethod: "MetricRabbitMQMemLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemUsed", GoMethod: "MetricRabbitMQMemUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemCpuUtilization", GoMethod: "MetricSystemCpuUtilization"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqBrokerDeploymentBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BrokerDeploymentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRabbitMqBroker)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRabbitMqBrokerDeployment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerDeploymentBaseProps",
		reflect.TypeOf((*RabbitMqBrokerDeploymentBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerDeploymentProps",
		reflect.TypeOf((*RabbitMqBrokerDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEndpoints",
		reflect.TypeOf((*RabbitMqBrokerEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerEngineVersion",
		reflect.TypeOf((*RabbitMqBrokerEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_RabbitMqBrokerEngineVersion{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		reflect.TypeOf((*RabbitMqBrokerInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "assignConfigurationIdProperty", GoMethod: "AssignConfigurationIdProperty"},
			_jsii_.MemberMethod{JsiiMethod: "configureLogRetention", GoMethod: "ConfigureLogRetention"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAckRate", GoMethod: "MetricAckRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricChannelCount", GoMethod: "MetricChannelCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConfirmRate", GoMethod: "MetricConfirmRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricConnectionCount", GoMethod: "MetricConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumerCount", GoMethod: "MetricConsumerCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricExchangeCount", GoMethod: "MetricExchangeCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageCount", GoMethod: "MetricMessageCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageReadyCount", GoMethod: "MetricMessageReadyCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricMessageUnacknowledgedCount", GoMethod: "MetricMessageUnacknowledgedCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricPublishRate", GoMethod: "MetricPublishRate"},
			_jsii_.MemberMethod{JsiiMethod: "metricQueueCount", GoMethod: "MetricQueueCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFree", GoMethod: "MetricRabbitMQDiskFree"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQDiskFreeLimit", GoMethod: "MetricRabbitMQDiskFreeLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQFdUsed", GoMethod: "MetricRabbitMQFdUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOReadAverageTime", GoMethod: "MetricRabbitMQIOReadAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQIOWriteAverageTime", GoMethod: "MetricRabbitMQIOWriteAverageTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemLimit", GoMethod: "MetricRabbitMQMemLimit"},
			_jsii_.MemberMethod{JsiiMethod: "metricRabbitMQMemUsed", GoMethod: "MetricRabbitMQMemUsed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSystemCpuUtilization", GoMethod: "MetricSystemCpuUtilization"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqBrokerInstance{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RabbitMqBrokerDeploymentBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRabbitMqBroker)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstanceProps",
		reflect.TypeOf((*RabbitMqBrokerInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqCloudwatchLogsExports",
		reflect.TypeOf((*RabbitMqCloudwatchLogsExports)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResource",
		reflect.TypeOf((*RabbitMqCustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "getResponseField", GoMethod: "GetResponseField"},
			_jsii_.MemberMethod{JsiiMethod: "getResponseFieldReference", GoMethod: "GetResponseFieldReference"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqCustomResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResourcePolicy",
		reflect.TypeOf((*RabbitMqCustomResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "statements", GoGetter: "Statements"},
		},
		func() interface{} {
			return &jsiiProxy_RabbitMqCustomResourcePolicy{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResourceProps",
		reflect.TypeOf((*RabbitMqCustomResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-amazonmq.RabbitMqEventSource",
		reflect.TypeOf((*RabbitMqEventSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToSourceAccessConfigurations", GoMethod: "AddToSourceAccessConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "mqType", GoGetter: "MqType"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_RabbitMqEventSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EventSourceBase)
			_jsii_.InitJsiiProxy(&j.Type__awslambdaIEventSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.RabbitMqEventSourceProps",
		reflect.TypeOf((*RabbitMqEventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-amazonmq.SimpleAuthenticationUserManagementOptions",
		reflect.TypeOf((*SimpleAuthenticationUserManagementOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-amazonmq.StorageType",
		reflect.TypeOf((*StorageType)(nil)).Elem(),
		map[string]interface{}{
			"EBS": StorageType_EBS,
			"EFS": StorageType_EFS,
		},
	)
}
