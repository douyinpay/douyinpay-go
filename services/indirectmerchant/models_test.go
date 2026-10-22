package indirectmerchant

import "testing"

func TestLocatorValidation(t *testing.T) {
	if err := (Locator{IndirectMerchantID: "m"}).Validate(); err != nil {
		t.Fatal(err)
	}
	if err := (Locator{OutIndirectMerchantID: "out"}).Validate(); err == nil {
		t.Fatal("incomplete external locator should fail")
	}
	if err := (Locator{IndirectMerchantID: "m", OutChannelSPID: "sp"}).Validate(); err == nil {
		t.Fatal("channel id without external locator should fail")
	}
}
