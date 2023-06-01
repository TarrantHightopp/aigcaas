package aigcaas

type Client struct {
	SecretKey   string `json:"secret_key"`
	SecretId    string `json:"secret_id"`
	Mode        int    `json:"mode"`
	CallbackUrl string `json:"callback_url"`
}

func NewClient(secretKey, secretId string) *Client {
	return &Client{
		SecretKey: secretKey,
		SecretId:  secretId,
	}
}
