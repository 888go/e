package e

import (
	"fmt"
	"reflect"
	"testing"
)

func TestE取文本长度(t *testing.T) {
	cases := []struct {
		input    string
		expected int64
	}{
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "Hello, 世界！",
			expected: 10,
		},
		{
			input:    "Go语言",
			expected: 4,
		},
	}

	for _, c := range cases {
		actual := X取文本长度(c.input)
		if actual != c.expected {
			t.Errorf("X取文本长度(%q) == %d, expected %d", c.input, actual, c.expected)
		}
	}
}

func TestE取文本左边(t *testing.T) {
	cases := []struct {
		输入文本   string
		欲取出字符数 int64
		预期结果   string
	}{
		{
			输入文本:   "Hello, 世界！",
			欲取出字符数: 7,
			预期结果:   "Hello, ",
		},
		{
			输入文本:   "Go语言",
			欲取出字符数: 3,
			预期结果:   "Go语",
		},
		{
			输入文本:   "Go语言",
			欲取出字符数: 10,
			预期结果:   "Go语言",
		},
	}

	for _, c := range cases {
		实际结果 := X取文本左边(c.输入文本, c.欲取出字符数)
		if 实际结果 != c.预期结果 {
			t.Errorf("X取文本左边(%q, %d) == %q, 预期结果 %q", c.输入文本, c.欲取出字符数, 实际结果, c.预期结果)
		}
	}
}

func TestE取文本右边(t *testing.T) {
	cases := []struct {
		输入文本   string
		欲取出字符数 int64
		预期结果   string
	}{
		{
			输入文本:   "Hello, 世界！",
			欲取出字符数: 7,
			预期结果:   "lo, 世界！",
		},
		{
			输入文本:   "Go语言",
			欲取出字符数: 3,
			预期结果:   "o语言",
		},
		{
			输入文本:   "Go语言",
			欲取出字符数: 10,
			预期结果:   "Go语言",
		},
	}

	for _, c := range cases {
		实际结果 := X取文本右边(c.输入文本, c.欲取出字符数)
		if 实际结果 != c.预期结果 {
			t.Errorf("X取文本右边(%q, %d) == %q, 预期结果 %q", c.输入文本, c.欲取出字符数, 实际结果, c.预期结果)
		}
	}
}

func TestE取文本中间(t *testing.T) {
	cases := []struct {
		输入文本   string
		起始位置   int64
		欲取出字符数 int64
		预期结果   string
	}{
		{
			输入文本:   "Hello, 世界！",
			起始位置:   7,
			欲取出字符数: 3,
			预期结果:   "世界！",
		},
		{
			输入文本:   "Go语言",
			起始位置:   2,
			欲取出字符数: 2,
			预期结果:   "语言",
		},
		{
			输入文本:   "Go语言",
			起始位置:   10,
			欲取出字符数: 10,
			预期结果:   "",
		},
	}

	for _, c := range cases {
		实际结果 := X取文本中间(c.输入文本, c.起始位置, c.欲取出字符数)
		if 实际结果 != c.预期结果 {
			t.Errorf("X取文本中间(%q, %d, %d) == %q, 预期结果 %q", c.输入文本, c.起始位置, c.欲取出字符数, 实际结果, c.预期结果)
		}
	}
}
func TestE字符(t *testing.T) {
	tests := []struct {
		输入   int8
		期望输出 string
	}{
		{65, "A"},
		{97, "X父类"},
		{33, "!"},
	}

	for _, tt := range tests {
		if 输出 := X字符(tt.输入); 输出 != tt.期望输出 {
			t.Errorf("X字符(%v) = %v，期望输出 %v", tt.输入, 输出, tt.期望输出)
		}
	}
}

func TestE取代码(t *testing.T) {
	tests := []struct {
		欲取字符代码的文本 string
		期望输出      int
	}{
		{"Hello, 世界", 72},
		{"Hello, 世界", 72},
		{"Hello, 世界", 72},
	}

	for _, tt := range tests {
		if 输出 := X取代码(tt.欲取字符代码的文本); 输出 != tt.期望输出 {
			t.Errorf("X取代码(%v) = %v，期望输出 %v", tt.欲取字符代码的文本, 输出, tt.期望输出)
		}
	}
}

