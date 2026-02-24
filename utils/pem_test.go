package utils

import (
	"fmt"
	"testing"
)

func TestLoadPublicKey(t *testing.T) {
	pubkey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxxhLVkzMlxBrFYsAOBaeTrkVC0CN1DEpKzyV4ju8iWo3gXRqG/ET6fxLCSbf+wtaUEuOewIqq9SQpGpn/X0vsbODcHv3+O59QM4pSrcSeRMbA5sG3PaNl8EN2YS7I+SdmzGVpgPVzBBvGOWN4L1lTibgS52BBYWmV7lTdRTEmVTBT/ofxO0dYTp0Ri0xdZAseyCb6KTl2Ec5DgAQ5Tzz3S2ku2hqsZCZtUa+KSg+VYsWbt+4pZ56+yw0nxEhLmBw688fqgrIvLkQY7fUESrj+8SEFhdp1HJqcqKlcSYnv5RD3jhgOTQc/4r8Rfwgm3Tl0p6mQNoY87vMlTCOZ4zJ4wIDAQAB"
	key, _ := LoadPublicKey(pubkey)
	fmt.Println(key.E)
}
