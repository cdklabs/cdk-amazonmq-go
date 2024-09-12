package cdklabscdkamazonmq


// Experimental.
type RabbitMqBrokerConfigurationProps struct {
	// Experimental.
	Definition RabbitMqBrokerConfigurationDefinition `field:"required" json:"definition" yaml:"definition"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	EngineVersion RabbitMqBrokerEngineVersion `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Experimental.
	ConfigurationName *string `field:"optional" json:"configurationName" yaml:"configurationName"`
}

