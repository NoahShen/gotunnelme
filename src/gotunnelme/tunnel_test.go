package gotunnelme

import (
	"fmt"
	"testing"
)

func TestTunnel(t *testing.T) {
	Debug = true
	tunnel := NewTunnel()
	url, getUrlErr := tunnel.GetUrl("")
	if getUrlErr != nil {
		t.Fatal(getUrlErr)
	}
	fmt.Println("Get Url:", url)
	tunnelErr := tunnel.CreateTunnel(8787)
	if tunnelErr != nil {
		t.Fatal(tunnelErr)
	}
}
