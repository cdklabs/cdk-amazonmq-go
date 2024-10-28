package cdklabscdkamazonmq


// Amazon MQ for ActiveMQ's authentication strategy.
// Experimental.
type ActiveMqAuthenticationStrategy string

const (
	// Experimental.
	ActiveMqAuthenticationStrategy_SIMPLE ActiveMqAuthenticationStrategy = "SIMPLE"
	// Experimental.
	ActiveMqAuthenticationStrategy_LDAP ActiveMqAuthenticationStrategy = "LDAP"
)

