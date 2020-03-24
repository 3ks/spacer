package amend

// convertBetweenHalfAndFullPunctuation
// 全角和半角标点符号的转换
type convertBetweenHalfAndFullPunctuation struct{}

func (s convertBetweenHalfAndFullPunctuation) AmendText(source string) (new string, stop bool) {
	panic("implement me")
}

func (s convertBetweenHalfAndFullPunctuation) RuleName() string {
	return "space-between-en-and-zh"
}
