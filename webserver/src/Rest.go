package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func rest() {
	//	REST (REpresentational State Transfer) 这个概念，
	//	首次出现是在 2000 年 Roy Thomas Fielding（他 是 HTTP 规范的主要编写者之一）的博士论文中，
	//	它指的是一组架构约束条件和原则。满足这些约束条件和原则的应用程序或设计就是 RESTful 的。

	//  资源（Resources）
	//REST 是 "表现层状态转化"，其实它省略了主语。"表现层" 其实指的是 "资源" 的 "表现层"。
	//
	//那么什么是资源呢？就是我们平常上网访问的一张图片、一个文档、一个视频等。这些资源我们通过 URI 来定位，也就是一个 URI 表示一个资源。
	//
	//表现层（Representation）
	//
	//资源是做一个具体的实体信息，他可以有多种的展现方式。而把实体展现出来就是表现层，例如一个 txt 文本信息，
	//他可以输出成 html、json、xml 等格式，一个图片他可以 jpg、png 等方式展现，这个就是表现层的意思。
	//
	//URI 确定一个资源，但是如何确定它的具体表现形式呢？应该在 HTTP 请求的头信息中用 Accept 和 Content-Type 字段指定，
	//这两个字段才是对 "表现层" 的描述。
	//
	//状态转化（State Transfer）
	//
	//访问一个网站，就代表了客户端和服务器的一个互动过程。在这个过程中，肯定涉及到数据和状态的变化。
	//而 HTTP 协议是无状态的，那么这些状态肯定保存在服务器端，所以如果客户端想要通知服务器端改变数据和状态的变化，肯定要通过某种方式来通知它。
	//
	//客户端能通知服务器端的手段，只能是 HTTP 协议。具体来说，就是 HTTP 协议里面，
	//四个表示操作方式的动词：GET、POST、PUT、DELETE。
	//它们分别对应四种基本操作：GET 用来获取资源，POST 用来新建资源（也可以用于更新资源），PUT 用来更新资源，DELETE 用来删除资源。
	//
	//综合上面的解释，我们总结一下什么是 RESTful 架构：
	//
	//（1）每一个 URI 代表一种资源；
	//（2）客户端和服务器之间，传递这种资源的某种表现层；
	//（3）客户端通过四个 HTTP 动词，对服务器端资源进行操作，实现 "表现层状态转化"。
	//
	//Web 应用要满足 REST 最重要的原则是：客户端和服务器之间的交互在请求之间是无状态的，
	//即从客户端到服务器的每个请求都必须包含理解请求所必需的信息。如果服务器在请求之间的任何时间点重启，
	//客户端不会得到通知。此外此请求可以由任何可用服务器回答，这十分适合云计算之类的环境。
	//因为是无状态的，所以客户端可以缓存数据以改进性能。
	//
	//当 REST 架构的约束条件作为一个整体应用时，将生成一个可以扩展到大量客户端的应用程序。
	//它还降低了客户端和服务器之间的交互延迟。统一界面简化了整个系统架构，改进了子系统之间交互的可见性。
	//REST 简化了客户端和服务器的实现，而且对于使用 REST 开发的应用程序更加容易扩展。
	//
	// RESTful 的实现
	//Go 没有为 REST 提供直接支持，但是因为 RESTful 是基于 HTTP 协议实现的，
	//所以我们可以利用 net/http 包来自己实现，当然需要针对 REST 做一些改造，
	//REST 是根据不同的 method 来处理相应的资源，目前已经存在的很多自称是 REST 的应用，
	//其实并没有真正的实现 REST，我暂且把这些应用根据实现的 method 分成几个级别，请看下图：
	//
	//RESTful 服务充分利用每一个 HTTP 方法，包括 DELETE 和 PUT。可有时，HTTP 客户端只能发出 GET 和 POST 请求：
	//
	//HTML 标准只能通过链接和表单支持 GET 和 POST。在没有 Ajax 支持的网页浏览器中不能发出 PUT 或 DELETE 命令
	//
	//有些防火墙会挡住 HTTP PUT 和 DELETE 请求，要绕过这个限制，客户端需要把实际的 PUT 和 DELETE 请求通过 POST 请求穿透过来。
	//RESTful 服务则要负责在收到的 POST 请求中找到原始的 HTTP 方法并还原。
	//
	//

}

//Go 实现RESTful
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "hello, %s!\n", ps.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprint(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprint(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprint(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func try(){
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)

	log.Fatal(http.ListenAndServe(":8080", router))

	//上面的代码演示了如何编写一个 REST 的应用，我们访问的资源是用户，我们通过不同的 method 来访问不同的函数，
	//这里使用了第三方库 github.com/julienschmidt/httprouter，在前面章节我们介绍过如何实现自定义的路由器，
	//这个库实现了自定义路由和方便的路由规则映射，通过它，我们可以很方便的实现 REST 的架构。
	//通过上面的代码可知，REST 就是根据不同的 method 访问同一个资源的时候实现不同的逻辑处理。

}


