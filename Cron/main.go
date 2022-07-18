/**
* @Author: gxj
* @Data: 2022/7/19-2:41
* @DESC:测试用
**/

package main

import (
	"fmt"
	cron2 "redrock/RedrockClassWork/Cron/cron"
)

func main(){
	//创建一个cron引擎
	cron := cron2.Init()

	//增加一个任务
	fu := cron.Add(1, func() {
		fmt.Println("hello world")
	})
	fu()

	 cron.Delete(0)

	//启动一个cron引擎,任务为空则启动失败
	err := cron.Run()
	if err != nil {
		fmt.Println(err)
	}
	//停止引擎
	cron.Stop()
}

