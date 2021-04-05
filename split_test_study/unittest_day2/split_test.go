package split

import (
	"reflect"
	"testing"
)

func TestSplit2(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片

	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"e","b","c"}},
		{input: "a:b:c", sep: ",", want: []string{"b:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"s", "d"}},
		{input: "枯藤老树昏鸦", sep: "老", want: []string{"藤", "树昏鸦"}},
	}
	// 遍历切片， 桌椅执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%v, got:%v\n", tc.want, got)
		}
	}
}
