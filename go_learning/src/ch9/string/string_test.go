package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认的零值

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xBA\xFF"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))
	c := []rune(s) //提取字符串中的 unicode 值
	t.Log(c)
	t.Logf("%T", c)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}
