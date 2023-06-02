package aigcaas

import (
	"fmt"
	"testing"
)

func TestClient_Visualglm6bSync(t *testing.T) {
	var err error
	var resp *Vitgpt2Response
	req := Vitgpt2Request{
		ImageUrl: "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685163473641xIbXnSPMvVWxF6ee.png?Expires=1685726382&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=ujHff%2Fa7YUcw8EVchWD3tHHlIoQ%3D",
	}
	if resp, err = client.Vitgpt2Sync(&req); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(resp)
}
