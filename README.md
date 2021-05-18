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

响应结果:
```json
{
   "result": [
      {
         "score": 0.129719,
         "root": "非自然图像-彩色动漫",
         "keyword": "卡通动漫人物"
      },
      {
         "score": 0.121515,
         "root": "商品-工艺品",
         "keyword": "图像素材"
      },
      {
         "score": 0.08894,
         "root": "商品-绘画",
         "keyword": "图画"
      },
      {
         "score": 0.047998,
         "root": "商品-绘画",
         "keyword": "工笔画"
      },
      {
         "score": 0.007329,
         "root": "商品-农用物资",
         "keyword": "花卉"
      }
   ],
   "log_id": 1394601776337911808,
   "result_num": 5
}
```

## 文档
[文档中心](https://cloud.baidu.com/doc/index.html)  
[入门指南](https://cloud.baidu.com/doc/StartGuide/index.html)  
[鉴权认证](https://cloud.baidu.com/doc/Reference/s/Njwvz1wot)  
[通用物体和场景识别](https://ai.baidu.com/tech/imagerecognition/general)  
