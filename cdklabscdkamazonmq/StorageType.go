package cdklabscdkamazonmq


// Experimental.
type StorageType string

const (
	// Amazon Elastic Block Store.
	//
	// NOTE: Available only for single-instance ActiveMQ brokers.
	// Experimental.
	StorageType_EBS StorageType = "EBS"
	// Experimental.
	StorageType_EFS StorageType = "EFS"
)

