package aigcaas

import (
	"fmt"
	"testing"
)

var client = NewClient(`HYOsXaUFOPMZmbY`, `yoZCNgSZeQHLWCp`)

func TestClient_AnalogSync(t *testing.T) {
	commonStableDiffusionRequest := CommonRequest{
		Text: "dog",
	}
	var err error
	var commonStableDiffusionResponse *CommonResponse
	if commonStableDiffusionResponse, err = client.AnalogSync(&commonStableDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonStableDiffusionResponse)
}

func TestClient_AnalogAsync(t *testing.T) {
	var err error
	commonStableDiffusionRequest := CommonRequest{
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
	if resp, err = client.AsyncRequestId(`47CE03F2-0A1E-59DF-B7E6-963903FE0BB9`); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println("message --> ", resp.Message)
	fmt.Println("status --> ", resp.Status)
	fmt.Println("data -->", resp.Data)
	fmt.Println("help --> ", resp.Help)
}
