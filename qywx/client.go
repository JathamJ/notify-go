package qywx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	pushUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
)

type QyWx struct {
	Key    string
	Client *http.Client
}

type Response struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func NewQyWx(key string) *QyWx {
	return &QyWx{
		Key:    key,
		Client: &http.Client{},
	}
}

func (c *QyWx) Send(params interface{}) error {
	url := pushUrl + "?key=" + c.Key
	method := http.MethodPost

	paramStr, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("json Marshal failed, err: %v", err)
	}

	payload := strings.NewReader(string(paramStr))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return fmt.Errorf("NewRequest failed, err: %v", err)
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed, err: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ReadAll failed, err: %v", err)
	}

	result := new(Response)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("unmarshal failed, err: %v", err)
	}
	if result.Code != 0 {
		return fmt.Errorf("request error, code: %d, msg: %s", result.Code, result.Msg)
	}
	return nil
}
