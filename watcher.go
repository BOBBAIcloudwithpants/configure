package configure

import (
	"io/ioutil"
	"os"
	"time"
)

// 监听配置文件时触发的事件
type ListenFunc func(string)

// 监听配置文件使用的结构定义
type Watcher struct {
	Done   chan bool  // 子进程是否已经完成的管道
	Listen ListenFunc // 配置文件发生改变时
	File   *File      // 存储 Watcher 监听的文件的相关信息和解析器
}

// 新建一个监听实体
func newWatcher(done chan bool, f *File, listenFunc ListenFunc) *Watcher {
	return &Watcher{
		Done:   done,
		Listen: listenFunc,
		File:   f,
	}
}

func newWatcherWithOption(done chan bool, f *File, listenFunc ListenFunc, option Option) *Watcher {
	f.parser.option = option
	return &Watcher{
		Done:   done,
		Listen: listenFunc,
		File:   f,
	}
}

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

// 根据文件路径读取内容，返回文件全部内容的字符串
func Read(f string) (string, error) {
	r, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

// 判断 filePath 对应的文件是否被修改过
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
