package src

import (
	"container/heap"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//全局的 session 管理器
type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

//用来随需注册存储 session 的结构的 Register 函数的实现
var proviedes = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

//全局唯一的 Session ID
//Session ID 是用来识别访问 Web 应用的每一个用户，因此必须保证它是全局唯一的（GUID），下面代码展示了如何满足这一需求：
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//session 创建
//我们需要为每个来访用户分配或获取与他相关连的 Session，以便后面根据 Session 信息来验证操作。
//SessionStart 这个函数就是用来检测是否已经有某个 Session 与当前来访用户发生了关联，如果没有则创建之。

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/",
			HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//login
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

//操作值：设置、读取和删除
//SessionStart 函数返回的是一个满足 Session 接口的变量，那么我们该如何用他来对 session 数据进行操作呢？
//
//上面的例子中的代码 session.Get("uid") 已经展示了基本的读取数据的操作，现在我们再来看一下详细的操作
//
func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
	//因为 Session 有过期的概念，所以我们定义了 GC 操作，当访问过期时间满足 GC 的触发条件后将会引起 GC，
	//但是当我们进行了任意一个 session 操作，都会对 Session 实体进行更新，都会触发对最后访问时间的修改，
	//这样当 GC 的时候就不会误删除还在使用的 Session 实体。
	//
}

//session重置
//我们知道，Web 应用中有用户退出这个操作，那么当用户退出应用的时候，我们需要对该用户的 session 数据进行销毁操作，
//上面的代码已经演示了如何使用 session 重置操作，下面这个函数就是实现了这个功能：
//
// Destory sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration,
			MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

//session 销毁
func init() {
	go globalSessions.GC()
}

//我们可以看到 GC 充分利用了 time 包中的定时器功能，当超时 maxLifeTime 之后调用 GC 函数，
//这样就可以保证 maxLifeTime 时间内的 session 都是可用的，类似的方案也可以用于统计在线用户数之类的。
//
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.GC()
	})
}

