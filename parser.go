package configure

import (
	"regexp"
	"strings"
)


const (
	SectionExp = "\\[[a-zA-Z0-9]*\\]"
)


// 解析配置文件的解析器实体
type Parser struct {
	content string
	option  Option
}


// 解析配置文件，将结果写入 f, 如果有异常则返回异常
func (p *Parser) parse(f *File) (err error){
	var secs []*Section
	lines := strings.Split(p.content, p.option.Separation)
	var secStart []int

	lines = deleteEmpty(lines)

	for i, line := range lines {
		if matched, _ := regexp.MatchString(SectionExp, line); matched {
			secStart = append(secStart, i)
		}
	}

	secName := "default"
	prevSec := 0
	endSec := 0
	// 此时整个分区都是 default 分区
	if len(secStart) == 0 {
		section, err := NewSection(secName, lines, f)
		if err != nil {
			return err
		}
		secs = append(secs, section)
		return nil
	}

	// 若没有默认分区
	if secStart[0] == 0 {
		secName = getSecName(lines[0])
		for i := 0;i<len(secStart);i++ {
			if i == len(secStart) - 1 {
				// 到达最后一个区域，则最后的内容都归为第一个区域
				endSec = len(lines)
			} else {
				endSec = secStart[i+1]
			}
			prevSec := secStart[i]+1
			content := lines[prevSec:endSec]
			secName = getSecName(lines[secStart[i]])
			section, err := NewSection(secName, content, f)
			if err != nil {
				return err
			}
			secs = append(secs, section)
		}
 	} else {
 		// 存在默认分区，第一个分区就是默认分区
 		content := lines[prevSec:secStart[0]]
		section, err := NewSection(secName, content, f)
		if err != nil {
			return err
		}
		secs = append(secs, section)
		for i := 0;i<len(secStart);i++ {
			if i == len(secStart)-1 {
				// 到达最后一个区域，则最后的内容都归为第一个区域
				endSec = len(lines)
			} else {
				endSec = secStart[i+1]
			}
			prevSec := secStart[i]+1
			content := lines[prevSec: endSec]
			secName = getSecName(lines[secStart[i]])
			section, err := NewSection(secName, content,f)
			if err != nil {
				return err
			}
			secs = append(secs, section)
		}
	}
	f.sections = secs
	return nil
}

// 新建一个转换器
func newParser (s string, option Option) *Parser {
	return &Parser{
		option: option,
		content: s,
	}
}



