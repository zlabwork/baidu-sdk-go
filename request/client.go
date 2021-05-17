package request

import (
    "os"
    "strings"
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
        timeout:     1800,
    }
}

func (cli *Client) SetHost(host string) {
    cli.host = host
}

func (cli *Client) BuildRequest(method, uri string, header map[string]string, params map[string]string) ([]byte, error) {

    // 1. 请求时间
    loc, _ := time.LoadLocation("")
    t := time.Now().In(loc)

    // 2. 合并 headers
    headers := make(map[string]string)
    headers["authorization"] = ""
    headers["host"] = cli.host
    headers["date"] = t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
    headers["x-bce-date"] = t.Format("2006-01-02T15:04:05Z")
    headers["content-type"] = "application/x-www-form-urlencoded"
    headers["content-length"] = "0"
    for k, v := range header {
        headers[strings.ToLower(k)] = v
    }

    // 3. request
    cli.request = &requestData{
        reqTime: t,
        method:  method,
        host:    cli.host,
        uri:     uri,
        header:  headers,
        body:    params,
    }

    // 4. authorization
    sign := cli.createSignV1()
    cli.request.header["authorization"] = cli.authorizationV1(sign)

    // 5. request
    return cli.request.httpRequest()
}
