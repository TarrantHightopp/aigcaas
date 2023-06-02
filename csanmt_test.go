package aigcaas

import (
	"fmt"
	"testing"
)

// 连续语义增强机器翻译（CSANMT）
// 通过输入中文，可以获得英文翻译（CSANMT连续语义增强机器翻译-中英-通用领域）

func TestClient_CsanmtSync(t *testing.T) {
	csanmtRequest := CsanmtRequest{InputSequence: "你好，请问你是谁"}
	var csanmtResponse *CsanmtResponse
	var err error
	if csanmtResponse, err = client.CsanmtSync(&csanmtRequest); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(csanmtResponse)
}
