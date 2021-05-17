package request

import (
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "time"
)

type requestData struct {
    reqTime time.Time
    method  string
    host    string
    uri     string
    header  map[string]string
    body    map[string]string
}

func (req *requestData) httpRequest() ([]byte, error) {
    url := "https://" + req.host + req.uri

    // 1. 准备
    client := resty.New()
    client.SetDebug(true) // TODO :: http debug
    client.SetContentLength(true)
    request := client.R().
        SetHeaders(req.header).
        SetBody([]byte("")) // FIXME :: req.body

    // 2. 请求
    var resp *resty.Response
    var err error
    if req.method == "GET" {
        resp, err = request.Get(url)
    }
    if req.method == "POST" {
        resp, err = request.Post(url)
    }
    if req.method == "HEAD" {
        resp, err = request.Head(url)
    }

    // 3. 结果处理
    if err != nil {
        return nil, err
    }
    if !resp.IsSuccess() {
        return nil, errors.New("request failed: " + url)
    }
    if resp.StatusCode() != 200 {
        return nil, errors.New("response code is " + fmt.Sprintf("%d", resp.StatusCode()))
    }

    return resp.Body(), nil
}
