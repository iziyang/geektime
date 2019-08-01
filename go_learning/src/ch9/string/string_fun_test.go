package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	var s = "A,B,C"
	var parts = strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestCon(t *testing.T) {
	var s = strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
