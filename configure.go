package configure

import (
	"io/ioutil"
	"os"
)

var (
	DefaultSeparation          = "\n" // 分隔符，默认为换行符
	DefaultDescriptionIdentity = "#"  // 注释符号，默认为 #
)

type Option struct {

	// 如何切割给定的配置文件，默认为按行
	Separation string

	// 用于标识注释的符号
	DescriptionIdentity string
}

func newOption() Option {
	return Option{
		Separation:          DefaultSeparation,
		DescriptionIdentity: DefaultDescriptionIdentity,
	}
}

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

func WatchWithOption(filename string, listenFunc ListenFunc, option Option) (configure *File, err error) {
	r, _ := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	f, _ := os.Open(filename)
	if err != nil {
		return nil, err
	}

	file := newFile(f, string(r))
	err = file.Parse()
	if err != nil {
		LogError(err)
	} else {
		return file, nil
	}

	doneChan := make(chan bool)
	Watcher := newWatcherWithOption(doneChan, file, listenFunc, option)
	go Watcher.watch()
	<-doneChan
	return file, nil
}
