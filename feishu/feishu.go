package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	Request struct {
		MsgType string       `json:"msg_type"`
		Content AlertContent `json:"content"`
	}
	AlertContent struct {
		Text string `json:"text"`
	}
)

func NoticeByWebhook(webhookAddr, content string) (string, error) {
	a := AlertContent{
		Text: content,
	}
	r := &Request{
		"text",
		a,
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	reader := bytes.NewReader(b)
	res, err := netClient.Post(webhookAddr, "application/json", reader)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	return string(body), nil
}
