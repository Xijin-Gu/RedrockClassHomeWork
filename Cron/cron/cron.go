/**
* @Author: gxj
* @Data: 2022/7/19-2:19
* @DESC:一个cron框架，支持新建一个cron引擎，启动/重置/停止该引擎，并添加/删除任务
**/

package cron

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Task 定义task结构体
type Task struct {
	ID  int			//任务ID
	NextTime	time.Duration //下次任务与这次任务的间隔
	Task		func()		//任务执行内容
}

// Cron 定义Cron结构体
type Cron struct {

	Task []Task		//该引擎中存在的任务内容
	NUM int			//当前引擎中存在的任务数目
}

var (
	wg sync.WaitGroup
	errNilTask  = errors.New("task is nil")
)

//Init 初始化一个cron实例
func Init() *Cron{
	cron := &Cron{
		Task: nil,
		NUM: 0,
	}
	fmt.Println("cron is be created\n...")
	return cron
}

//Run 启动cron引擎
func (cron *Cron) Run() error {
	if cron.Task == nil {
		return errNilTask
	}
	for _,v := range cron.Task{
		wg.Add(1)
		go v.Do()
	}
	wg.Wait()
	return nil
}

//Stop 停止一个cron引擎
func (cron *Cron) Stop(){
	cron = &Cron{nil,0}

}

//Do 启动一个任务
func (task *Task) Do() {
	for i:=0;i<1; {
		time.Sleep(task.NextTime)
		task.Task()
		wg.Add(1)
		wg.Done()
	}
}

//Add 为cron引擎增加一个任务
func (cron *Cron) Add(second int,function func()) func(){
	var du time.Duration
	for i:=0;i<second;i++{
		du +=time.Second
	}
	cron.Task = append(cron.Task,Task{NextTime: du,Task: function})
	for i,v := range cron.Task{
		v.ID = i
	}
	return function
}

//Delete 为cron删除一个任务
func (cron *Cron) Delete(id int) {
	cron.Task = append(cron.Task[:id],cron.Task[id+1:]...)
}






