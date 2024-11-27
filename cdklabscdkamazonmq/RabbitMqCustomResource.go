package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Experimental.
type RabbitMqCustomResource interface {
	constructs.Construct
	awsec2.IConnectable
	awsiam.IGrantable
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	GetResponseField(key *string) *string
	// Experimental.
	GetResponseFieldReference(key *string) awscdk.Reference
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RabbitMqCustomResource
type jsiiProxy_RabbitMqCustomResource struct {
	internal.Type__constructsConstruct
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_RabbitMqCustomResource) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqCustomResource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitMqCustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewRabbitMqCustomResource(scope constructs.Construct, id *string, props *RabbitMqCustomResourceProps) RabbitMqCustomResource {
	_init_.Initialize()

	if err := validateNewRabbitMqCustomResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitMqCustomResource{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRabbitMqCustomResource_Override(r RabbitMqCustomResource, scope constructs.Construct, id *string, props *RabbitMqCustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResource",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RabbitMqCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRabbitMqCustomResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqCustomResource) GetResponseField(key *string) *string {
	if err := r.validateGetResponseFieldParameters(key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResponseField",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqCustomResource) GetResponseFieldReference(key *string) awscdk.Reference {
	if err := r.validateGetResponseFieldReferenceParameters(key); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		r,
		"getResponseFieldReference",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitMqCustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

