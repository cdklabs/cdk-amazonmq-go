//go:build no_runtime_type_checking

package cdklabscdkamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func validateActiveMqBrokerUserManagement_LdapParameters(options *LdapUserStoreOptions) error {
	return nil
}

func validateActiveMqBrokerUserManagement_SimpleParameters(options *SimpleAuthenticationUserManagementOptions) error {
	return nil
}

