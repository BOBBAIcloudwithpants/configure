package configure

import (
	"strings"
)

// 配置项中的 key-value pair 的结构体定义
type Item struct {
	name string				// 这个配置项的名称
	val string				// 配置项的值
	description string		// 配置项的描述
	section *Section		// 上一级的Section
}

// 创建一个新的pair
func newItem(section *Section, name string, val string, desp string) *Item {

	return &Item{
		section: section,
		name: name,
		val: val,
		description: desp,
	}
}

// 返回value
func (t *Item) Val() string {
	return t.val
}

// 返回key
func (t *Item) Name() string {
	return t.name
}

// 返回这个配置项的描述
func (t *Item) Description() string {
	return t.description
}

// 为这个配置项设置其所属的区域
func (t *Item) SetSection(s *Section) {
	t.section = s
}

// 新建一个配置项
func NewItem(section *Section, s string, description string) *Item {
	spt := strings.Split(s, " = ")
	return newItem(section, spt[0], spt[1], description)
}






