package cdklabscdkamazonmq


// Experimental.
type RabbitMqBrokerConfigurationOptions struct {
	// Experimental.
	Definition RabbitMqBrokerConfigurationDefinition `field:"required" json:"definition" yaml:"definition"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

