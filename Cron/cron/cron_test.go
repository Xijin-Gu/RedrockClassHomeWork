/**
* @Author: gxj
* @Data: 2022/7/19-3:42
* @DESC:对cron测试函数
**/

package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestInit(t *testing.T) {
	got := Init()
	want := &Cron{
		Task: nil,
		NUM: 0,
	}
	assert.Equal(t, got,want)


}