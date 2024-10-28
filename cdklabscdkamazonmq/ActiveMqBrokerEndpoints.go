package cdklabscdkamazonmq


// Experimental.
type ActiveMqBrokerEndpoints struct {
	// The AMQP endpoint of the broker.
	// Experimental.
	Amqp *BrokerEndpoint `field:"required" json:"amqp" yaml:"amqp"`
	// Experimental.
	Console *BrokerEndpoint `field:"required" json:"console" yaml:"console"`
	// The MQTT endpoint of the broker.
	// Experimental.
	Mqtt *BrokerEndpoint `field:"required" json:"mqtt" yaml:"mqtt"`
	// The OpenWire endpoint of the broker.
	// Experimental.
	OpenWire *BrokerEndpoint `field:"required" json:"openWire" yaml:"openWire"`
	// The STOMP endpoint of the broker.
	// Experimental.
	Stomp *BrokerEndpoint `field:"required" json:"stomp" yaml:"stomp"`
	// The WSS endpoint of the broker.
	// Experimental.
	Wss *BrokerEndpoint `field:"required" json:"wss" yaml:"wss"`
}

