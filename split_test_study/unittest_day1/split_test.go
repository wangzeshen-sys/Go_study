package split

import (
	"reflect"
	"testing"
)

func TestSplit1(t *testing.T) { // 测试函数必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":") // 程序输出的结构
	want := []string{"a", "b", "c"} // 期望结果
	if !reflect.DeepEqual(want, got) { // 因为slice不饿能直接比较，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v\n", want, got) // 参数失败输出错误提示
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got){
		t.Errorf("excepted:%v, got:%v\n", want, got)
	}
}