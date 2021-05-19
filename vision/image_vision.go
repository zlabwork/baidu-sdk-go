package vision

import (
    "github.com/zlabwork/baidu-sdk-go/request"
)

const (
    endpoint = "aip.baidubce.com"
)

type vision struct {
    cli *request.Client
}

func NewVision() *vision {
    v := &vision{
        cli: request.NewClient(),
    }
    v.cli.SetHost(endpoint)
    return v
}

// 场景检测
// @link https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Xk3bcxe21
func (vi *vision) Scene(url string) ([]byte, error) {

    // 1. uri
    uri := "/rest/2.0/image-classify/v2/advanced_general"

    // 2. header
    header := make(map[string]string)
    header["Content-Type"] = "application/x-www-form-urlencoded"

    // 3. body
    body := make(map[string]string)
    body["url"] = url
    body["baike_num"] = "0"

    // 4. result
    resp, err := vi.cli.BuildRequest("POST", uri, header, body)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
