package vision

import (
    "baidu_sdk/request"
)

const (
    endpoint = "https://aip.baidubce.com"
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
func (vi *vision) Scene() {
    uri := "/rest/2.0/image-classify/v2/advanced_general"
    vi.cli.BuildRequest("GET", uri, nil, "")
}
