package amend

// ContinuousPunctuation
// 处理连续的标点符号
type ContinuousPunctuation struct{}

func (s ContinuousPunctuation) AmendText(source string) (new string, stop bool) {
	panic("implement me")
}

func (s ContinuousPunctuation) RuleName() string {
	return "space-between-en-and-zh"
}
