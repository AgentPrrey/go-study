
# 第五章控制结构

## 测试多返回值函数的错误

## switch结构
选择分支
```
switch var1 {
case val1:
...
case val2:
...
default:
...
}
```
基本结构

##  for 结构
循环控制

## Break 与 continue
break 退出循环
continue 跳过当前循环

## 标签与 goto
跳到标签
 
# 第6章 函数Function
## 6.1介绍
函数 是 基本的代码块
 
函数顺序无关紧要；

但是由于可读性的要求，可以把main()函数写在最前，其他函数按调用顺序往后写

执行特定任务的代码最好只在程序里面出现一次（Don’t Repeat Yourself）

# 6.2 函数参数与返回值
# 6.3 传递变长函数
# 6.4 defer和追踪

关键字defer

允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数

用于释放某些已分配资源的延时释放

当多个defer被注册，会以逆序执行（栈式结构，后进先出）

defer同样可以接受参数（逆序执行）

defer 允许我们在一些函数执行完后进行收尾工作

关闭东西，打印最终报告，解锁加锁资源等；

用defer实现代码追踪，通过打印的方式确定离开某个函数
```
func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }
```
以上是进入和离开的打印信息，显然defer要与leaving匹配

也可以使用 defer 语句来记录函数的参数与返回值，打印出来

https://github.com/AgentPrrey/go-study/blob/master/code/function2/6.12%20defer_logvalues.go

(另一种在调试时使用 defer 语句的手法)

## 6.5 内置函数
不需要进行导入就可以使用的内置函数，有时可以针对不同的类型进行操作

一个简单的列表

https://learnku.com/docs/the-way-to-go/built-in-function/3603

关于内置函数的列表在这里面

## 6.6. 递归函数
一个函数在函数体内调用自身，称为递归  

一般来说，大量的递归调用会导致程序栈内存分配耗尽（栈溢出）

也可以使用相互调用的递归函数：多个函数之间相互调用形成闭环

这些函数的声明顺序可以是任意的

## 6.7. 将函数作为参数


## 关键字
* package 定义包（1）
* import 导入包（2）
* func 函数定义（3）
* const 定义常量（4）
* var 定义变量 （5）
* switch 选择结构（6）
* if 选择结构 （7）
* else 选择结构 （8）
* default 选择结构的默认选项（9）
* case 选择结构标签 （10）
* fallthrough case带有fallthrough 强制执行下一条 不判断条件（11）
* for 循环 （12）
* break 跳出循环（13）
* continue 跳过本次循环（14）
* goto 跳转语句（15）
* defer 延时执行 栈结构（16）
* range 取字符串的a[i] 赋值给v打印 可以处理特殊字节（多byte）（17）
* 
# 备注
https://learnku.com/docs/the-way-to-go