package gotunnelme

import (
	"testing"
)

func TestTunnel(t *testing.T) {
	Debug = true
	err := CreateTunnel("noah", 8787)
	if err != nil {
		t.Fatal(err)
	}

}
