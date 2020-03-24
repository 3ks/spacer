package amend

// spaceBetweenWordAndPunctuation
// 处理文字和标点符号之间的空格
type spaceBetweenWordAndPunctuation struct{}

func (s spaceBetweenWordAndPunctuation) AmendText(source string) (new string, stop bool) {
	panic("implement me")
}

func (s spaceBetweenWordAndPunctuation) RuleName() string {
	return "space-between-en-and-zh"
}
