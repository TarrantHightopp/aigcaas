package aigcaas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 人像卡通化（DCT-Net 模型）
// 输入一张人物图像，实现端到端全图卡通化转换，生成漫画风格虚拟形象，返回风格化后的结果图像。针对不同风格图像。
// 只支持同步调用

var DCTNetApplicationName = "dct"

const (
	DCTNetApiNameCompound       string = `compound`        // 通用风格
	DCTNetApiNameDesignCompound string = `design_compound` // 插画风格
	DCTNetApiName3d             string = `3d`              // 3D风格
	DCTNetApiNameArtstyle       string = `artstyle`        // 艺术风格
	DCTNetApiNameHanddrawn      string = `handdrawn`       // 手绘风格
	DCTNetApiNameWzry           string = `wzry`            // 王者荣耀英雄风格
	DCTNetApiNameMhf            string = `mhf`             // 漫画风格
	DCTNetApiNameSketch         string = `sketch`          // 素描风格

)

func (c *Client) dctNet(req *DCTNetRequest, apiName string) (response *http.Response, err error) {
	var byteInfo = make([]byte, 0)
	if byteInfo, err = json.Marshal(req); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s//%s/api/%s", URL, DCTNetApplicationName, apiName)
	return c.Send(url, string(byteInfo))
}

// DctNetSync 同步调用
func (c *Client) DctNetSync(req *DCTNetRequest, apiName string) (*CommonStableDiffusionResponse, error) {
	var err error
	var response *http.Response
	if response, err = c.dctNet(req, apiName); err != nil {
		return nil, err
	}
	return c.StableDiffusionSync(response)
}
