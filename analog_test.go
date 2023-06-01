package aigcaas

import (
	"fmt"
	"testing"
)

var client = NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)

func TestClient_AnalogSync(t *testing.T) {
	commonStableDiffusionRequest := CommonStableDiffusionRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonStableDiffusionResponse
	if commonStableDiffusionResponse, err = client.AnalogSync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonStableDiffusionResponse)
}

func TestClient_AnalogAsync(t *testing.T) {
	var err error
	commonStableDiffusionRequest := CommonStableDiffusionRequest{
		Text: "dog",
	}
	client.Mode = Polling
	var requestId string
	if requestId, err = client.AnalogAsync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}

	fmt.Println(requestId)
}

func TestClient_AsyncRequestId(t *testing.T) {
	var err error
	var resp *AsyncRequestIdResponse
	if resp, err = client.AsyncRequestId(`4e33b6da-c737-4356-b9ef-9d085b2c5cb4`); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println("message --> ", resp.Message)
	fmt.Println("status --> ", resp.Status)
	fmt.Println("data -->", resp.Data)
	fmt.Println("help --> ", resp.Help)
}
