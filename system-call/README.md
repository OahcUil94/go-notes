# Go系统调用

- [linux系统调用参考](https://github.com/torvalds/linux/blob/master/arch/x86/entry/syscalls/syscall_64.tbl)

## 获取pid

例如`39	common	getpid			sys_getpid`, 标号39表示用来获取pid, 代码如下:

```go
package main

import (
	"fmt"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println(pid)
}
```

## 模拟strace命令

### 介绍strace命令

- strace命令是Linux进程照妖镜, strace是个功能强大的Linux调试分析诊断工具, 可用于跟踪程序执行时进程系统调用(system call)和所接收的信号, 尤其是针对源码不可读或源码无法再编译的程序.
- 在Linux系统中, 用户进程不能直接访问计算机硬件设备, 当进程需要访问硬件设备(如读取磁盘文件或接收网络数据等)时, 必须由用户态模式切换至内核态模式, 通过系统调用访问硬件设备.
- strace可跟踪进程产生的系统调用, 包括参数、返回值和执行所消耗的时间.
- 若strace没有任何输出, 并不代表此时进程发生阻塞, 也可能程序进程正在执行某些不需要与系统其它部分发生通信的事情.
- strace从内核接收信息, 且无需以任何特殊方式来构建内核.

编辑`print.go`文件:
 
```go
package main

func main() {
    println("Hello World!!!")
}
```

执行`go build print.go`构建二进制文件, 执行`strace -c ./print`: 

```
[vagrant@localhost code]$ strace -c ./print
Hello World!!!
% time     seconds  usecs/call     calls    errors syscall
------ ----------- ----------- --------- --------- ----------------
 55.60    0.000149           1       114           rt_sigaction
 19.03    0.000051          51         1           futex
 11.19    0.000030          10         3           clone
 10.07    0.000027           3         8           rt_sigprocmask
  2.24    0.000006           6         1           write
  1.49    0.000004           2         2           sigaltstack
  0.37    0.000001           1         1           gettid
  0.00    0.000000           0        18           mmap
  0.00    0.000000           0         1           execve
  0.00    0.000000           0         1           uname
  0.00    0.000000           0         1           arch_prctl
  0.00    0.000000           0         1           sched_getaffinity
  0.00    0.000000           0         1         1 openat
------ ----------- ----------- --------- --------- ----------------
100.00    0.000268                   153         1 total
```

直接执行`strace ./print`, 输出:

```
rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
clone(child_stack=0xc000042000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM) = 3387
rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
clone(child_stack=0xc000044000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM) = 3388
rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
--- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=3386, si_uid=1000} ---
rt_sigreturn({mask=[]})                 = 0
--- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=3386, si_uid=1000} ---
rt_sigreturn({mask=[]})                 = 0
futex(0xc000032548, FUTEX_WAKE_PRIVATE, 1) = 1
futex(0xc000032548, FUTEX_WAKE_PRIVATE, 1) = 1
rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
clone(child_stack=0xc00003e000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM) = 3389
rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f9d013a2000
write(2, "Hello World!!!\n", 15Hello World!!!
)        = 15
exit_group(0)                           = ?
+++ exited with 0 +++
``` 

### ptrace系统调用

`ptrace`可以拦截系统调用并且修改该系统调用的参数, 提供了一种机制使得父进程可以观察和控制子进程的执行过程, 
还可以检查和修改该子进程的可执行文件在内存中的镜像及该子进程所使用的寄存器中的值, 主要用于实现对进程插入断点和跟踪子进程的系统调用.

源码`syscall_linux.go`可以查看到封装的很多`ptrace`函数.

## 参考资料

- [Go 语言中的系统调用](https://zhuanlan.zhihu.com/p/58285124)
- [Go是如何进行系统调用的](https://www.zhihu.com/question/264073701/answer/627197588)
- [系统调用真正的效率瓶颈在哪里](https://www.zhihu.com/question/32043825)
- [Linux系统调用过程分析](https://zhuanlan.zhihu.com/p/79236207)