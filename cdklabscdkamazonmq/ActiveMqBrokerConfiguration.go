package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type ActiveMqBrokerConfiguration interface {
	BrokerConfiguration
	// Experimental.
	Arn() *string
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
	// Experimental.
	Revision() *float64
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
	AssociateWith(broker IActiveMqBrokerDeployment) ConfigurationAssociation
	// Experimental.
	CreateRevision(options *ActiveMqBrokerConfigurationOptions) IActiveMqBrokerConfiguration
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveMqBrokerConfiguration
type jsiiProxy_ActiveMqBrokerConfiguration struct {
	jsiiProxy_BrokerConfiguration
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveMqBrokerConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewActiveMqBrokerConfiguration(scope constructs.Construct, id *string, props *ActiveMqBrokerConfigurationProps) ActiveMqBrokerConfiguration {
	_init_.Initialize()

	if err := validateNewActiveMqBrokerConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveMqBrokerConfiguration{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewActiveMqBrokerConfiguration_Override(a ActiveMqBrokerConfiguration, scope constructs.Construct, id *string, props *ActiveMqBrokerConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		[]interface{}{scope, id, props},
		a,
	)
}

// Experimental.
func ActiveMqBrokerConfiguration_FromAttributes(scope constructs.Construct, logicalId *string, attrs *BrokerConfigurationAttributes) IActiveMqBrokerConfiguration {
	_init_.Initialize()

	if err := validateActiveMqBrokerConfiguration_FromAttributesParameters(scope, logicalId, attrs); err != nil {
		panic(err)
	}
	var returns IActiveMqBrokerConfiguration

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		"fromAttributes",
		[]interface{}{scope, logicalId, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActiveMqBrokerConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func ActiveMqBrokerConfiguration_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerConfiguration_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ActiveMqBrokerConfiguration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActiveMqBrokerConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ActiveMqBrokerConfiguration) AssociateWith(broker IActiveMqBrokerDeployment) ConfigurationAssociation {
	if err := a.validateAssociateWithParameters(broker); err != nil {
		panic(err)
	}
	var returns ConfigurationAssociation

	_jsii_.Invoke(
		a,
		"associateWith",
		[]interface{}{broker},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerConfiguration) CreateRevision(options *ActiveMqBrokerConfigurationOptions) IActiveMqBrokerConfiguration {
	if err := a.validateCreateRevisionParameters(options); err != nil {
		panic(err)
	}
	var returns IActiveMqBrokerConfiguration

	_jsii_.Invoke(
		a,
		"createRevision",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveMqBrokerConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (a *jsiiProxy_ActiveMqBrokerConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
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

func (a *jsiiProxy_ActiveMqBrokerConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

