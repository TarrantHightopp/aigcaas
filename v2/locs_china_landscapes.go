package aigcaas

import (
	"fmt"
	"net/http"
)

// Locs China Landscapes v2
// Locs China Landscapes v2 基于模型 SD 1.5 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS
// https://www.aigcaas.cn/product/detail/111.html

const LocsChinaLandscapesApplicationName = `locsChinaLandscapes_locsChinaLandscapes`

func (c *Client) locsChinaLandscapes(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, LocsChinaLandscapesApplicationName, c.ApiName)
	var response *http.Response
	if response, err = c.Send(url, req); err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if err = c.ReadResponseBody(&commonResponse, response); err != nil {
		return nil, err
	}
	return commonResponse, nil
}

func (c *Client) LocsChinaLandscapesText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.locsChinaLandscapes(commonText2ImgRequest)
}

func (c *Client) LocsChinaLandscapesImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.locsChinaLandscapes(commonImg2ImgRequest)
}
