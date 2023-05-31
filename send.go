package aigcaas

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (c *Client) Send(url, body string) (*http.Response, error) {
	var err error
	var request *http.Request
	if request, err = http.NewRequest("POST", url, strings.NewReader(body)); err != nil {
		return nil, err
	}
	var timestamp = time.Now().Unix()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := rand.Intn(10000)
	sign := strconv.FormatInt(timestamp, 10) + c.SecretKey + strconv.Itoa(nonce)
	hash := sha256.New()
	hash.Write([]byte(sign))
	var token = hex.EncodeToString(hash.Sum(nil))
	request.Header.Add("SecretID", c.SecretId)
	request.Header.Add("Nonce", strconv.Itoa(nonce))
	request.Header.Add("Token", token)
	request.Header.Add("Timestamp", strconv.FormatInt(timestamp, 10))
	request.Header.Add("Content-Type", "application/json")
	var client = http.Client{}
	return client.Do(request)
}
