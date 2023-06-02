package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	FacefuApplicationName string = "facefu"
	FacefuApiName         string = "default"
)

func (c *Client) facefu(req *FacefuRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, FacefuApplicationName, FacefuApiName)
	return c.Send(url, string(byteInfo))
}

// FacefuSync 同步
func (c *Client) FacefuSync(req *FacefuRequest) (commonStableDiffusionResponse *CommonResponse, err error) {
	var response *http.Response
	if response, err = c.facefu(req); err != nil {
		return nil, err
	}
	return c.Sync(response)
}
