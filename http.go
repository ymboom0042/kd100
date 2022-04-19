// Package src /**
package mian

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *Client) Post(method string, values url.Values) (res []byte, err error) {
	reqUrl := fmt.Sprintf("%s/%s", gateway, method)
	resp, err := c.client.PostForm(reqUrl, values)
	if err != nil {
		fmt.Println(err)
		return nil, newErr("请求失败")
	}

	defer resp.Body.Close()
	res, _ = ioutil.ReadAll(resp.Body)
	return
}
