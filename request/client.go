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
    host        string
    timeout     int64
}

func NewClient() *Client {
    id := os.Getenv("BAIDU_ACCESS_KEY_ID")
    secret := os.Getenv("BAIDU_ACCESS_KEY_SECRET")
    return &Client{
        credentials: credentials{id, secret},
        timeout:     900,
    }
}

func (cli *Client) SetHost(host string) {
    cli.host = host
}
