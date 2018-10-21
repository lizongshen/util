package util

import (
	"testing"
)

func TestVerifyHTTP(t *testing.T) {
	if !VerifyProxyIP("	121.31.147.247:8123") {
		t.Errorf("IP ERROR!")
	}
}
