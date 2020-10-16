package configure

import (
	"io/ioutil"
	"os"
)


var (
	DefaultSeparation = "\n"

	DefaultAllowDuplicateSection = false

	DefaultDescriptionIdentity = "#"
)

type Option struct {

	// 如何切割给定的配置文件，默认为按行
	Separation string

	// 是否允许出现名字相同的分区
	AllowDuplicateSection bool

	// 用于标识注释的符号
	DescriptionIdentity string

}

func newOption() Option {
	return Option{
		Separation: DefaultSeparation,
		AllowDuplicateSection: DefaultAllowDuplicateSection,
		DescriptionIdentity: DefaultDescriptionIdentity,
	}
}


func Watch(filename string) (configure *File, err error) {
	r,_ := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	f, _ := os.Open(filename)
	if err != nil {
		return nil, err
	}

	//
	file := newFile(f, string(r))
	err = file.Parse()
	if err != nil {
		LogError(err)
	} else {
		return file, nil
	}

	doneChan := make(chan bool)
	Watcher := NewWatcher(doneChan,file)
	go Watcher.Watch()
	<-doneChan
	return file, nil
}

