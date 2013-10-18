package gotunnelme

import (
	"fmt"
	"testing"
)

func TestTunnel(t *testing.T) {
	Debug = true
	tunnel := NewTunnel()
	url, getUrlErr := tunnel.GetUrl("noah")
	if getUrlErr != nil {
		t.Fatal(getUrlErr)
	}
	fmt.Println("Get Url:", url)
	tunnel.CreateTunnel(8787)

}
