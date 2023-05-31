package aigcaas

type Client struct {
	SecretKey string `json:"secret_key"`
	SecretId  string `json:"secret_id"`
}

func NewClient(secretKey, secretId string) *Client {
	return &Client{
		SecretKey: secretKey,
		SecretId:  secretId,
	}
}
