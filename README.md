该仓库包含了2022暑期红岩网校的课后作业

# 架构设计与编码规范

**作业：**

类似[网站](https://github.com/trekhleb/state-of-the-art-shitcode/blob/master/README.zh-CN.md)中那样，以bad and good 的形式修改期末考核代码。

**提交：**

通过Markdown代码块形式提交作业。

代码块格式提交代码与注释相关

一些重复性高的修改只提交一份。

源代码会修改更新

<<<<<<< HEAD
关于注释：由于为作业代码，所以注释在使用英文的同时会保留中文注释（英文也是用机翻，没中文看不懂自己写了啥）

# Cron

**作业：**

在golang中编写一个**cron框架**，并且有该框架的**使用示例**。该框架应有如下功能：

-  新建一个cron引擎实例 

-  在cron引擎中，添加/删除一个定时任务 

-  启动/停止/重置cron引擎 

启动引擎后就为每个任务开启一个goroutine。直到任务结束。

问题：

- 无法解析cron格式语法

- 没有解决随时停止重置cron引擎，启动后无法关闭

关于注释：由于为作业代码，所以注释会保留中文注释（没中文看不懂自己写了啥）

修改ing

# 接口反射的实现及使用

**作业：**

有一个接口切片：

```go
type TSlice struct{
    i []interface{}
    //你需要的其他字段
    type	type	//添加字段type，当队列存入第一个元素时，通过typeof确定该队列的元素类型。
}

t:=TSlice{}
t.push()
t.pop()
t.top()
t.length()
```

需要实现：

- 实现队列，先入先出

- push方法（入一个），pop方法（出一个），top方法（返回队头元素），length方法（返回长度）

- 切片的元素实例可以是任意类型，但同一个TSlice实例的元素只能是一个类型

单测代码覆盖率为83.3%

# Apache,Nginx,Kubernetes

### 作业

### lv0

分别使用apache和nginx部署一个html页面，内容随意。交配置文件中关键内容（截图或者复制文本）

### lv0.5

使用nginx监听9999端口，反向代理到8081，8082，8083端口。使用任意策略的负载均衡。交配置文件中关键内容（截图或者复制文本）

### lv1

使用k8s部署自己的上次考核的作业。部署好之后交如下指令的结果（截图或者复制文本）

```
kubectl describe <你使用的workloads类型> <name>
kubectl describe service <name>
```

























