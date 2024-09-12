package cdklabscdkamazonmq


// Experimental.
type ConfigurationAssociationProps struct {
	// Experimental.
	Broker IBrokerDeployment `field:"required" json:"broker" yaml:"broker"`
	// Experimental.
	Configuration IBrokerConfiguration `field:"required" json:"configuration" yaml:"configuration"`
}

