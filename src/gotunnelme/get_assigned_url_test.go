package gotunnelme

import (
	"fmt"
	"testing"
)

func _TestGetAssignedUrl(t *testing.T) {
	Debug = false
	assignedUrlInfo, err := GetAssignedUrl("noah")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("assignedUrlInfo: %+v\n", assignedUrlInfo)
}
