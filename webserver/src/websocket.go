package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//WebSocket 采用了一些特殊的报头，使得浏览器和服务器只需要做一个握手的动作，就可以在浏览器和服务器之间建立一条连接通道。
	//且此连接会保持在活动状态，你可以使用 JavaScript 来向连接写入或从中接收数据，就像在使用一个常规的 TCP Socket 一样。
	//它解决了 Web 实时化的问题，相比传统 HTTP 有如下好处：
	//一个 Web 客户端只建立一个 TCP 连接
	//Websocket 服务端可以推送 (push) 数据到 web 客户端.
	//有更加轻量级的头，减少数据传送量
	//WebSocket URL 的起始输入是
	//ws:// 或是 wss://（在 SSL 上）

	//websocket 原理
	// WebSocket 的协议颇为简单，在第一次 handshake 通过以后，连接便建立成功，
	//其后的通讯数据都是以 ”\x00″ 开头，以 ”\xFF” 结尾。在客户端，
	//这个是透明的，WebSocket 组件会自动将原始数据 “掐头去尾”。
	//
	//浏览器发出 WebSocket 连接请求，然后服务器发出回应，然后连接建立成功，这个过程通常称为 “握手” (handshaking)。

	//Go 实现 WebSocket
	//Go 语言标准包里面没有提供对 WebSocket 的支持，但是在由官方维护的 go.net 子包中有对这个的支持，你可以通过如下的命令获取该包：
	//go get golang.org/x/net/websocket
	//WebSocket 分为客户端和服务端，接下来我们将实现一个简单的例子：用户输入信息，
	//客户端通过 WebSocket 将信息发送给服务器端，服务器端收到信息之后主动 Push 信息到客户端，
	//然后客户端将输出其收到的信息，客户端的代码如下：
	//
	//
	//<html>
	//<head></head>
	//<body>
	//<script type="text/javascript">
	//var sock = null;
	//var wsuri = "ws://127.0.0.1:1234";
	//
	//window.onload = function() {
	//
	//	console.log("onload");
	//
	//	sock = new WebSocket(wsuri);
	//
	//	sock.onopen = function() {
	//		console.log("connected to " + wsuri);
	//	}
	//
	//	sock.onclose = function(e) {
	//		console.log("connection closed (" + e.code + ")");
	//	}
	//
	//	sock.onmessage = function(e) {
	//		console.log("message received: " + e.data);
	//	}
	//};
	//
	//function send() {
	//	var msg = document.getElementById('message').value;
	//	sock.send(msg);
	//};
	//</script>
	//<h1>WebSocket Echo Test</h1>
	//<form>
	//<p>
	//	Message: <input id="message" type="text" value="Hello, world!">
	//</p>
	//</form>
	//<button onclick="send();">Send Message</button>
	//</body>
	//</html>

	//可以看到客户端 JS，很容易的就通过 WebSocket 函数建立了一个与服务器的连接 sock，当握手成功后，
	//会触发 WebScoket 对象的 onopen 事件，告诉客户端连接已经成功建立。客户端一共绑定了四个事件。
	//
	//1）onopen 建立连接后触发
	//2）onmessage 收到消息后触发
	//3）onerror 发生错误时触发
	//4）onclose 关闭连接时触发
	//

	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant't send")
			break
		}

	}
}
