package formatter

import (
	"fmt"
	"spacer/amend"
	"spacer/rule"
)

var amends map[string]amend.Amend
var enableAmends []string
var disableAmends map[string]bool

func init() {
	// 加载所有规则
	loadAllRule()
	// 加载默认启用的规则
	loadDefaultRule()
	// 加载用户配置启用的规则
	loadEnableRule()
	// 加载用户配置禁用的规则
	loadDisableRule()
}

func Handling(source string) string {
	goon, stop := false, false
	for _, ruleName := range enableAmends {
		// todo 捋一捋先后顺序
		if _, ok := disableAmends[ruleName]; ok {
			continue
		}
		source, goon, stop = amends[ruleName].AmendText(source)
		if goon {
			// todo 从 Handing 的调用方获取两行的内容，可能需要重新设计以下
		}
		if stop {
			break
		}
	}
	return source + "\n"
}

func loadAllRule() {
	amends = make(map[string]amend.Amend)
	amends[rule.SpacesBetweenChineseAndEnglish{}.RuleName()] = rule.SpacesBetweenChineseAndEnglish{}
	// todo 中英文之间的空格
	// todo 文字和半角全角标点符号
	// todo 全角字母数字符号转为半角
	// todo 连续标点符号（允许的，不允许的，特俗处理的）
}

func loadDefaultRule() {
	enableAmends = make([]string, 0)
	enableAmends = append(enableAmends, "a", "b", "c")
}

func loadEnableRule() {
	// todo 临时填充，后续通过配置文件读取
	tmpEnable := []string{"x", "y", "z"}
	for _, v := range tmpEnable {
		if _, ok := amends[v]; ok {
			enableAmends = append(enableAmends, v)
		} else {
			fmt.Printf("ignore invalid enable rule: %v\n", v)
		}
	}
}

func loadDisableRule() {
	// todo 临时填充，后续通过配置文件读取
	tmpDisable := []string{"j", "k", "l"}
	for _, v := range tmpDisable {
		if _, ok := amends[v]; ok {
			disableAmends[v] = true
		} else {
			fmt.Printf("ignore invalid disable rule: %v\n", v)
		}
	}
}
