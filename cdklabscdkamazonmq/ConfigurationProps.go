package cdklabscdkamazonmq


// Experimental.
type ConfigurationProps struct {
	// Experimental.
	Data *string `field:"required" json:"data" yaml:"data"`
	// Experimental.
	Engine BrokerEngine `field:"required" json:"engine" yaml:"engine"`
	// Experimental.
	AuthenticationStrategy ActiveMqAuthenticationStrategy `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Experimental.
	ConfigurationName *string `field:"optional" json:"configurationName" yaml:"configurationName"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
}

