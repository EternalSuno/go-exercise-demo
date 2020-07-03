package sockett

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func socket() {
	//Socket 起源于 Unix，而 Unix 基本哲学之一就是 “一切皆文件”，
	//都可以用 “打开 open –> 读写 write/read –> 关闭 close” 模式来操作。
	//Socket 就是该模式的一个实现，网络的 Socket 数据传输是一种特殊的 I/O，Socket 也是一种文件描述符。
	//Socket 也具有一个类似于打开文件的函数调用：Socket ()，该函数返回一个整型的 Socket 描述符，
	//随后的连接建立、数据传输等操作都是通过该 Socket 实现的。
	//
	//常用的 Socket 类型有两种：流式 Socket（SOCK_STREAM）和数据报式 Socket（SOCK_DGRAM）。
	//流式是一种面向连接的 Socket，针对于面向连接的 TCP 服务应用；
	//数据报式 Socket 是一种无连接的 Socket，对应于无连接的 UDP 服务应用。
	//
	//Socket 如何通信
	//
	//网络中的进程之间如何通过 Socket 通信呢？
	//首要解决的问题是如何唯一标识一个进程，否则通信无从谈起！
	//在本地可以通过进程 PID 来唯一标识一个进程，但是在网络中这是行不通的。
	//其实 TCP/IP 协议族已经帮我们解决了这个问题，网络层的 “ip 地址” 可以唯一标识网络中的主机，
	//而传输层的 “协议 + 端口” 可以唯一标识主机中的应用程序（进程）。
	//这样利用三元组（ip 地址，协议，端口）就可以标识网络的进程了，网络中需要互相通信的进程，就可以利用这个标志在他们之间进行交互。
	//
	//
	//使用 TCP/IP 协议的应用程序通常采用应用编程接口：
	//UNIX BSD 的套接字（socket）和 UNIX System V 的 TLI（已经被淘汰），
	//来实现网络进程之间的通信。就目前而言，几乎所有的应用程序都是采用 socket，而现在又是网络时代，网络中进程通信是无处不在，
	//这就是为什么说 “一切皆 Socket”。
	//

	//Socket 基础
	//通过上面的介绍我们知道 Socket 有两种：
	//TCP Socket 和 UDP Socket，TCP 和 UDP 是协议，而要确定一个进程的需要三元组，需要 IP 地址和端口。
	//IPv4 地址
	//目前的全球因特网所采用的协议族是 TCP/IP 协议。IP 是 TCP/IP 协议中网络层的协议，是 TCP/IP 协议族的核心协议。
	//目前主要采用的 IP 协议的版本号是 4 (简称为 IPv4)，发展至今已经使用了 30 多年。
	//
	//IPv4 的地址位数为 32 位，也就是最多有 2 的 32 次方的网络设备可以联到 Internet 上。
	//近十年来由于互联网的蓬勃发展，IP 位址的需求量愈来愈大，使得 IP 位址的发放愈趋紧张，前一段时间，据报道 IPV4 的地址已经发放完毕，
	//

	//IPv6 地址
	//IPv6 是下一版本的互联网协议，也可以说是下一代互联网的协议，它是为了解决 IPv4 在实施过程中遇到的各种问题而被提出的，
	//IPv6 采用 128 位地址长度，几乎可以不受限制地提供地址。按保守方法估算 IPv6 实际可分配的地址，
	//整个地球的每平方米面积上仍可分配 1000 多个地址。在 IPv6 的设计过程中除了一劳永逸地解决了地址短缺问题以外，
	//还考虑了在 IPv4 中解决不好的其它问题，主要有端到端 IP 连接、服务质量（QoS）、安全性、多播、移动性、即插即用等。
	//地址格式类似这样：2002:c0e8:82e7:0:0:0:c0e8:82e7
	//
	//

	//Go 支持的 IP 类型
	//在 Go 的 net 包中定义了很多类型、函数和方法用来网络编程，其中 IP 的定义如下：
	//type IP []byte
	//在 net 包中有很多函数来操作 IP，但是其中比较有用的也就几个，
	//其中 ParseIP(s string) IP 函数会把一个 IPv4 或者 IPv6 的地址转化成 IP 类型

	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)

	//TCP Socket
	//在 Go 语言的 net 包中有一个类型 TCPConn，这个类型可以用来作为客户端和服务器端交互的通道，他有两个主要的函数
	//func (c *TCPConn) Write(b []byte) (int, error)
	//func (c *TCPConn) Read(b []byte) (int, error)

	//TCPConn 可以用在客户端和服务器端来读写数据。
	// TCPAddr 类型 标识一个TCP地址信息
	//type TCPAddr struct {
	//	IP IP
	//	Port int
	//	Zone string //IPv6 scoped addressing zone
	//}

	//在 Go 语言中通过 ResolveTCPAddr 获取一个 TCPAddr
	//func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
	//net 参数是 "tcp4"、"tcp6"、"tcp" 中的任意一个，分别表示 TCP (IPv4-only), TCP (IPv6-only) 或者
	//TCP (IPv4, IPv6 的任意一个)。
	//addr 表示域名或者 IP 地址，例如 "www.google.com:80" 或者 "127.0.0.1:22"。
	//

	//"HEAD / HTTP/1.0\r\n\r\n"
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)

	//首先程序将用户的输入作为参数 service 传入 net.ResolveTCPAddr 获取一个 tcpAddr,
	//然后把 tcpAddr 传入 DialTCP 后创建了一个 TCP 连接 conn，通过 conn 来发送请求信息，
	//最后通过 ioutil.ReadAll 从 conn 中读取全部的文本，也就是服务端响应反馈的信息。
	//

	//TCP server
	//也可以通过 net 包来创建一个服务器端程序，在服务器端我们需要绑定服务到指定的非激活端口，并监听此端口，
	//当有客户端请求到达的时候可以接收到来自客户端连接的请求。net 包中有相应功能的函数，函数定义如下
	//
	//func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
	//func (l *TCPListener) Accept() (Conn, error)
	//参数说明同 DialTCP 的参数一样。下面我们实现一个简单的时间同步服务，监听 7777 端口
	//service := ":7777"
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//checkError(err)
	//listener, err := net.Listen("tcp", tcpAddr)
	//checkError(err)
	//for {
	//	conn,err := listener.Accept()
	//	if err != nil {
	//		continue
	//	}
	//	daytime := time.Now().String()
	//	conn.Write([]byte(daytime))
	//	conn.Close()
	//}
	//上面的服务跑起来之后，它将会一直在那里等待，直到有新的客户端请求到达。
	//当有新的客户端请求到达并同意接受 Accept 该请求的时候他会反馈当前的时间信息。
	//值得注意的是，在代码中 for 循环里，当有错误发生时，直接 continue 而不是退出，
	//是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，然后当前连接的客户端直接报错而退出，
	//从而不会影响到当前服务端运行的整个服务。
	//
	//上面的代码有个缺点，执行的时候是单任务的，不能同时接收多个请求，那么该如何改造以使它支持多并发呢？
	//Go 里面有一个 goroutine 机制，请看下面改造后的代码

	//

	//这个服务端没有处理客户端实际请求的内容。如果我们需要通过从客户端发送不同的请求来获取不同的时间格式，而且需要一个长连接
	//service := "1200"
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//checkError(err)
	//listener, err := net.ListenTCP("tcp", tcpAddr)
	//checkError(err)
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		continue
	//	}
	//	go handleClient(conn)
	//}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 *time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128)
	}
}
//在上面这个例子中，我们使用 conn.Read() 不断读取客户端发来的请求。
//由于我们需要保持与客户端的长连接，所以不能在读取完一次请求后就关闭连接。
//由于 conn.SetReadDeadline() 设置了超时，当一定时间内客户端无请求发送，
//conn 便会自动关闭，下面的 for 循环即会因为连接已关闭而跳出。
//需要注意的是，request 在创建时需要指定一个最大长度以防止 flood attack；每次读取到请求处理完毕后，
//需要清理 request，因为 conn.Read() 会将新读取到的内容 append 到原内容之后。
//


