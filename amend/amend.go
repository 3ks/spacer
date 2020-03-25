// Spacer 本质上是遍历目录，获取需要处理的文件。对于每一个文件，再遍历每一行的内容。
// 最核心的地方在于每一行的内容如何处理，而 Spacer 将该功能抽象为 `Amend` 接口，
// 由实现了该接口的一系列 `规则` 自己去决定如何处理这一行的内容。
// 如果你想扩展规则，只需要实现 `Amend` 接口即可：
package amend

type Amend interface {
	// AmendText 的三个返回值分别代表：处理后的内容，需要处理多行内容，阻止后续其它的规则的处理。
	// 其中 goon 一般用于需要处理多行的情况，例如跨行的链接、表格等。
	// stop 表示阻止后续其它的规则的处理，并立即返回本行内容。
	AmendText(source string) (new string, goon, stop bool)
	RuleName() string
}
