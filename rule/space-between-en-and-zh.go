package rule

// SpacesBetweenChineseAndEnglish
// 处理中英文之间的空格
// todo 应该对该规则具体行为进行详细描述。也就是一类规则，应该实现两条规则。例如一条规则加空格，另一条规则不加空格，用户只需要根据需求进行选择。
type SpacesBetweenChineseAndEnglish struct{}

func (s SpacesBetweenChineseAndEnglish) AmendText(source string) (new string, goon, stop bool) {
	panic("implement me")
}

func (s SpacesBetweenChineseAndEnglish) RuleName() string {
	return "space-between-en-and-zh"
}
