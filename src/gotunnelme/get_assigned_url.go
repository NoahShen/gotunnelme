package gotunnelme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	localtunnelServer = "http://localtunnel.me/"
)

type AssignedUrlInfo struct {
	Id           string `json:"id,omitempty"`
	Url          string `json:"url,omitempty"`
	Port         int    `json:"port,omitempty"`
	MaxConnCount int    `json:"max_conn_count,omitempty"`
}

func GetAssignedUrl(assignedDomain string) (*AssignedUrlInfo, error) {
	if len(assignedDomain) == 0 {
		assignedDomain = "?new"
	}
	url := fmt.Sprintf(localtunnelServer+"%s", assignedDomain)
	request, _ := http.NewRequest("GET", url, nil)
	response, httpErr := http.DefaultClient.Do(request)
	if httpErr != nil {
		return nil, httpErr
	}
	defer response.Body.Close()
	bytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}
	if Debug {
		fmt.Printf("***GetAssignedUrl: %s\n", string(bytes))
	}

	assignedUrlInfo := &AssignedUrlInfo{}
	if unmarshalErr := json.Unmarshal(bytes, assignedUrlInfo); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return assignedUrlInfo, nil
}
