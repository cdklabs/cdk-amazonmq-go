# AWS::AmazonMQ L2 Construct Library

<!--BEGIN STABILITY BANNER-->---


Features                                     | Stability
---------------------------------------------|--------------------------------------------------------
Higher level constructs for ActiveMQ Brokers | ![Experimental](https://img.shields.io/badge/experimental-important.svg?style=for-the-badge)
Higher level constructs for RabbitMQ Bokers  | ![Experimental](https://img.shields.io/badge/experimental-important.svg?style=for-the-badge)

> **Experimental:** Higher level constructs in this module that are marked as experimental are
> under active development. They are subject to non-backward compatible changes or removal in any
> future version. These are not subject to the [Semantic Versioning](https://semver.org/) model and
> breaking changes will be announced in the release notes. This means that while you may use them,
> you may need to update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## Table of Contents

* [Introduction](#introduction)

  * [Security](#security)
* [ActiveMQ Brokers](#activemq-brokers)

  * [ActiveMQ Broker Deployments](#activemq-broker-deployments)
  * [ActiveMQ Broker Endpoints](#activemq-broker-endpoints)
  * [Allowing Connections to ActiveMQ Brokers](#allowing-connections-to-activemq-brokers)
  * [ActiveMQ Broker Configurations](#activemq-broker-configurations)
  * [ActiveMQ Broker User Management](#activemq-broker-user-management)

    * [ActiveMQ Broker Simple Authentication](#activemq-broker-simple-authentication)
    * [ActiveMQ Broker LDAP Integration](#activemq-broker-ldap-integration)
  * [Monitoring ActiveMQ Brokers](#monitoring-activemq-brokers)
  * [ActiveMQ Broker Integration with AWS Lambda](#activemq-broker-integration-with-aws-lambda)
* [RabbitMQ Brokers](#rabbitmq-brokers)

  * [RabbitMQ Broker Deployments](#rabbitmq-broker-deployments)
  * [RabbitMQ Broker Endpoints](#rabbitmq-broker-endpoints)
  * [Allowing Connections to a RabbitMQ Broker](#allowing-connections-to-a-rabbitmq-broker)
  * [RabbitMQ Broker Configurations](#rabbitmq-broker-configurations)
  * [Monitoring RabbitMQ Brokers](#monitoring-rabbitmq-brokers)
  * [RabbitMQ Broker Integration with AWS Lambda](#rabbitmq-broker-integration-with-aws-lambda)

## Introduction

Amazon MQ is a managed service that makes it easy to create and run Apache ActiveMQ and RabbitMQ message brokers at scale. This library brings L2 AWS CDK constructs for Amazon MQ and introduces a notion of *broker deployment* and distincts between *a broker* and *a broker deployment*.

* *broker deployment* represents the configuration that defines how the broker (or a set of brokers in a particular configuration) will be deployed. Effectively, this is the representation of the `AWS::AmazonMQ::Broker` resource type, and will expose the relevant attributes of the resource type (such as ARN, Id).
* *broker* represents the means for accessing the broker, that is its endpoints and (in the case of ActiveMQ) IPv4 address(es).

This stems from the fact that when creating the `AWS::AmazonMQ::Broker` resource for ActiveMQ in the `ACTIVE_STANDBY_MULTI_AZ` deployment mode, the resulting AWS resource will in fact contain a set of two, distinct brokers.

The separation allows for expressing the resources as types in two ways:

* *is*, where a *broker deployment* implements the *broker* behavioral interface
* *has*, where a *broker deployment* contains (a set of) *brokers*.

### Security

In order to build secure solutions follow the guidelines and recommendations in the *[Security](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/using-amazon-mq-securely.html)* section of the AWS documentation for the Amazon MQ.

## ActiveMQ Brokers

Amazon MQ allows for creating AWS-managed ActiveMQ brokers. The brokers enable exchanging messages over [a number of protocols](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker.html#broker-protocols), e.g. AMQP 1.0, OpenWire, STOMP, MQTT.

### ActiveMQ Broker Deployments

The following example creates a minimal, [single-instance ActiveMQ Broker deployment](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/single-broker-deployment.html):

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var brokerUser iSecret


broker := cdklabscdkamazonmq.NewActiveMqBrokerInstance(stack, jsii.String("ActiveMqBroker"), &ActiveMqBrokerInstanceProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.ActiveMqBrokerEngineVersion_V5_18(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_T3, awscdk.InstanceSize_MICRO),
	UserManagement: *cdklabscdkamazonmq.ActiveMqBrokerUserManagement_Simple(&SimpleAuthenticationUserManagementOptions{
		Users: []activeMqUser{
			&activeMqUser{
				Username: brokerUser.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
				Password: brokerUser.*SecretValueFromJson(jsii.String("password")),
			},
		},
	}),
})
```

The example below shows how to instantiate an active-standby redundant pair. `ActiveMqBrokerRedundantPair` doesn't implement `IActiveMqBroker`, but has two properties: `first`, and `second` that do. This stems from the fact that [ActiveMq redundant-pair deployment](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/active-standby-broker-deployment.html) exposes two, separate brokers that work in an active-standby configuration. The names are `first` (instead of `active`) and `second` (instead of `standby`) as there cannot be a guarantee which broker will be the `active` and which - the `standby`.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var brokerUser iSecret
var vpc iVpc
var vpcSubnets subnetSelection


brokerPair := cdklabscdkamazonmq.NewActiveMqBrokerRedundantPair(stack, jsii.String("ActiveMqBrokerPair"), &ActiveMqBrokerRedundantPairProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.ActiveMqBrokerEngineVersion_V5_18(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_M5, awscdk.InstanceSize_LARGE),
	UserManagement: *cdklabscdkamazonmq.ActiveMqBrokerUserManagement_Simple(&SimpleAuthenticationUserManagementOptions{
		Users: []activeMqUser{
			&activeMqUser{
				Username: brokerUser.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
				Password: brokerUser.*SecretValueFromJson(jsii.String("password")),
			},
		},
	}),
	Vpc: Vpc,
	VpcSubnets: VpcSubnets,
})
```

### ActiveMQ Broker Endpoints

Each created broker instance implements `IActiveMqBroker` and has `endpoints` property representing each allowed transport with url and port.

One can use the endpoints as in the example below

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var broker iActiveMqBroker


awscdk.NewCfnOutput(this, jsii.String("AmqpEndpointUrl"), &CfnOutputProps{
	Value: broker.Endpoints.Amqp.Url,
})
awscdk.NewCfnOutput(this, jsii.String("AmqpEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Amqp.Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("StompEndpointUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.Stomp.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("StompEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Stomp.*Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("OpenWireEndpointUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.OpenWire.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("OpenWireEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*OpenWire.*Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("MqttEndpointUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.Mqtt.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("MqttEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Mqtt.*Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("WssEndpointUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.Wss.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("WssEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Wss.*Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("WebConsoleUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.Console.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("WebConsolePort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Console.*Port.toString(),
})

awscdk.NewCfnOutput(this, jsii.String("IpAddress"), &CfnOutputProps{
	Value: broker.IpAddress,
})
```

For the redundant pair deployments one can access all the endpoints under properties `first` and `second`, as each implements `IActiveMqBroker`.

### Allowing Connections to ActiveMQ Brokers

For ActiveMQ broker deployments that are not publically accessible and with specified VPC and subnets you can control who can access the Broker using `connections` attribute. By default no connection is allowed and it has to be explicitly allowed.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var deployment iActiveMqBrokerDeployment
var broker iActiveMqBroker


// for the applications to interact over the STOMP protocol
deployment.Connections.AllowFrom(awscdk.Peer_Ipv4(jsii.String("1.2.3.4/8")), awscdk.Port_Tcp(broker.Endpoints.Stomp.Port))

// for the applications to interact over the OpenWire protocol
deployment.Connections.AllowFrom(awscdk.Peer_Ipv4(jsii.String("1.2.3.4/8")), awscdk.Port_Tcp(broker.Endpoints.OpenWire.Port))

// for the Web Console access
deployment.Connections.AllowFrom(awscdk.Peer_Ipv4(jsii.String("1.2.3.4/8")), awscdk.Port_Tcp(broker.Endpoints.Console.Port))
```

Mind that `connections` will be defined only if VPC and subnets are specified. For an instance of `ActiveMqBrokerRedundantPair` one would access the broker endpoints under either `first` or `second` property.

***Security:*** It is a security best practice *[to block unnecessary protocols with VPC security groups](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/using-amazon-mq-securely.html#amazon-mq-vpc-security-groups)*.

### ActiveMQ Broker Configurations

By default Amazon MQ will create a default configuration for the broker(s) on your deployment. You can introduce custom configurations by explicitly creating one as in the example below:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var brokerUser iSecret
var configurationData string


customConfiguration := cdklabscdkamazonmq.NewActiveMqBrokerConfiguration(stack, jsii.String("CustomConfiguration"), &ActiveMqBrokerConfigurationProps{
	ConfigurationName: jsii.String("ConfigurationName"),
	Description: jsii.String("ConfigurationDescription"),
	EngineVersion: *cdklabscdkamazonmq.ActiveMqBrokerEngineVersion_V5_18(),
	AuthenticationStrategy: *cdklabscdkamazonmq.ActiveMqAuthenticationStrategy_SIMPLE,
	Definition: *cdklabscdkamazonmq.ActiveMqBrokerConfigurationDefinition_Data(configurationData),
})

broker := cdklabscdkamazonmq.NewActiveMqBrokerInstance(stack, jsii.String("Broker"), &ActiveMqBrokerInstanceProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.ActiveMqBrokerEngineVersion_V5_18(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_T3, awscdk.InstanceSize_MICRO),
	UserManagement: *cdklabscdkamazonmq.ActiveMqBrokerUserManagement_Simple(&SimpleAuthenticationUserManagementOptions{
		Users: []activeMqUser{
			&activeMqUser{
				Username: brokerUser.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
				Password: brokerUser.*SecretValueFromJson(jsii.String("password")),
			},
		},
	}),
	Configuration: customConfiguration,
})
```

A configuration can be associated with a specific broker also after the broker creation. Then, it is required to be explicitly associated with the broker.

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var configuration iActiveMqBrokerConfiguration
var deployment iActiveMqBrokerDeployment


configuration.AssociateWith(deployment)
```

This library also allows to modify an existing configuration. Such update of a particular configuration is [creating a new configuration *revision*](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-creating-applying-configurations.html#creating-new-configuration-revision-console) so that a history of revisions can be viewed in the AWS Console. The new revision can be then associated with the broker so it uses it as a working configuration.

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var configuration iActiveMqBrokerConfiguration
var deployment iActiveMqBrokerDeployment
var newData string


newRevision := configuration.CreateRevision(&ActiveMqBrokerConfigurationOptions{
	Description: jsii.String("We need to modify an AuthorizationEntry"),
	Definition: cdklabscdkamazonmq.ActiveMqBrokerConfigurationDefinition_Data(newData),
})

newRevision.AssociateWith(deployment)
```

### ActiveMQ Broker User Management

#### ActiveMQ Broker Simple Authentication

Using ActiveMQ built-in [Simple Authentication](http://activemq.apache.org/security.html#Security-SimpleAuthenticationPlugin) users need to be provided during the broker deployment definition.

***Security:*** In the Simple Authentication User Management authorization is managed in the configuration. It is a security best practice to *[always configure an authorization map](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/using-amazon-mq-securely.html#always-configure-authorization-map)*.

#### ActiveMQ Broker LDAP Integration

Amazon MQ for ActiveMQ enables LDAP integration. An example below shows a minimal setup to configure an Amazon MQ for ActiveMQ broker.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var serviceAccountSecret iSecret


broker := cdklabscdkamazonmq.NewActiveMqBrokerInstance(stack, jsii.String("ActiveMqBrokerInstance"), &ActiveMqBrokerInstanceProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.ActiveMqBrokerEngineVersion_V5_18(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_T3, awscdk.InstanceSize_MICRO),
	UserManagement: *cdklabscdkamazonmq.ActiveMqBrokerUserManagement_Ldap(&LdapUserStoreOptions{
		Hosts: []*string{
			jsii.String("ldap.example.com"),
		},
		UserSearchMatching: jsii.String("uid={0}"),
		UserRoleName: jsii.String("amq"),
		UserBase: jsii.String("ou=users,dc=example,dc=com"),
		RoleBase: jsii.String("ou=roles,dc=example,dc=com"),
		RoleSearchMatching: jsii.String("cn={0}"),
		RoleName: jsii.String("amq"),
		ServiceAccountPassword: serviceAccountSecret.SecretValueFromJson(jsii.String("password")),
		ServiceAccountUsername: serviceAccountSecret.*SecretValueFromJson(jsii.String("username")),
	}),
})
```

### Monitoring ActiveMQ Brokers

This library introduces [a set of metrics](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/security-logging-monitoring-cloudwatch.html#activemq-logging-monitoring) that we can use for the `IActiveMqBrokerDeployment` monitoring. Each can be accessed as a method on the `IActiveMqBrokerDeployment` with the convention `metric[MetricName]`. An example below shows how one can use that:

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var deployment iActiveMqBrokerDeployment


consumerCountMetric := deployment.MetricConsumerCount()
consumerCountMetric.CreateAlarm(stack, jsii.String("ConsumerCountAlarm"), &CreateAlarmOptions{
	Threshold: jsii.Number(100),
	EvaluationPeriods: jsii.Number(3),
	DatapointsToAlarm: jsii.Number(2),
})
```

### ActiveMQ Broker Integration with AWS Lambda

Amazon MQ for ActiveMQ broker queues can be used as event sources for AWS Lambda functions. For authentication only the ActiveMQ SimpleAuthenticationPlugin is supported. Lambda consumes messages using the OpenWire/Java Message Service (JMS) protocol. No other protocols are supported for consuming messages. Within the JMS protocol, only TextMessage and BytesMessage are supported. Lambda also supports JMS custom properties. For more details on the requirements of the integration read [the documentation](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html).

The example below presents an example of creating such an event source mapping:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var target iFunction
var creds iSecret // with username and password fields
var broker iActiveMqBrokerDeployment
var queueName string


target.AddEventSource(cdklabscdkamazonmq.NewActiveMqEventSource(&ActiveMqEventSourceProps{
	Broker: Broker,
	Credentials: creds,
	QueueName: jsii.String(QueueName),
}))
```

***Security:*** When adding an Amazon MQ for ActiveMQ as an AWS Lambda function's event source the library updates the execution role's permissions to satisfy [Amazon MQ requirements for provisioning the event source mapping](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions).

In the case of a private deployment the defined event source mapping will create a set of Elastic Network Interfaces (ENIs) in the subnets in which the broker deployment created communication endpoints. Thus, in order to allow the event source mapping to communicate with the broker one needs to additionally allow inbound traffic from the ENIs on the OpenWire port. As ENIs will use the same security group that governs the access to the broker endpoints you can simply allow communication from the broker's security group to itself on the OpenWire port as in the example below:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var deployment iActiveMqBrokerDeployment
var broker iActiveMqBroker


deployment.Connections.AllowInternally(awscdk.Port_Tcp(broker.Endpoints.OpenWire.Port), jsii.String("Allowing for the ESM"))
```

## RabbitMQ Brokers

Amazon MQ allows for creating AWS-managed RabbitMQ brokers. The brokers enable exchanging messages over AMQP 0-9-1 protocol.

### RabbitMQ Broker Deployments

The following example creates a minimal, single-instance RabbitMQ broker deployment:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var adminSecret iSecret


broker := cdklabscdkamazonmq.NewRabbitMqBrokerInstance(stack, jsii.String("RabbitMqBroker"), &RabbitMqBrokerInstanceProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.RabbitMqBrokerEngineVersion_V3_13(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_T3, awscdk.InstanceSize_MICRO),
	Admin: &Admin{
		Username: adminSecret.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
		Password: adminSecret.*SecretValueFromJson(jsii.String("password")),
	},
})
```

The next example creates a minimal RabbitMQ broker cluster:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var adminSecret iSecret


broker := cdklabscdkamazonmq.NewRabbitMqBrokerCluster(stack, jsii.String("RabbitMqBroker"), &RabbitMqBrokerClusterProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.RabbitMqBrokerEngineVersion_V3_13(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_M5, awscdk.InstanceSize_LARGE),
	Admin: &Admin{
		Username: adminSecret.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
		Password: adminSecret.*SecretValueFromJson(jsii.String("password")),
	},
})
```

### RabbitMQ Broker Endpoints

Each created broker has `endpoints` property with the AMQP endpoint url and port.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var broker iRabbitMqBroker


awscdk.NewCfnOutput(this, jsii.String("AmqpEndpointUrl"), &CfnOutputProps{
	Value: broker.Endpoints.Amqp.Url,
})
awscdk.NewCfnOutput(this, jsii.String("AmqpEndpointPort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Amqp.Port.toString(),
})
awscdk.NewCfnOutput(this, jsii.String("WebConsoleUrl"), &CfnOutputProps{
	Value: broker.*Endpoints.Console.*Url,
})
awscdk.NewCfnOutput(this, jsii.String("WebConsolePort"), &CfnOutputProps{
	Value: broker.*Endpoints.*Console.*Port.toString(),
})
```

### Allowing Connections to a RabbitMQ Broker

For the RabbitMQ broker deployments that are not publically accessible and with specified VPC and subnets you can control who can access the broker using `connections` attribute.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var deployment iRabbitMqBrokerDeployment
var broker iRabbitMqBroker


// for the applications to interact over the AMQP protocol
deployment.Connections.AllowFrom(awscdk.Peer_Ipv4(jsii.String("1.2.3.4/8")), awscdk.Port_Tcp(broker.Endpoints.Amqp.Port))

// for the Web Console access
deployment.Connections.AllowFrom(awscdk.Peer_Ipv4(jsii.String("1.2.3.4/8")), awscdk.Port_Tcp(broker.Endpoints.Console.Port))
```

Mind that `connections` will be defined only if VPC and subnets are specified.

### RabbitMQ Broker Configurations

If you do not specify a custom RabbitMQ Broker configuration, Amazon MQ for RabbitMQ will create a default configuration for the broker on your behalf. You can introduce custom configurations by explicitly creating one as in the example below:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var adminSecret iSecret


customConfiguration := cdklabscdkamazonmq.NewRabbitMqBrokerConfiguration(stack, jsii.String("CustomConfiguration"), &RabbitMqBrokerConfigurationProps{
	ConfigurationName: jsii.String("ConfigurationName"),
	Description: jsii.String("ConfigurationDescription"),
	EngineVersion: *cdklabscdkamazonmq.RabbitMqBrokerEngineVersion_V3_13(),
	Definition: *cdklabscdkamazonmq.RabbitMqBrokerConfigurationDefinition_Parameters(&RabbitMqBrokerConfigurationParameters{
		ConsumerTimeout: awscdk.Duration_Minutes(jsii.Number(20)),
	}),
})

broker := cdklabscdkamazonmq.NewRabbitMqBrokerInstance(stack, jsii.String("Broker"), &RabbitMqBrokerInstanceProps{
	PubliclyAccessible: jsii.Boolean(false),
	Version: *cdklabscdkamazonmq.RabbitMqBrokerEngineVersion_V3_13(),
	InstanceType: awscdk.InstanceType_Of(awscdk.InstanceClass_T3, awscdk.InstanceSize_MICRO),
	Admin: &Admin{
		Username: adminSecret.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
		Password: adminSecret.*SecretValueFromJson(jsii.String("password")),
	},
	Configuration: customConfiguration,
})
```

A configuration can be associated with a specific broker also after the deployment. Then, it is required to be explicitly associated with the broker.

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var configuration iRabbitMqBrokerConfiguration
var deployment iRabbitMqBrokerDeployment


configuration.AssociateWith(deployment)
```

This library also allows to modify an existing configuration. Such update of a particular configuration is [creating a new configuration *revision*](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/rabbitmq-creating-applying-configurations.html#creating-new-rabbitmq-configuration-revision-console) so that a history of revisions can be viewed in the AWS Console. The new revision can be then associated with the broker so it uses it as a working configuration.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var configuration iRabbitMqBrokerConfiguration
var deployment iRabbitMqBrokerDeployment
var newConsumerTimeout duration


newRevision := configuration.CreateRevision(&RabbitMqBrokerConfigurationOptions{
	Description: jsii.String("We need to modify the consumer timeout"),
	Definition: cdklabscdkamazonmq.RabbitMqBrokerConfigurationDefinition_Parameters(&RabbitMqBrokerConfigurationParameters{
		ConsumerTimeout: newConsumerTimeout,
	}),
})

newRevision.AssociateWith(deployment)
```

### Monitoring RabbitMQ Brokers

This library introduces [a set of metrics](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/security-logging-monitoring-cloudwatch.html#rabbitmq-logging-monitoring) that we can use for the `IRabbitMqBrokerDeployment` monitoring. Each can be accessed as a method on the `IRabbitMqBrokerDeployment` with the convention `metric[MetricName]`. An example below shows how one can use that:

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var stack stack
var deployment iRabbitMqBrokerDeployment


consumerCountMetric := deployment.MetricConsumerCount()
consumerCountMetric.CreateAlarm(stack, jsii.String("ConsumerCountAlarm"), &CreateAlarmOptions{
	Threshold: jsii.Number(100),
	EvaluationPeriods: jsii.Number(3),
	DatapointsToAlarm: jsii.Number(2),
})
```

### RabbitMQ Broker Integration with AWS Lambda

Amazon MQ for RabbitMQ broker queues can be used as event sources for AWS Lambda functions. For authentication only the PLAIN authentication mechanism is supported. Lambda consumes messages using the AMQP 0-9-1 protocol. No other protocols are supported for consuming messages. For more details on the requirements of the integration read [the documentation](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html).

The example below presents an example of creating such an event source mapping:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var target iFunction
var creds iSecret // with username and password fields
var broker iRabbitMqBrokerDeployment
var queueName string


target.AddEventSource(cdklabscdkamazonmq.NewRabbitMqEventSource(&RabbitMqEventSourceProps{
	Broker: Broker,
	Credentials: creds,
	QueueName: jsii.String(QueueName),
}))
```

***Security:*** When adding an Amazon MQ for RabbitMQ as an AWS Lambda function's event source the library updates the execution role's permissions to satisfy [Amazon MQ requirements for provisioning the event source mapping](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#events-mq-permissions).

In the case of a private deployment the defined event source mapping will create a set of Elastic Network Interfaces (ENIs) in the subnets in which the broker deployment created communication VPC Endpoints. Thus, in order to allow the event source mapping to communicate with the broekr one needs to additionally allow inbound traffic from the ENIs. As ENIs will use the same security group that governs the access to the VPC Endpoints you can simply allow communication from the broker's security group to itself on the AMQP port as in the example below:

```go
import "github.com/cdklabs/cdk-amazonmq-go/cdklabscdkamazonmq"

var deployment iRabbitMqBrokerDeployment


deployment.Connections.AllowDefaultPortInternally()
```
