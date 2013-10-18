package gotunnelme

import (
	"fmt"
	"testing"
)

func TestGetLatestAqiEntity(t *testing.T) {
	Debug = false
	assignedUrlInfo, err := GetAssignedUrl("noah")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("assignedUrlInfo: %+v\n", assignedUrlInfo)
}
