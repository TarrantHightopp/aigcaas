package aigcaas

import (
	"fmt"
	"net/http"
)

// Dark Sushi Mix 大颗寿司Mix
// Dark Sushi Mix 大颗寿司Mix 基于模型 Other 训练而来，发布在 Civitai，并由第三方开发者上传并分享到 AIGCaaS

const DarkSushiMixApplicationName = `darkSushiMixMix_colorful`

func (c *Client) darkSushiMix(req interface{}) (commonResponse *CommonResponse, err error) {
	var url = fmt.Sprintf("%s/%s/api/%s", URL, DarkSushiMixApplicationName, c.ApiName)
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

func (c *Client) DarkSushiMixText2Img(commonText2ImgRequest *CommonText2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameTEXT2IMG
	return c.darkSushiMix(commonText2ImgRequest)
}

func (c *Client) DarkSushiMixImg2Img(commonImg2ImgRequest *CommonImg2ImgRequest) (commonResponse *CommonResponse, err error) {
	c.ApiName = ApiNameIMG2IMG
	return c.darkSushiMix(commonImg2ImgRequest)
}