func TestE寻找文本(t *testing.T) {
	tests := []struct {
		被搜寻的文本   string
		欲寻找的文本   string
		起始搜寻位置   int
		是否不区分大小写 bool
		期望结果     int
	}{
		{
			被搜寻的文本:   "hello world",
			欲寻找的文本:   "lo",
			起始搜寻位置:   0,
			是否不区分大小写: false,
			期望结果:     4,
		},
		{
			被搜寻的文本:   "GoLang编程",
			欲寻找的文本:   "lang",
			起始搜寻位置:   0,
			是否不区分大小写: true,
			期望结果:     3,
		},
		{
			被搜寻的文本:   "",
			欲寻找的文本:   "test",
			起始搜寻位置:   0,
			是否不区分大小写: false,
			期望结果:     -1,
		},
	}

	for _, test := range tests {
		if X寻找文本(test.被搜寻的文本, test.欲寻找的文本, test.起始搜寻位置, test.是否不区分大小写) != test.期望结果 {
			t.Errorf("X寻找文本(%v, %v, %v, %v) 结果错误, 期望结果为 %v", test.被搜寻的文本, test.欲寻找的文本, test.起始搜寻位置, test.是否不区分大小写, test.期望结果)
		}
	}
}

func TestE倒找文本(t *testing.T) {
	tests := []struct {
		被搜寻的文本   string
		欲寻找的文本   string
		起始搜寻位置   int
		是否不区分大小写 bool
		期望结果     int
	}{
		{
			被搜寻的文本:   "hello world",
			欲寻找的文本:   "lo",
			起始搜寻位置:   0,
			是否不区分大小写: false,
			期望结果:     3,
		},
		{
			被搜寻的文本:   "GoLang编程",
			欲寻找的文本:   "lang",
			起始搜寻位置:   7,
			是否不区分大小写: true,
			期望结果:     2,
		},
		{
			被搜寻的文本:   "",
			欲寻找的文本:   "test",
			起始搜寻位置:   0,
			是否不区分大小写: false,
			期望结果:     -1,
		},
	}

	for _, test := range tests {
		if X倒找文本(test.被搜寻的文本, test.欲寻找的文本, test.起始搜寻位置, test.是否不区分大小写) != test.期望结果 {
			t.Errorf("X倒找文本(%v, %v, %v, %v) 结果错误, 期望结果为 %v", test.被搜寻的文本, test.欲寻找的文本, test.起始搜寻位置, test.是否不区分大小写, test.期望结果)
		}
	}
}

func TestE到大写(t *testing.T) {
	输入 := "hello"
	预期 := "HELLO"
	实际 := X到大写(输入)
	if 实际 != 预期 {
		t.Errorf("对输入%s,预期%s,获得%s", 输入, 预期, 实际)
	}
}

func TestE到小写(t *testing.T) {
	输入 := "HELLO"
	预期 := "hello"
	实际 := X到小写(输入)
	if 实际 != 预期 {
		t.Errorf("对输入%s,预期%s,获得%s", 输入, 预期, 实际)
	}
}

func TestE到全角(t *testing.T) {
	输入 := "hello"
	预期 := "ｈｅｌｌｏ"
	实际 := X到全角(输入)
	if 实际 != 预期 {
		t.Errorf("对输入%s,预期%s,获得%s", 输入, 预期, 实际)
	}
}

func TestE到半角(t *testing.T) {
	输入 := "hello"
	预期 := "hello"
	实际 := X到半角(输入)
	if 实际 != 预期 {
		t.Errorf("对输入%s,预期%s,获得%s", 输入, 预期, 实际)
	}
}

