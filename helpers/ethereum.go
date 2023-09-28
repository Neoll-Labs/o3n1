package helpers

import "regexp"

func IsValidEthAddress(address string) bool {
	// Define a regular expression pattern for a valid Ethereum address
	ethAddressPattern := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	return ethAddressPattern.MatchString(address)
}
