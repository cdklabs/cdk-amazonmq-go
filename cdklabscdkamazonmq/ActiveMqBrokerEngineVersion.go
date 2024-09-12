package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"
)

// See:  : https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/activemq-version-management.html
//
// Experimental.
type ActiveMqBrokerEngineVersion interface {
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveMqBrokerEngineVersion
type jsiiProxy_ActiveMqBrokerEngineVersion struct {
	_ byte // padding
}

// Experimental.
func NewActiveMqBrokerEngineVersion(version *string) ActiveMqBrokerEngineVersion {
	_init_.Initialize()

	if err := validateNewActiveMqBrokerEngineVersionParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveMqBrokerEngineVersion{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Experimental.
func NewActiveMqBrokerEngineVersion_Override(a ActiveMqBrokerEngineVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		[]interface{}{version},
		a,
	)
}

// Experimental.
func ActiveMqBrokerEngineVersion_Of(version *string) ActiveMqBrokerEngineVersion {
	_init_.Initialize()

	if err := validateActiveMqBrokerEngineVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns ActiveMqBrokerEngineVersion

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func ActiveMqBrokerEngineVersion_V5_15_16() ActiveMqBrokerEngineVersion {
	_init_.Initialize()
	var returns ActiveMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		"V5_15_16",
		&returns,
	)
	return returns
}

func ActiveMqBrokerEngineVersion_V5_16_7() ActiveMqBrokerEngineVersion {
	_init_.Initialize()
	var returns ActiveMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		"V5_16_7",
		&returns,
	)
	return returns
}

func ActiveMqBrokerEngineVersion_V5_17_6() ActiveMqBrokerEngineVersion {
	_init_.Initialize()
	var returns ActiveMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		"V5_17_6",
		&returns,
	)
	return returns
}

func ActiveMqBrokerEngineVersion_V5_18() ActiveMqBrokerEngineVersion {
	_init_.Initialize()
	var returns ActiveMqBrokerEngineVersion
	_jsii_.StaticGet(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerEngineVersion",
		"V5_18",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ActiveMqBrokerEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

