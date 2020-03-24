package amend

// expandJSON
// 展开 JSON 字符串
type expandJSON struct{}

func (s expandJSON) AmendText(source string) (new string, stop bool) {
	panic("implement me")
}

func (s expandJSON) RuleName() string {
	return "space-between-en-and-zh"
}
