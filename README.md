## 安装
```bash
go get github.com/zlabwork/baidu-sdk-go
```

## 使用
```golang
client := vision.NewVision()
url := "https://img1.baidu.com/it/u=2340497325,2166644129&fm=26&fmt=auto&gp=0.jpg"
resp, _ := client.Scene(url)
```

## 文档
[文档中心](https://cloud.baidu.com/doc/index.html)  
[入门指南](https://cloud.baidu.com/doc/StartGuide/index.html)  
[鉴权认证](https://cloud.baidu.com/doc/Reference/s/Njwvz1wot)  
[通用物体和场景识别](https://ai.baidu.com/tech/imagerecognition/general)  
