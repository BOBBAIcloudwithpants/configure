package configure

import (
	"os"
)


// 配置文件的结构体
type File struct {

	file *os.File			// 读取的配置文件的 reader
	sections []*Section	    // 配置文件中包含的 section
	parser *Parser			// 用于解析配置文件的 parser
}

// 创建一个配置文件实例, 关于 parser 的解析时选项使用默认选项
func newFile(file *os.File, content string) *File {
	return &File{
		file: file,
		parser: newParser(content, newOption()),
	}
}

// 创建一个配置文件实例，parser 的解析时选项使用传入的参数 'option'
func newFileWithOpt(file *os.File, content string, option Option) *File {
	return &File{
		file: file,
		parser: newParser(content, option),
	}
}

// 解析配置文件，并且写入 sections 中
func (f * File) Parse() error {
	if err := f.parser.parse(f); err != nil {
		return err
	}
	return nil
}

// 设置 parser 要解析的文件内容
func (f * File) SetContent(content string) {
	f.parser.content = content
}

// 根据 section 的名字寻找对应的分区，如果没有找到则返回 ObjectNotFound 类异常
func (f *File) section(name string) (s *Section,e error) {
	// 如果 name 为空则为寻找 default 分区
	if name == "" {
		name = "default"
	}

		for i := range f.sections {
			if f.sections[i].Name() == name {
				s = f.sections[i]
				return s, nil
			}
		}
		e = ObjectNotFound{Type: "Section", Name: name}
		return nil, e
}

// 同样根据名字来寻找对应的分区，如果这个分区不存在则返回 nil
func (f *File) Section(name string) (s *Section) {
	s, err := f.section(name)
	if err != nil {
		return nil
	} else {
		return s
	}
}

// 获取配置文件的文件名
func (f *File) Filename() string {
	return f.file.Name()
}



