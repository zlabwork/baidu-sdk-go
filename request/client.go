package request

import (
    "net/url"
    "os"
    "strconv"
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

func (cli *Client) BuildRequest(method, uri string, header map[string]string, body map[string]string) ([]byte, error) {

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
    for k, v := range header {
        headers[strings.ToLower(k)] = v
    }

    // 3. body
    rawBody := ""
    if body != nil {
        for k, v := range body {
            rawBody += "&" + k + "=" + url.QueryEscape(v)
        }
        rawBody = rawBody[1:]
    }
    headers["content-length"] = strconv.Itoa(len(rawBody))

    // 4. request
    cli.request = &requestData{
        reqTime: t,
        method:  method,
        host:    cli.host,
        uri:     uri,
        header:  headers,
        body:    rawBody,
    }

    // 5. authorization
    sign := cli.createSignV1()
    cli.request.header["authorization"] = cli.authorizationV1(sign)

    // 6. request
    return cli.request.httpRequest()
}
