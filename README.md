# go-crawler
使用go语言写的一个爬虫

go version: 1.9.7

主要技术：
* 使用redis(HSETNX, HEXISTS)去重
* 用[es](https://www.elastic.co/)存储用户的数据
* 使用`time.Tick`限制抓取速度

如何生成本地模拟证书：
```bash
cd go-crawler
go build -o gencert cert/main.go
./gencert
```
