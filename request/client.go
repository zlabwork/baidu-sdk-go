package request

import (
    "os"
    "time"
)

type credentials struct {
    id     string
    secret string
}

type Client struct {
    credentials credentials
    timeout     int64
    host        string
    request     *requestData
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
    cli.host = host
}

func (cli *Client) BuildRequest(method, uri string, header map[string]string, body string) {

    // 1. 请求时间
    t := time.Now()

    // 2. header
    ph := pubRequestHeader{
        Authorization: "",
        Length:        "0",
        Type:          "application/x-www-form-urlencoded",
        Md5:           "",
        Date:          t.Format(time.RFC1123Z),
        Host:          cli.host,
        BceDate:       t.Format("2006-01-02T15:04:05Z"),
    }

    // 3. request
    cli.request = &requestData{
        reqTime:   t,
        pubHeader: ph,
        method:    method,
        host:      cli.host,
        uri:       uri,
        header:    header,
        body:      body,
    }

    // 4. authorization
    sign := cli.createSignV1(method, uri)
    cli.request.authorization = cli.authorizationV1(sign)

    // 5. request
    cli.request.httpRequest()
}
