package configure

// 去掉读取到程序里的文件内容中多余的换行
func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// 去掉注释行首的符号
func splitFirst(str string) string {
	str = str[2:]
	return str
}

// 从中括号中获得 section name
func getSecName(str string) string {
	str = str[1 : len(str)-1]
	return str
}
