package cdklabscdkamazonmq


// Experimental.
type ConfigurationProps struct {
	// Experimental.
	Data *string `field:"required" json:"data" yaml:"data"`
	// Experimental.
	Engine BrokerEngine `field:"required" json:"engine" yaml:"engine"`
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	AuthenticationStrategy ActiveMqAuthenticationStrategy `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
}

