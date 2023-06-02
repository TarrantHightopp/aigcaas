package aigcaas

import (
	"fmt"
	"testing"
)

func TestClient_DctNetSync(t *testing.T) {
	var err error
	dCTNetRequest := DCTNetRequest{
		ImageUrl:    "https://x-ai-test.oss-cn-hangzhou.aliyuncs.com/images/uploads/2023-05-27/1685184985115EseK2F6MkAPO09vQ.png?Expires=1685698590&OSSAccessKeyId=TMP.3Kjy6V2sBYt95J6rdre4yESQqWJ7sZNxhKrw6mMiKhCxbLSLn6Sy5JRLqpP2rrpCoBRVNoWHgpFviDE2Re3RBo1qjhW1CB&Signature=S%2F6lGXunqDDL%2FDI0p1Qd6VeoWlQ%3D",
		ImageBase64: "",
	}
	var commonResponse *CommonResponse
	if commonResponse, err = client.DctNetSync(&dCTNetRequest, DCTNetApiNameCompound); err != nil {
		t.Errorf("err --> %v", err)
	}
	fmt.Println(commonResponse)
}
