package cdklabscdkamazonmq


// Experimental.
type BrokerEndpoint struct {
	// The port at which the endpoint awaits communication.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The full URL of the broker endpoint, including the port.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
}

