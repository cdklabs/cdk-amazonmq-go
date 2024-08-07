package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A representation of an active/standby broker that is comprised of two brokers in two different Availability Zones.
//
// see: https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/active-standby-broker-deployment.html
// Experimental.
type ActiveMqBrokerRedundantPair interface {
	ActiveMqBrokerDeploymentBase
	// Experimental.
	Arn() *string
	// Manages connections for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The first broker of the redundant pair for the deployment.
	// Experimental.
	First() IActiveMqBroker
	// Experimental.
	Id() *string
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The second broker of the redundant pair for the deployment.
	// Experimental.
	Second() IActiveMqBroker
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	ConfigureLogRetention()
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Experimental.
	Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveMqBrokerRedundantPair
type jsiiProxy_ActiveMqBrokerRedundantPair struct {
	jsiiProxy_ActiveMqBrokerDeploymentBase
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) First() IActiveMqBroker {
	var returns IActiveMqBroker
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Second() IActiveMqBroker {
	var returns IActiveMqBroker
	_jsii_.Get(
		j,
		"second",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerRedundantPair) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewActiveMqBrokerRedundantPair(scope constructs.Construct, id *string, props *ActiveMqBrokerRedundantPairProps) ActiveMqBrokerRedundantPair {
	_init_.Initialize()

	if err := validateNewActiveMqBrokerRedundantPairParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveMqBrokerRedundantPair{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewActiveMqBrokerRedundantPair_Override(a ActiveMqBrokerRedundantPair, scope constructs.Construct, id *string, props *ActiveMqBrokerRedundantPairProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActiveMqBrokerRedundantPair_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerRedundantPair_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func ActiveMqBrokerRedundantPair_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerRedundantPair_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ActiveMqBrokerRedundantPair_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerRedundantPair_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerRedundantPair",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) ConfigureLogRetention() {
	_jsii_.InvokeVoid(
		a,
		"configureLogRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricParameters(metricName, options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricAmqpMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricAmqpMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricAmqpMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricBurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricBurstBalanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricBurstBalance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricConsumerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricConsumerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricCpuCreditBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricCpuCreditBalanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricCpuCreditBalance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricCurrentConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricCurrentConnectionsCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricCurrentConnectionsCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricDequeueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricDequeueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricDispatchCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricDispatchCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricDispatchCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricEnqueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricEnqueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricEnqueueTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricEnqueueTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricEnqueueTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricEstablishedConnectionsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricEstablishedConnectionsCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricEstablishedConnectionsCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricExpiredCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricExpiredCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricExpiredCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricHeapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHeapUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeapUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricInactiveDurableTopicSubscribersCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricInactiveDurableTopicSubscribersCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricInactiveDurableTopicSubscribersCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricInFlightCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricInFlightCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricInFlightCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricJobSchedulerStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricJobSchedulerStorePercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricJobSchedulerStorePercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricJournalFilesForFastRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricJournalFilesForFastRecoveryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricJournalFilesForFastRecovery",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricJournalFilesForFullRecovery(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricJournalFilesForFullRecoveryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricJournalFilesForFullRecovery",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricMemoryUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricMemoryUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricMemoryUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricMqttMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricMqttMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricMqttMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricNetworkConnectorConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricNetworkConnectorConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricNetworkConnectorConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricNetworkIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricNetworkInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricNetworkIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricNetworkOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricNetworkOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricNetworkOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricOpenTransactionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricOpenTransactionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricOpenTransactionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricOpenwireMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricOpenwireMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricOpenwireMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricProducerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricProducerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricQueueSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricQueueSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricQueueSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricReceiveCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricReceiveCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricReceiveCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricStompMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricStompMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStompMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricStorePercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricStorePercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStorePercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTempPercentUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTempPercentUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTempPercentUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTotalConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTotalConsumerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTotalConsumerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTotalDequeueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTotalDequeueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTotalDequeueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTotalEnqueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTotalEnqueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTotalEnqueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTotalMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTotalMessageCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTotalMessageCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricTotalProducerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTotalProducerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTotalProducerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricVolumeReadOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricVolumeReadOpsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricVolumeReadOps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricVolumeWriteOps(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricVolumeWriteOpsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricVolumeWriteOps",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) MetricWsMaximumConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricWsMaximumConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricWsMaximumConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerRedundantPair) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

