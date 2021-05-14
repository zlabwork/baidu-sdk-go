package request

import (
    "os"
)

type credentials struct {
    id     string
    secret string
}

type Client struct {
    credentials credentials
    timeout     int64
    request     requestData
}

func NewClient() *Client {
    id := os.Getenv("BAIDU_ACCESS_ID")
    secret := os.Getenv("BAIDU_ACCESS_SECRET")
    return &Client{
        credentials: credentials{id, secret},
        timeout:     900,
    }
}

func (cli *Client) SetHost(host string) {
    cli.request.host = host
}
