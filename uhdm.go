package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 图像去摩尔纹（uhdm）
// 给定一张输入带有摩尔纹的图像，输出去除摩尔纹图像；拍摄数字屏幕上显示的内容时，相机的颜色滤波器阵列(CFA)和屏幕LCD亚像素之间存在频率堆叠效应，呈现的图像混合了彩色的条纹，此类图像称为摩尔纹图像。

const (
	UhdmApplicationName string = "uhdm"
	UhdmApiName         string = "default"
)

func (c *Client) uhdm(req *UhdmRequest) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, UhdmApplicationName, UhdmApiName)
	return c.Send(url, string(byteInfo))
}

// UhdmSync 同步调用
func (c *Client) UhdmSync(req *UhdmRequest) (*CommonResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.uhdm(req); err != nil {
		return nil, err
	}
	return c.Sync(response)
}
