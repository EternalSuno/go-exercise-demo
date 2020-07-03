package rpc

import (
	"errors"
	"fmt"
	"net/http"
)

func rpc()  {
	//RPC 就是想实现函数调用模式的网络化。客户端就像调用本地函数一样，然后客户端把这些参数打包之后通过网络传递到服务端，
	//服务端解包到处理过程中执行，然后执行的结果反馈给客户端。
	//
	//RPC（Remote Procedure Call Protocol）—— 远程过程调用协议，是一种通过网络从远程计算机程序上请求服务，
	//而不需要了解底层网络技术的协议。它假定某些传输协议的存在，如 TCP 或 UDP，以便为通信程序之间携带信息数据。
	//通过它可以使函数调用模式网络化。在 OSI 网络通信模型中，RPC 跨越了传输层和应用层。
	//RPC 使得开发包括网络分布式多程序在内的应用程序更加容易。
	//
	//运行时，一次客户机对服务器的 RPC 调用，其内部操作大致有如下十步：
	//
	//1. 调用客户端句柄；执行传送参数
	//2. 调用本地系统内核发送网络消息
	//3. 消息传送到远程主机
	//4. 服务器句柄得到消息并取得参数
	//5. 执行远程过程
	//6. 执行的过程将结果返回服务器句柄
	//7. 服务器句柄返回结果，调用远程系统内核
	//8. 消息传回本地主机
	//9. 客户句柄由内核接收消息
	//10. 客户接收句柄返回的数据
	//
	//
	//
	//Go RPC
	//Go 标准包中已经提供了对 RPC 的支持，而且支持三个级别的 RPC：TCP、HTTP、JSONRPC。
	//但 Go 的 RPC 包是独一无二的 RPC，它和传统的 RPC 系统不同，它只支持 Go 开发的服务器与客户端之间的交互，因为在内部，
	//它们采用了 Gob 来编码。
	//Go RPC 的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
	//
	//函数必须是导出的 (首字母大写)
	//必须有两个导出类型的参数，
	//第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
	//函数还要有一个返回值 error
	//
	//func (t *T) MethodName(argType T1, replyType *T2) error
	//
	//任何的 RPC 都需要通过网络来传递数据，Go RPC 可以利用 HTTP 和 TCP 来传递数据，
	//利用 HTTP 的好处是可以直接复用 net/http 里面的一些函数。
	// HTTP RPC
	//
}
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}


