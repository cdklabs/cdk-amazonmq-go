package cdklabscdkamazonmq


// Experimental.
type RabbitMqBrokerConfigurationProps struct {
	// Experimental.
	Definition RabbitMqBrokerConfigurationDefinition `field:"required" json:"definition" yaml:"definition"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	EngineVersion RabbitMqBrokerEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

