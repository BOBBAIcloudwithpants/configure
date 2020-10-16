# Configure

## 简介
Package `configure` 能够监控，读取符合格式的配置文件并转换为层次 <key, value> pair，在程序中被调用.

## 功能
- 解析配置文件为配置对象
- 从配置对象中读取配置项
- 监听配置文件的改动，在不符合格式时在 `stderr` 中输出错误信息，方便用户作出配置文件的格式调整

## Environment
- go 1.14+

## Installation
```
go get -u github.com/bobbaicloudwithpants/configure
```