//func handleClient(conn net.Conn) {
//	defer conn.Close()
//	daytime := time.Now().String()
//	conn.Write([]byte(daytime))
//}
//通过把业务处理分离到函数 handleClient，我们就可以进一步地实现多并发执行了。
//看上去是不是很帅，增加 go 关键词就实现了服务端的多并发，从这个小例子也可以看出 goroutine 的强大之处。

//控制TCP连接
// TCP 有很多连接控制函数，我们平常用到比较多的有如下几个函数：
//
//func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)
//设置建立连接的超时时间，客户端和服务器端都适用，当超过设置时间时，连接自动关闭。
//
//func (c *TCPConn) SetReadDeadline(t time.Time) error
//func (c *TCPConn) SetWriteDeadline(t time.Time) error
//用来设置 写入 / 读取 一个连接的超时时间。当超过设置时间时，连接自动关闭。
//
//func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
//设置 keepAlive 属性，是操作系统层在 tcp 上没有数据和 ACK 的时候，会间隔性的发送 keepalive 包，
//操作系统可以通过该包来判断一个 tcp 连接是否已经断开，在 windows 上默认 2 个小时没有收到数据和
//keepalive 包的时候认为 tcp 连接已经断开，这个功能和我们通常在应用层加的心跳包的功能类似。
//

//UDP Socket
//处理 UDP Socket 和 TCP Socket 不同的地方就是在服务器端处理多个客户端请求数据包的方式不同，
//UDP 缺少了对客户端连接请求的 Accept 函数。其他基本几乎一模一样，只有 TCP 换成了 UDP 而已。
//UDP 几个主要的函数如下
// func ResolveUDPAddr (net, addr string) (*UDPAddr, os.Error)
// func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPCoon, err os.Error)
// func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
// func (c *UPDConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, error os.Error)
// func (c *UPDConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
//

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
