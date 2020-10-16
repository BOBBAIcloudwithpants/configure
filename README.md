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

## Usage

### 0. 文件格式
`configure` 目前仅支持下面的配置文件格式：
- linux/macOS 下:
```
[section_name1]

# description
# of
# key-value 1
key1 = val1

# description of key-value 2
key2 = val2

[section_name2]

# description of key-value3
key3 = val3
```
- Windows 下:
```
[section_name1]

; description
; of
; key-value 1
key1 = val1

; description of key-value 2
key2 = val2

[section_name2]

; description of key-value3
key3 = val3
```

- Example:

```
# possible values : production, development
app_mode = development

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana

[server]
# Protocol (http or https)
protocol = http

# The http port  to use
http_port = 9999

# Redirect to correct domain if host header does not match domain
# Prevents DNS rebinding attacks
enforce_domain = true
```

### 1. 解析给定格式的配置文件

```go
package main

import (
	"fmt"
	"gitee.com/baijiadong/service_computing/hw4/configure"
	"log"
)

func defaultListen(filepath string) {
	log.Println(fmt.Sprintf("file '%s' has been changed", filepath))
}

func main() {

	file, err := configure.Watch("test.txt", defaultListen)
	if err == nil {
		// 输出：test.txt
		fmt.Println(file.Filename())

		// 输出：http
		fmt.Println(file.Section("server").Key("protocol").Val())

		// 输出: possible values : production, development
		// 这里，由于 'app_mode' 不属于任何分区，因此被解析到 default 分区中
		// default 分区通过 file.Section("") 或者 filt.Section("default") 来获取
		fmt.Println(file.Section("").Key("app_mode").Description())
	} else {
		configure.LogError(err)
	}
}
```

### 2. 监听配置文件的格式，发生改变并且格式无法解析时能够持续监听并且提示错误信息

假设此时配置文件被修改成如下的内容:
```
# possible values : production, development
app_mode = development
[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data= /home/git/grafana

[server]
# Protocol (hdddgggttp or https)
protocol = http
```
可以看到，这个文件的 `第6行`, 即 `data= /home/git/grafana` 是不符合格式的，因为 `=` 两侧应当都有且仅有一个空格，然后我们运行下面的代码:
```go
package main

import (
	"fmt"
	"github.com/bobbaicloudwithpants/configure"     // import 配置包
)

func defaultListen(filepath string) {
	log.Println(fmt.Sprintf("file '%s' has been changed", filepath))
}

func main() {

  // 假设当前目录下已经存在名为 test.txt 的配置文件，配置文件的内容与上面的 Example 相同
	file, err := configure.Watch("test.txt", defaultListen)
	if err == nil {
    // 输出：test.txt
    fmt.Println(file.Filename())
	} else {
		configure.LogError(err)
	}
}
```
控制台会出现如下输出:
```
2020/10/16 16:56:23 You have unrecognized format in config file 'test.txt', sentence: 'data= /home/git/grafana'
```
同时，进程在 `Watch` 函数被阻塞，这是因为配置文件格式并不正确，需要改成正确的格式才能正常进行。于是我们修改文件为如下:

```
# possible values : production, development
app_mode = development
[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana

[server]
# Protocol (hdddgggttp or https)
protocol = http
```
保存修改之后，控制台出现如下的输出:
```
2020/10/16 16:57:31 file 'test.txt' has been changed
```
并且进程不再被阻塞。