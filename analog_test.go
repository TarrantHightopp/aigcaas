package aigcaas

import (
	"fmt"
	"testing"
)

func TestClient_Analog(t *testing.T) {
	client := NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)
	commonStableDiffusionRequest := CommonStableDiffusionRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonStableDiffusionResponse
	if commonStableDiffusionResponse, err = client.Analog(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonStableDiffusionResponse)
}
