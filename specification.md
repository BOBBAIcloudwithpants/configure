# Configure 设计报告

## 个人信息
|      |          |
| ---- | -------- |
| 姓名 | 白家栋   |
| 学号 | 18342001 |
| 专业 | 软件工程 |

## 源码和API文档地址
- 源码: https://github.com/BOBBAIcloudwithpants/configure.git
- API 文档: https://github.com/BOBBAIcloudwithpants/configure/wiki/Configure-API-Document

## 设计概要
在 Configure 包中，主要为用户提供 `Watch` 接口, 即:
```go
func Watch(filename string, listenFunc ListenFunc) (configure *File, err error)
```
Watch 接口主要实现这两个功能:
- 1. 返回格式正确的配置文件的解析结果
- 2. 当配置文件格式不正确时，能够监听配置文件的修改情况，如果格式正确，则立刻返回


## 设计模式
根据我的类的设计，接口 `Watch` 的工作方式如下:
![](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjrky3uv51j309h0iyq3z.jpg)

## 源码解析
Watch 函数的实现如下:
```go
func Watch(filename string, listenFunc ListenFunc) (configure *File, err error) {
	// 首先读取文件的内容
	r, _ := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 获取文件路径对应的 reader
	f, _ := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// 创建一个file，用于初步解析文件以及接收解析后的结果
	file := newFile(f, string(r))
	err = file.Parse()
	if err != nil {
		LogError(err)
	} else {
		return file, nil
	}

	// 创建Watcher，并且用 doneChan 管道阻塞
	doneChan := make(chan bool)
	Watcher := newWatcher(doneChan, file, listenFunc)
	go Watcher.watch()
	<-doneChan
	return file, nil
}
```

可以看到，通过 `gorouting`, Watcher.watch() 作为一个协程开始监听配置文件的改动。watch 函数的实现如下:
```go
// 监控文件变化，如果配置文件不符合格式，则持续阻塞，否则，返回此时的正确格式解析出的配置文件的 key-value pair
func (w *Watcher) watch() {
	go func(doneChan chan bool) {
		defer func() {
			doneChan <- true
		}()
		for {
			err := watchFile(w.File.Filename())
			w.Listen(w.File.Filename())
			// 此时说明文件已经发生了改变
			if err != nil {
				// 输出读取文件发生的错误，继续阻塞
				LogError(err)
			} else {
				// 读出文件内容
				c, err := Read(w.File.Filename())
				if err != nil {
					LogError(err)
				} else {

					// 为 File 设置最新的文件内容，并且对文件内容 Parse
					w.File.SetContent(c)
					err := w.File.Parse()
					if err != nil {
						// 如果产生错误，说明此时文件格式有误，继续阻塞
						LogError(err)
					} else {
						// 如果没有错误，说明文件格式正确，跳出循环
						// 调用用户定义的listen函数，默认的listen函数仅仅打印一个日志
						break
					}
				}
			}
		}
	}(w.Done)
	<-w.Done
}
```
可以看到，如果 **配置文件格式正确**，则跳出for循环，通过`defer`来向doneChan 传值，这时 watch 结束阻塞，同时主进程也结束阻塞返回 configuration; 如果 **配置文件发生变化后格式不正确**，则输出错误，继续阻塞在 for 循环中。要注意的是，每次文件发生变化，都会调用 `listen` 函数，也就是用户定义的文件发生修改时的事件。判断文件是否发生变化的函数 `watchFile` 实现如下:    
```go
func watchFile(filePath string) error {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
```
下面介绍解析配置文件的实现方式。由于要解析的文件具有固定格式，这次我使用了正则表达式匹配的方式来解析文件。即:
```go
const (
	SectionExp = "\\[[a-zA-Z0-9]*\\]"
)
```
通过 `SectionExp` 来得到每个分区的位置，从而解析出每个分区的内容，然后:
```go
var (
	Description = "# .*"
	Equation = ".* = .*"
)
```
在每个 Section 内，通过 `Description` 来得到 item 的描述；通过 `Equation` 来得到每个 key, value 对，创建新的 Item，放入 section 中。具体实现就不一一贴出源码，您可以自行去源码的 `parser.go`, `section.go` 文件中查看。

## 测试设计和结果

对于每个模块，我设计了相应的测试函数，下面分模块进行介绍。    

### Item

1. 测试正常获取 Item 的 Description
```go
func TestItem_Description(t *testing.T)
```

2. 测试正常获取 Item 的 Name
```go
func TestItem_Name(t *testing.T)
```

3. 测试获取 Item 的 Val
```go
func TestItem_Val(t *testing.T)
```

### Section

1. 测试获取 Section 的名字
```go
func TestSection_Name(t *testing.T)
```

2. 测试解析文件后获取 Item
```go
func TestSection_ItemKeyVal(t *testing.T)
```

3. 测试获取区域下 Item 的描述
```go
func TestSection_Description(t *testing.T)
```

### Parser

1. 测试解析配置文件之后正确获取 default 区域名
```go
func TestParser_ParseName(t *testing.T)
```

2. 测试解析没有默认分区的配置文件正常获取区域名
```go
func TestParser_ParseNameNoDefault(t *testing.T)
```