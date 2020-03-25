package rule

// ExpandJSON
// 展开 JSON 字符串
type ExpandJSON struct{}

func (s ExpandJSON) AmendText(source string) (new string, goon, stop bool) {
	panic("implement me")
}

func (s ExpandJSON) RuleName() string {
	return "space-between-en-and-zh"
}
