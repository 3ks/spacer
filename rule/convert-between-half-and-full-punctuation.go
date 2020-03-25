package rule

// ConvertBetweenHalfAndFullPunctuation
// 全角和半角标点符号的转换
type ConvertBetweenHalfAndFullPunctuation struct{}

func (s ConvertBetweenHalfAndFullPunctuation) AmendText(source string) (new string, goon, stop bool) {
	panic("implement me")
}

func (s ConvertBetweenHalfAndFullPunctuation) RuleName() string {
	return "space-between-en-and-zh"
}
