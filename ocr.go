package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var OcrApplicationName string = "ocr"

const (
	OcrDeblur = "source" // 原始模拟（输出文字与坐标）
	OcrSimple = "simple" // 简单模式（仅输出文字）
)

func (c *Client) ocr(req *OrcRequest, apiName string) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, OcrApplicationName, apiName)
	return c.Send(url, string(byteInfo))
}

func (c Client) OcrSync(req *OrcRequest, apiName string) {
	var err error
	var response *http.Response
	if response, err = c.ocr(req, apiName); err != nil {
		panic(err)
	}
	var byteInfo = make([]byte, 0)
	if byteInfo, err = io.ReadAll(response.Body); err != nil {
		panic(err)
	}

	fmt.Println(string(byteInfo))
}