func TestE删首空(t *testing.T) {
	input := "  这是一段带有首部空格的文本。"
	expected := "这是一段带有首部空格的文本。"

	result := X删首空(input)

	if result != expected {
		t.Errorf("X删首空 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE删尾空(t *testing.T) {
	input := "这是一段带有尾部空格的文本。  "
	expected := "这是一段带有尾部空格的文本。"

	result := X删尾空(input)

	if result != expected {
		t.Errorf("X删尾空 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE删首尾空(t *testing.T) {
	input := "  这是一段带有首尾空格的文本。  "
	expected := "这是一段带有首尾空格的文本。"

	result := X删首尾空(input)

	if result != expected {
		t.Errorf("X删首尾空 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE删全部空(t *testing.T) {
	input := "  这 是 一 段 带 有 空 格 的 文 本 。  "
	expected := "这是一段带有空格的文本。"

	result := X删全部空(input)

	if result != expected {
		t.Errorf("X删全部空 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE文本替换(t *testing.T) {
	input := "这是一个测试文本，用来进行文本替换的测试。"
	expected := "这是一个测试替换，用来进行文本替换的测试。"

	result := X文本替换(input, 6, 2, "替换")

	if result != expected {
		t.Errorf("X文本替换 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE子文本替换(t *testing.T) {
	input := "这是一个测试文本，用来进行文本替换的测试。替换替换"
	expected := "这是一个测试文本，用来进行文本修改的测试。修改修改"

	result := X子文本替换(input, "替换", "修改", 0, -1, false)

	if result != expected {
		t.Errorf("X子文本替换 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE取空白文本(t *testing.T) {
	expected := "     "

	result := X取空白文本(5)

	if result != expected {
		t.Errorf("X取空白文本 测试失败，期望输出：%s，实际输出：%s", expected, result)
	}
}

func TestE取重复文本(t *testing.T) {
	input := "测试"
	expected := "测试测试测试"

	result := X取重复文本(3, input)

	if result != expected {
		t.Errorf("X取重复文本 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func TestE分割文本(t *testing.T) {
	input := "这,是,一个,用,逗号分割,的文本。"
	expected := []string{"这", "是", "一个", "用", "逗号分割", "的文本。"}

	result := X分割文本(input, ",")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("X分割文本 测试失败，输入：%s，期望输出：%s，实际输出：%s", input, expected, result)
	}
}

func ExampleE倒找文本() {
	fmt.Println(X倒找文本("这是一个测试文本abc", "测试", 0, false))
	// Output: 4
}

func ExampleE取文本中间() {
	fmt.Println(X取文本中间("这是一个测试文本abc", 4, 2))
	// Output: 测试
}

func TestE文本取随机数字(t *testing.T) {
	t.Log("X文本取随机数字", X文本取随机数字(10))
	t.Log("X文本取随机数字", X文本取随机数字(10))
	t.Log("X文本取随机数字", X文本取随机数字(10))
	t.Log("X文本取随机数字", X文本取随机数字(10))
	t.Log("X文本取随机数字", X文本取随机数字(10))

	t.Log("X文本取随机字母", X文本取随机字母(10, 2))
	t.Log("X文本取随机字母", X文本取随机字母(10, 2))
	t.Log("X文本取随机字母", X文本取随机字母(10, 2))
	t.Log("X文本取随机字母", X文本取随机字母(10, 2))
	t.Log("X文本取随机字母", X文本取随机字母(10, 2))

	t.Log("X文本取随机字母和数字", X文本取随机字母和数字(10))
	t.Log("X文本取随机字母和数字", X文本取随机字母和数字(10))
	t.Log("X文本取随机字母和数字", X文本取随机字母和数字(10))
	t.Log("X文本取随机字母和数字", X文本取随机字母和数字(10))
	t.Log("X文本取随机字母和数字", X文本取随机字母和数字(10))

}

func Test文本_取出中间文本(t *testing.T) {

	t.Log("文本_取左边", X文本取左边("我爱efun好棒", "efun"))
	t.Log("文本_取右边", X文本取右边("我爱efun好棒", "efun"))
	t.Log("文本_取出中间文本", X文本取出中间文本("我爱efun好棒", "我爱", "好棒"))
	t.Log("StrCut", StrCut("我爱efun好棒", "我爱$好棒"))
	t.Log("StrCut", StrCut("我爱efun好棒", "efun$"))
	t.Log("StrCut", StrCut("我爱efun好棒", "$efun"))

}

func Test文本_删除左边(t *testing.T) {
	//一个文本3个字节
	t.Log("文本_删左边", X文本删左边("我爱efun好棒", 2))
	t.Log("文本_删右边", X文本删右边("我爱efun好棒", 2))
	t.Log("文本_删右边", X文本删中间("我爱efun好棒", 2, 4))

	t.Log("文本_取出文本中汉字", X文本取出文本中汉字("我爱efun好棒"))

	for i, s := range X文本逐字分割("我爱efun好棒") {
		t.Log("X文本逐字分割", i, s)
	}
	t.Log("X文本颠倒", X文本颠倒("我爱efun好棒"))
	t.Log("X文本取随机姓氏", X文本取随机姓氏())
	t.Log("X文本取随机姓氏", X文本取随机姓氏())
	t.Log("X文本取随机姓氏", X文本取随机姓氏())
	t.Log("X文本取随机姓氏", X文本取随机姓氏())
	t.Log("X文本取随机姓氏", X文本取随机姓氏())
	t.Log("X文本自动补零", X文本自动补零("1", 5))
	t.Log("X文本自动补零", X文本自动补零("12", 5))
	t.Log("X文本自动补零", X文本自动补零("123", 5))

}

func Test文本_判断大小写(t *testing.T) {

	t.Log("E文本是否为汉字我", X文本是否为汉字("我"))
	t.Log("E文本是否为汉字e", X文本是否为汉字("e"))

	t.Log("X文本是否为字母", X文本是否为字母("我"))
	t.Log("X文本是否为字母.", X文本是否为字母("."))
	t.Log("E文本是否为字母e", X文本是否为字母("e"))

	t.Log("E文本是否为大写字母A", X文本是否为大写字母("A"))
	t.Log("E文本是否为大写字母a", X文本是否为大写字母("X父类"))

	t.Log("E文本是否为小写字母A", X文本是否为小写字母("A"))
	t.Log("E文本是否为小写字母a", X文本是否为小写字母("X父类"))

	t.Log("X文本是否为数字", X文本是否为数字("A"))
	t.Log("E文本是否为数字1", X文本是否为数字("1"))

}
func TestE文本区分_只取字母(t *testing.T) {
	t.Log("X文本区分_只取字母", X文本区分_只取字母("我爱efun好棒"))
	t.Log("X文本区分_只取数字", X文本区分_只取数字("我爱efun520好棒"))
	t.Log("X文本区分_只取汉子", X文本区分_只取汉子("我爱efun520好棒"))
	t.Log("X文本区分_只取符号", X文本区分_只取符号("我爱efun520好棒!@#$%^&*()_+{}|:[]\\;,./"))

	t.Log("X文本首字母改大写", X文本首字母改大写("efun"))
	t.Log("X文本首字母改大写", X文本首字母改大写("我爱efun"))

}
func TestE取文本字数(t *testing.T) {
	t.Log("X取文本字数", X取文本字数("我爱efun好棒"))
	t.Log("X判断文本", X判断文本("我爱efun好棒", "1", "2", "3", "efun"))
	t.Log("X判断文本s", X判断文本s("我爱efun好棒", "1", "2", "3", "efun"))

	t.Log("X判断文本前缀", X判断文本前缀("我爱efun好棒", "我爱"))
	t.Log("X判断文本前缀", X判断文本前缀("我爱efun好棒", "efun"))
	t.Log("X判断文本后缀", X判断文本后缀("我爱efun好棒", "好棒"))
	t.Log("X判断文本后缀", X判断文本后缀("我爱efun好棒", "efun"))

}

func TestE转驼峰(t *testing.T) {
	t.Log("X文本单词首字母大写", X文本单词首字母大写("i love you"))
	t.Log("X文本句子首字母大写", X文本句子首字母大写("i love you"))
	t.Log("X文本自动换行", X文本自动换行("i love you i love you i love you i love you i love you i love you", 5, "\r\n"))
	var p float64
	t.Log("X文本相似文本", X文本相似文本("i love you ", "uoy evol i", &p))
	t.Log("X文本相似文本", p)
	t.Log("X文本相似文本", X文本相似文本("i love you ", "i love you", &p))
	t.Log("X文本相似文本", p)
	t.Log("X文本随机文本", X文本随机文本("i love you "))

	strArr := []string{"a123", "b321", "c123456"}
	t.Log("X文本搜索切片文本", X文本搜索切片文本(strArr, "b321"))

}
