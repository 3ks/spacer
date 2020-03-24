package amend

// spaceLink
// 处理链接边界字符之间的空格
type spaceLink struct{}

func (s spaceLink) AmendText(source string) (new string, stop bool) {
	panic("implement me")
}

func (s spaceLink) RuleName() string {
	return "space-between-en-and-zh"
}
