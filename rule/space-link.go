package rule

// SpaceLink
// 处理链接边界字符之间的空格
type SpaceLink struct{}

func (s SpaceLink) AmendText(source string) (new string, goon, stop bool) {
	panic("implement me")
}

func (s SpaceLink) RuleName() string {
	return "space-between-en-and-zh"
}
