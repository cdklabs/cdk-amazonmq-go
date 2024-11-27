package cdklabscdkamazonmq

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources"
)

// A RabbitMQ Management HTTP API call.
// Experimental.
type RabbitMqApiCall struct {
	// The RabbitMQ Management HTTP API call path.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// A property used to configure logging during lambda function execution.
	//
	// Note: The default Logging configuration is all. This configuration will enable logging on all logged data
	// in the lambda handler. This includes:
	//  - The event object that is received by the lambda handler
	//  - The response received after making a API call
	//  - The response object that the lambda handler will return
	//  - SDK versioning information
	// - Caught and uncaught errors.
	// Default: Logging.all()
	//
	// Experimental.
	Logging customresources.Logging `field:"optional" json:"logging" yaml:"logging"`
	// The HTTP Method used when invoking the RabbitMQ Management HTTP API call.
	// Default: GET.
	//
	// Experimental.
	Method HttpMethods `field:"optional" json:"method" yaml:"method"`
	// Restrict the data returned by the custom resource to specific paths in the API response.
	//
	// Use this to limit the data returned by the custom resource if working with API calls that could potentially result in custom response objects exceeding the hard limit of 4096 bytes.
	// Experimental.
	OutputPaths *[]*string `field:"optional" json:"outputPaths" yaml:"outputPaths"`
	// The payload expected by the RabbitMQ Management HTTP API call.
	// Experimental.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The physical resource id of the custom resource for this call.
	//
	// Mandatory for onCreate call.
	// In onUpdate, you can omit this to passthrough it from request.
	// Default: - no physical resource id.
	//
	// Experimental.
	PhysicalResourceId customresources.PhysicalResourceId `field:"optional" json:"physicalResourceId" yaml:"physicalResourceId"`
}

