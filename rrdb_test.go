package aigcaas

import (
	"fmt"
	"testing"
)

// 图像去色带（RealESRGAN）
// 适用于因色彩精度不够导致色带的图像，能将色带去除。

func TestClient_RrdbSync(t *testing.T) {
	var err error
	var resp *CommonResponse
	req := &RrdbRequest{
		ImageUrl: "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685184807947dLCWr4BDegD7wvRf.png?Expires=1685725887&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=fh0f8VdWYT18ygTv%2B%2B6iNCb%2BYPc%3D",
	}
	if resp, err = client.RrdbSync(req); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(resp)
	fmt.Println("status --> ", resp.Status)
}
