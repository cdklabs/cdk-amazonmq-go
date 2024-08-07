package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A representation of a single-instance broker comprised of one broker in one Availability Zone behind a Network Load Balancer (NLB).
// Experimental.
type RabbitMqBrokerInstance interface {
	RabbitMqBrokerDeploymentBase
	IRabbitMqBroker
	// Experimental.
	Arn() *string
	// Experimental.
	Configuration() IRabbitMqBrokerConfiguration
	// Manages connections for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// Experimental.
	Endpoints() *RabbitMqBrokerEndpoints
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RabbitMqBrokerInstance
type jsiiProxy_RabbitMqBrokerInstance struct {
	jsiiProxy_RabbitMqBrokerDeploymentBase
	jsiiProxy_IRabbitMqBroker
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Configuration() IRabbitMqBrokerConfiguration {
	var returns IRabbitMqBrokerConfiguration
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Endpoints() *RabbitMqBrokerEndpoints {
	var returns *RabbitMqBrokerEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqBrokerInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewRabbitMqBrokerInstance(scope constructs.Construct, id *string, props *RabbitMqBrokerInstanceProps) RabbitMqBrokerInstance {
	_init_.Initialize()

	if err := validateNewRabbitMqBrokerInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitMqBrokerInstance{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRabbitMqBrokerInstance_Override(r RabbitMqBrokerInstance, scope constructs.Construct, id *string, props *RabbitMqBrokerInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RabbitMqBrokerInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRabbitMqBrokerInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func RabbitMqBrokerInstance_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRabbitMqBrokerInstance_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func RabbitMqBrokerInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRabbitMqBrokerInstance_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqBrokerInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RabbitMqBrokerInstance) ConfigureLogRetention() {
	_jsii_.InvokeVoid(
		r,
		"configureLogRetention",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitMqBrokerInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) Metric(metricName *string, options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricParameters(metricName, options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metric",
		[]interface{}{metricName, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricAckRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricAckRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricAckRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricChannelCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricChannelCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricChannelCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricConfirmRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricConfirmRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricConfirmRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricConsumerCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricConsumerCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricConsumerCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricExchangeCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricExchangeCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricExchangeCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricMessageCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricMessageCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricMessageCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricMessageReadyCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricMessageReadyCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricMessageReadyCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricMessageUnacknowledgedCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricMessageUnacknowledgedCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricMessageUnacknowledgedCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricPublishRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricPublishRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricPublishRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricQueueCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricQueueCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricQueueCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQDiskFree(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQDiskFreeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQDiskFree",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQDiskFreeLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQDiskFreeLimitParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQDiskFreeLimit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQFdUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQFdUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQFdUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQIOReadAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQIOReadAverageTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQIOReadAverageTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQIOWriteAverageTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQIOWriteAverageTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQIOWriteAverageTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQMemLimit(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQMemLimitParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQMemLimit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricRabbitMQMemUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricRabbitMQMemUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricRabbitMQMemUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) MetricSystemCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricSystemCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricSystemCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqBrokerInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

