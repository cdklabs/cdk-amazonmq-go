package cdklabscdkamazonmq


// Experimental.
type RabbitMqBrokerEndpoints struct {
	// Experimental.
	Amqp *BrokerEndpoint `field:"required" json:"amqp" yaml:"amqp"`
	// Experimental.
	Console *BrokerEndpoint `field:"required" json:"console" yaml:"console"`
}

