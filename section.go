package configure

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
)

var (
	Description = "# .*"
	Equation = ".* = .*"
)

func init() {
	if runtime.GOOS == "windows" {
		Description = "; .*"
	} else {
		Description = "# .*"
	}
}

var (
	EmptySection = errors.New("ERROR: This section is empty")
	NoSuchItem = errors.New("ERROR: This section has no item of this name")
	)

type Section struct {
	name string				// 区域名称
	items []*Item			// 该区域下的全部 Item
	file *File				// 该区域所属的配置文件
}

func newSection(name string,  f *File) *Section {
	return &Section{
		name: name,
		file: f,
	}
}

func (s *Section)appendItem(item *Item) {
	for i, _ := range s.items {
		if s.items[i].name == item.name {
			s.items[i].val = item.val
			return
		}
	}
	s.items = append(s.items, item)
}

func (s *Section)Name() string {
	return s.name
}

func (s *Section)AppendInFile(f *File) {
	s.file = f
}


func (s *Section) key(name string) (item *Item, err error) {
	err = nil
	for i := range s.items {
		if s.items[i].Name() == name {
			item = s.items[i]
			return
		}
	}
	err = ObjectNotFound{Name: name, Type: "item"}
	return
}

func (s *Section) Key(name string) (item *Item) {
	item, err := s.key(name)
	if err != nil {
		return nil
	} else {
		return item
	}
}

func NewSection(name string, content []string, f *File) (s *Section, e error){
	s = newSection(name, f)
	desp := ""
	for _, line := range content {

			// 为描述
		if matched, _ := regexp.MatchString(Description, line); matched {
			desp+=splitFirst(line)
		} else if matched, _ := regexp.MatchString(Equation, line); matched {
			// 为创建新的key
			item := NewItem(s, line, desp)

			desp = ""
			s.appendItem(item)
		} else {
			fmt.Println(line)
			// 异常
			e = ConfigFormatError{Sentence: line, Filename: f.file.Name()}
			return nil, e
		}
	}

	return s, nil
}










