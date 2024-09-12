package cdklabscdkamazonmq


// Experimental.
type ActiveMqBrokerConfigurationProps struct {
	// Experimental.
	Definition ActiveMqBrokerConfigurationDefinition `field:"required" json:"definition" yaml:"definition"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	EngineVersion ActiveMqBrokerEngineVersion `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Sets authentication strategy for the broker.
	// Default: - undefined; a SIMPLE authentication strategy will be applied.
	//
	// Experimental.
	AuthenticationStrategy ActiveMqAuthenticationStrategy `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Experimental.
	ConfigurationName *string `field:"optional" json:"configurationName" yaml:"configurationName"`
}