func sessionAndCookie() {
	//	cookie，简而言之就是在本地计算机保存一些用户操作的历史信息（当然包括登录信息），
	//	并在用户再次访问该站点时浏览器通过 HTTP 协议将本地 cookie 内容发送给服务器，
	//	从而完成验证，或继续上一步操作。

	//session，简而言之就是在服务器上保存用户操作的历史信息。
	//服务器使用 session id 来标识 session，session id 由服务器负责产生，
	//保证随机性与唯一性，相当于一个随机密钥，避免在握手或传输中暴露用户真实密码。
	//但该方式下，仍然需要将发送请求的客户端与 session 进行对应，
	//所以可以借助 cookie 机制来获取客户端的标识（即 session id），
	//也可以通过 GET 方式将 id 提交给服务器。

	//Cookie
	//cookie 是有时间限制的，根据生命期不同分成两种：会话 cookie 和持久 cookie；
	//
	//如果不设置过期时间，则表示这个 cookie 的生命周期为从创建到浏览器关闭为止，
	//只要关闭浏览器窗口，cookie 就消失了。这种生命期为浏览会话期的 cookie 被称为会话 cookie。
	//会话 cookie 一般不保存在硬盘上而是保存在内存里。
	//
	//如果设置了过期时间 (setMaxAge (606024))，浏览器就会把 cookie 保存到硬盘上，
	//关闭后再次打开浏览器，这些 cookie 依然有效直到超过设定的过期时间。
	//存储在硬盘上的 cookie 可以在不同的浏览器进程间共享，比如两个 IE 窗口。
	//而对于保存在内存的 cookie，不同的浏览器有不同的处理方式。
	//

	//Go 设置cookie
	// Go语言中通过net/http 包中的SetCookie 来设置
	//http.SetCookie(w http.ResponseWriter, cookie *Cookie)
	//w 表示需要写入的 response, cookie 是一个struct, 让我们来看下cookie 对象是怎么样的
	//type Cookie struct {
	//	Name string
	//	Value string
	//	Path string
	//	Domain string
	//	Expires time.Time
	//	RawExpires string
	//	// MaxAge=0 means no 'Max-Age' attribute specified.
	//	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	//	// MaxAge>0 means Max-Age attribute present and given in seconds
	//	MaxAge int
	//	Secure bool
	//	HttpOnly bool
	//	Raw string
	//	Unparsed []string //Raw text of unparsed attribute-value pairs
	//}

	//设置cookie
	//expiration := time.Now()
	//expiration = expiration.AddDate(1, 0, 0)
	//cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	//http.SetCookie(w, &cookie)

	//Go 读取cookie
	// cookie, _ := r.Cookie("username")
	//fmt.Fprint(w, cookie)

	//另一种读取方式
	//for _, cookie := range r.Cookies() {
	//	fmt.Fprint(w, cookie.Name)
	//}

	//session
	//session，中文经常翻译为会话，其本来的含义是指有始有终的一系列动作 / 消息，
	//比如打电话是从拿起电话拨号到挂断电话这中间的一系列过程可以称之为一个 session。
	//然而当 session 一词与网络协议相关联时，它又往往隐含了 “面向连接” 和 / 或 “保持状态” 这样两个含义。
	//session 在 Web 开发环境下的语义又有了新的扩展，它的含义是指一类用来在客户端与服务器端之间保持状态的解决方案。
	//有时候 Session 也用来指这种解决方案的存储结构。
	//
	//session 机制是一种服务器端的机制，服务器使用一种类似于散列表的结构 (也可能就是使用散列表) 来保存信息。
	//
	//但程序需要为某个客户端的请求创建一个 session 的时候，服务器首先检查这个客户端的请求里是否包含了一个 session 标识－称为 session id，
	//如果已经包含一个 session id 则说明以前已经为此客户创建过 session，服务器就按照 session id 把这个 session 检索出来使用
	//(如果检索不到，可能会新建一个，这种情况可能出现在服务端已经删除了该用户对应的 session 对象，
	//但用户人为地在请求的 URL 后面附加上一个 JSESSION 的参数)。如果客户请求不包含 session id，
	//则为此客户创建一个 session 并且同时生成一个与此 session 相关联的 session id，
	//这个 session id 将在本次响应中返回给客户端保存。
	//
	//session 机制本身并不复杂，然而其实现和配置上的灵活性却使得具体情况复杂多变。
	//这也要求我们不能把仅仅某一次的经验或者某一个浏览器，服务器的经验当作普遍适用的。
	//
	//
	//session 的基本原理是由服务器为每个会话维护一份信息数据，客户端和服务端依靠一个全局唯一的标识来访问这份数据，
	//以达到交互的目的。当用户访问 Web 应用时，服务端程序会随需要创建 session，这个过程可以概括为三个步骤：
	//
	//生成全局唯一标识符（sessionid）；
	//开辟数据存储空间。一般会在内存中创建相应的数据结构，但这种情况下，系统一旦掉电，所有的会话数据就会丢失，
	//如果是电子商务类网站，这将造成严重的后果。所以为了解决这类问题，你可以将会话数据写到文件里或存储在数据库中，
	//当然这样会增加 I/O 开销，但是它可以实现某种程度的 session 持久化，也更有利于 session 的共享；
	//将 session 的全局唯一标示符发送给客户端。
	//
	//以上三个步骤中，最关键的是如何发送这个 session 的唯一标识这一步上。
	//考虑到 HTTP 协议的定义，数据无非可以放到请求行、头域或 Body 里，所以一般来说会有两种常用的方式：cookie 和 URL 重写。
	//
	//
	//Cookie
	//服务端通过设置 Set-cookie 头就可以将 session 的标识符传送到客户端，而客户端此后的每一次请求都会带上这个标识符，
	//另外一般包含 session 信息的 cookie 会将失效时间设置为 0 (会话 cookie)，即浏览器进程有效时间。
	//至于浏览器怎么处理这个 0，每个浏览器都有自己的方案，但差别都不会太大 (一般体现在新建浏览器窗口的时候)；
	//
	//URL 重写
	//所谓 URL 重写，就是在返回给用户的页面里的所有的 URL 后面追加 session 标识符，这样用户在收到响应之后，
	//无论点击响应页面里的哪个链接或提交表单，都会自动带上 session 标识符，从而就实现了会话的保持。
	//虽然这种做法比较麻烦，但是，如果客户端禁用了 cookie 的话，此种方案将会是首选。
	//
	//
	//Go 实现 session管理
	//session 管理设计
	//
	//全局 session 管理器
	//保证 sessionid 的全局唯一性
	//为每个客户关联一个 session
	//session 的存储 (可以存储到内存、文件、数据库等)
	//session 过期处理

	//Session 管理器
	//定义全局的session 管理器

	//type Manager struct {
	//	cookieName  string     // private cookiename
	//	lock        sync.Mutex // protects session
	//	provider    Provider
	//	maxLifeTime int64
	//}
	//
	//func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	//	provider, ok := provides[provideName]
	//	if !ok {
	//		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	//	}
	//	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
	//}

	//Go实现整个的流程应该也是这样的，在 main 包中创建一个全局的 session 管理器

	//var globalSessions *session.Manager
	//然后在init函数中初始化
	//func init() {
	//	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
	//}

	//我们知道 session 是保存在服务器端的数据，它可以以任何的方式存储，比如存储在内存、数据库或者文件中。
	//因此我们抽象出一个 Provider 接口，用以表征 session 管理器底层存储结构。

	//type Provider interface {
	//	SessionInit(sid string) (Session, error)
	//	SessionRead(sid string) (Session, error)
	//	SessionDestroy(sid string) error
	//	SessionGC(maxLifeTime int64)
	//}

	//SessionInit 函数实现 Session 的初始化，操作成功则返回此新的 Session 变量
	//SessionRead 函数返回 sid 所代表的 Session 变量，如果不存在，那么将以 sid 为参数调用 SessionInit 函数创建并返回一个新的 Session 变量
	//SessionDestroy 函数用来销毁 sid 对应的 Session 变量
	//SessionGC 根据 maxLifeTime 来删除过期的数据
	//

	type Session interface {
		Set(key, value interface{}) error //set session value
		Get(key interface{}) interface{}  //get session value
		Delete(key interface{}) error     // delete session value
		SessionID() string
	}

	//以上设计思路来源于 database/sql/driver，先定义好接口，然后具体的存储 session 的结构实现相应的接口并注册后，
	//相应功能这样就可以使用了，以下是用来随需注册存储 session 的结构的 Register 函数的实现。
	//

}
