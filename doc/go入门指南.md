
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

也可以使用 defer 语句来记录函数的参数与返回值

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