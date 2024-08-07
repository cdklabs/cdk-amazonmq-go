package cdklabscdkamazonmq


// Experimental.
type ActiveMqBrokerConfigurationProps struct {
	// Experimental.
	Definition ActiveMqBrokerConfigurationDefinition `field:"required" json:"definition" yaml:"definition"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	AuthenticationStrategy ActiveMqAuthenticationStrategy `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Experimental.
	EngineVersion ActiveMqBrokerEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

