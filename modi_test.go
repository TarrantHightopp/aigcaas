package aigcaas

import (
	"fmt"
	"testing"
)

// 迪士尼风格扩散生成模型
// 本模型自 Stable Diffusion 1.5 微调而来，微调数据来自某著名动画工作室的电影截图

func TestClient_ModiAsync(t *testing.T) {
	commonStableDiffusionRequest := CommonStableDiffusionRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonStableDiffusionResponse
	if commonStableDiffusionResponse, err = client.ModiSync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err -> %v", err)
	}

	fmt.Println(commonStableDiffusionResponse)
}

func TestClient_ModiSync(t *testing.T) {
	commonStableDiffusionRequest := CommonStableDiffusionRequest{
		Text: "dog",
	}
	var err error
	var requestId string
	if requestId, err = client.ModiAsync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err -> %v", err)
	}

	fmt.Println(requestId)
}
