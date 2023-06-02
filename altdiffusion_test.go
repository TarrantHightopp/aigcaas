package aigcaas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_AltDiffusionSync(t *testing.T) {
	var err error
	altDiffusionRequest := &AltDiffusionRequest{
		Prompt: "黑暗精灵公主，非常详细，幻想，非常详细，数字绘画，概念艺术，敏锐的焦点，插图",
	}
	var requestId string
	if requestId, err = client.AltDiffusionAsync(altDiffusionRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(requestId)
}

func TestClient_AltDiffusionAsyncRequestId(t *testing.T) {
	var err error
	var altDiffusionSyncResponse *AltDiffusionSyncResponse
	if altDiffusionSyncResponse, err = client.AltDiffusionAsyncRequestId(`3aa771a2-dfc2-4996-8b5b-776c86cd3384`); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(altDiffusionSyncResponse)
	var altDiffusionSyncInfoResponse AltDiffusionSyncInfoResponse
	if err = json.Unmarshal([]byte(altDiffusionSyncResponse.Info), &altDiffusionSyncInfoResponse); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(altDiffusionSyncInfoResponse)
}
