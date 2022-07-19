/**
* @Author: gxj
* @Data: 2022/7/20-2:10
* @DESC:queue测试函数
**/

package queue

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNew (t *testing.T){
	got := New()
	want := &TSlice{
		I:nil,
		T:nil,
	}
	if got.T != want.T{
		t.Errorf("want:%#v,got:%#v",want,got)
	}
}


func TestTSlice_Push(t *testing.T) {
	var T = &TSlice{
		nil,
		nil,
	}
	var(
		a = 1
		b = 2
		c = "c"
	)
	got := T.Push(a)
	if got != nil{
		t.Errorf("want:%#v,got:%#v",nil,got)
	}
	got = T.Push(b)
	if got != nil{
		t.Errorf("want:%#v,got:%#v",nil,got)
	}
	got = T.Push(c)
	if got != ErrType{
		t.Errorf("want:%#v,got:%#v",nil,got)
	}
}

func TestTSlice_Pop(t *testing.T) {
	var T = &TSlice{
		nil,
		nil,
	}
	var want = []error{
		nil,
		ErrNil,
	}
	_ = T.Push(1)
	for i:=0;i<2;i++{
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := T.Pop()
			if got != want[i]{
				t.Errorf("want:%#v,got:%#v",want,got)
			}
		})
	}

}


func TestTSlice_Top(t *testing.T) {
	var T = &TSlice{
		I: nil,
		T:nil,
	}
	var want = []error{
		nil,
		nil,
		ErrNil,
	}
	T.T = reflect.TypeOf(1)
	T.I = append(T.I,0,1)
	for i:=0;i<3;i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := T.Top()
			if got != want[i]{
				t.Errorf("want:%#v,got:%#v",want,got)
			}
			T.Pop()
		})
	}
}


func TestTSlice_Length(t *testing.T) {
	var T = &TSlice{
		I: nil,
		T:nil,
	}
	var want = []error{
		nil,
		nil,
		ErrNil,
	}
	T.T = reflect.TypeOf(1)
	T.I = append(T.I,0,1)
	for i:=0;i<3;i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := T.Top()
			if got != want[i]{
				t.Errorf("want:%#v,got:%#v",want,got)
			}
			T.Pop()
		})
	}
}