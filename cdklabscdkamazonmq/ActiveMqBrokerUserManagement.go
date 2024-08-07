package cdklabscdkamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq/jsii"
)

// Experimental.
type ActiveMqBrokerUserManagement interface {
}

// The jsii proxy struct for ActiveMqBrokerUserManagement
type jsiiProxy_ActiveMqBrokerUserManagement struct {
	_ byte // padding
}

// Experimental.
func NewActiveMqBrokerUserManagement() ActiveMqBrokerUserManagement {
	_init_.Initialize()

	j := jsiiProxy_ActiveMqBrokerUserManagement{}

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerUserManagement",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewActiveMqBrokerUserManagement_Override(a ActiveMqBrokerUserManagement) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerUserManagement",
		nil, // no parameters
		a,
	)
}

// Experimental.
func ActiveMqBrokerUserManagement_Ldap(options *LdapUserStoreOptions) IActiveMqBrokerUserManagement {
	_init_.Initialize()

	if err := validateActiveMqBrokerUserManagement_LdapParameters(options); err != nil {
		panic(err)
	}
	var returns IActiveMqBrokerUserManagement

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerUserManagement",
		"ldap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Experimental.
func ActiveMqBrokerUserManagement_Simple(options *SimpleAuthenticationUserManagementOptions) IActiveMqBrokerUserManagement {
	_init_.Initialize()

	if err := validateActiveMqBrokerUserManagement_SimpleParameters(options); err != nil {
		panic(err)
	}
	var returns IActiveMqBrokerUserManagement

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-amazonmq.ActiveMqBrokerUserManagement",
		"simple",
		[]interface{}{options},
		&returns,
	)

	return returns
}

