/**
* @Author: gxj
* @Data: 2022/7/20-1:19
* @DESC:基本队列的实现
**/

package queue

import (
	"errors"
	"fmt"
	"reflect"
)

//TSlice 队列载体
type TSlice struct{
	I 		[]interface{}
	T 	 	reflect.Type
}

var (
	ErrType = errors.New("数据类型错误")
	ErrNil = errors.New("队列为空")
)

//New 创建一个新的队列
func New() *TSlice {
	fmt.Println("创建了一个新队列")
	var T = &TSlice{
		I: nil,
		T:nil,
	}
	return T
}
//Push 添加一个新元素
func (T *TSlice) Push (v interface{})error{
	if T.T == nil {
		T.T = reflect.TypeOf(v)
	}
	if T.T != reflect.TypeOf(v){
			return ErrType
	}
	T.I = append(T.I,v)
	fmt.Println(T)
	return nil
}

//Pop 移除一个元素，由先进先出可知，要移除队首元素
func (T *TSlice) Pop ()error{
	if len(T.I) == 0{
		return ErrNil
	}
	d := T.I[0]
	T.I = T.I[1:len(T.I)]
	fmt.Println(d,"已经被删除")
	return nil
}
//Top 查看队头元素
func (T *TSlice) Top() error{
	if len(T.I) == 0{
		return ErrNil
	}
	fmt.Println(T.I[0])
	return nil
}

//Length 查看队列的长度
func (T *TSlice) Length()error{
	if len(T.I) == 0{
		return ErrNil
	}
	fmt.Println("length:",len(T.I))
	return nil
}