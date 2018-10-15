# go-crawler
go语言学习项目

go version: 1.9.7

主要技术：
* 使用redis(HSETNX, HEXISTS)去重

如何生成本地模拟证书：
```bash
cd go-crawler
go build -o gencert cert/main.go
./gencert
```
