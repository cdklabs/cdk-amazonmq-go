package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The IAM Policy that will be applied to the calls.
// Experimental.
type RabbitMqCustomResourcePolicy interface {
	// statements for explicit policy.
	// Experimental.
	Statements() *[]awsiam.PolicyStatement
}

// The jsii proxy struct for RabbitMqCustomResourcePolicy
type jsiiProxy_RabbitMqCustomResourcePolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_RabbitMqCustomResourcePolicy) Statements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"statements",
		&returns,
	)
	return returns
}


// Explicit IAM Policy Statements.
// Experimental.
func RabbitMqCustomResourcePolicy_FromStatements(statements *[]awsiam.PolicyStatement) RabbitMqCustomResourcePolicy {
	_init_.Initialize()

	if err := validateRabbitMqCustomResourcePolicy_FromStatementsParameters(statements); err != nil {
		panic(err)
	}
	var returns RabbitMqCustomResourcePolicy

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResourcePolicy",
		"fromStatements",
		[]interface{}{statements},
		&returns,
	)

	return returns
}

func RabbitMqCustomResourcePolicy_ANY_RESOURCE() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.RabbitMqCustomResourcePolicy",
		"ANY_RESOURCE",
		&returns,
	)
	return returns
}

