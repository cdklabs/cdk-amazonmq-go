package cdklabscdkamazonmq


// Experimental.
type ActiveMqCloudwatchLogsExports struct {
	// Export audit logs to CloudWatch.
	// Default: - undefined; do not export audit logs.
	//
	// Experimental.
	Audit *bool `field:"optional" json:"audit" yaml:"audit"`
	// Export general logs to CloudWatch.
	// Default: - undefined; do not export general logs.
	//
	// Experimental.
	General *bool `field:"optional" json:"general" yaml:"general"`
}

