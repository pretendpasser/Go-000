学习笔记

For qusetion:

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。


For answer:

os.Signal
首先要创建一个带缓冲区的管道用来接收信号。
通过signal.Notify()来接收信号并存放到管道中。
另一边通过从管道中取来判断信号。

https://pkg.go.dev/os/signal#example-Notify

https://pkg.go.dev/net/http

https://pkg.go.dev/golang.org/x/sync/errgroup

https://pkg.go.dev/github.com/go-kratos/kratos/pkg/sync/errgroup

https://pkg.go.dev/syscall

无论是http server还是linux signal，都通过context上下文进行控制。以做到一个退出，全部退出。