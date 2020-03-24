// Spacer 本质上是遍历目录，获取需要处理的文件。对于每一个文件，再遍历每一行的内容。
// 最核心的地方在于每一行的内容如何处理，而 Spacer 将该功能抽象为 `Amend` 接口，
// 由实现了该接口的一系列 `规则` 自己去决定如何处理这一行的内容。
// 如果你想扩展规则，只需要实现 `Amend` 接口即可：
package amend

type Amend interface {
	AmendText(source string) (new string, stop bool)
	RuleName() string
}
