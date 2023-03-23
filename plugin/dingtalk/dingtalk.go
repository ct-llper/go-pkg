package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var DingTalker = new(dingTalk)

type dingTalk struct{}

// DingTalkText 钉钉-文本内容
// msg := fmt.Sprintf("类型: %v\n> 信息: %s",
//	messageType,
//	message,
// )
func (u *dingTalk) DingTalkText(secret, webhook string, content string, mobileList []string) (err error) {
	url := dingSign(secret, webhook)

	var (
		data      = make(map[string]interface{})
		dataStr   = make(map[string]string)
		atMobiles = make(map[string]interface{})
	)
	atMobiles["atMobiles"] = mobileList
	dataStr["content"] = content

	data["msgtype"] = "text"
	data["text"] = dataStr
	data["at"] = atMobiles

	b, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)

	return err
}

// DingTalkLink 钉钉-链接
// secret 	SEC10dc2d1914091026ca347e329c34f016e97b1e5190cfcbf14df9b807e02f43b7
// webhook 	https://oapi.dingtalk.com/robot/send?access_token=25540c1c38c7f9d490e0afb121262491ed62a6a2ce5c34488984e635a65d8a5e
// content := make(map[string]string)
// content["title"] = "标题"
// content["text"] = "内容"
// content["messageUrl"] = "跳转地址"
func (u *dingTalk) DingTalkLink(secret, webhook string, content map[string]string) (err error) {
	url := dingSign(secret, webhook)

	data := make(map[string]interface{})
	data["msgtype"] = "link"
	data["link"] = content

	b, _ := json.Marshal(data)
	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(b))

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	return err
}

// Sign 签名
func dingSign(secret, webhook string) string {
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	sign := hmacSha256(stringToSign, secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	return url
}

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
