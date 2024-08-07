package cdklabscdkamazonmq


// Experimental.
type RabbitMqCloudwatchLogsExports struct {
	// Export general logs to CloudWatch.
	// Default: - undefined; do not export general logs.
	//
	// Experimental.
	General *bool `field:"optional" json:"general" yaml:"general"`
}

