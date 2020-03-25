package rule

// SpaceBetweenWordAndPunctuation
// 处理文字和标点符号之间的空格
type SpaceBetweenWordAndPunctuation struct{}

func (s SpaceBetweenWordAndPunctuation) AmendText(source string) (new string, goon, stop bool) {
	panic("implement me")
}

func (s SpaceBetweenWordAndPunctuation) RuleName() string {
	return "space-between-en-and-zh"
}
