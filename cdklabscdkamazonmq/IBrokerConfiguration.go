package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/internal"
)

// Experimental.
type IBrokerConfiguration interface {
	awscdk.IResource
	// Experimental.
	Arn() *string
	// Experimental.
	Id() *string
	// Experimental.
	Revision() *float64
}

// The jsii proxy for IBrokerConfiguration
type jsiiProxy_IBrokerConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IBrokerConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrokerConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrokerConfiguration) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

