package src

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"strconv"
)

func csrf() {
	//CSRF（Cross-site request forgery），中文名称：跨站请求伪造，
	//也被称为：one click attack/session riding，缩写为：CSRF/XSRF。
	//
	//要完成一次 CSRF 攻击，受害者必须依次完成两个步骤 ：
	//
	//1. 登录受信任网站 A，并在本地生成 Cookie 。
	//2. 在不退出 A 的情况下，访问危险网站 B。
	//看到这里，读者也许会问：“如果我不满足以上两个条件中的任意一个，就不会受到 CSRF 的攻击”。是的，确实如此，但你不能保证以下情况不会发生：
	//
	//你不能保证你登录了一个网站后，不再打开一个 tab 页面并访问另外的网站，特别现在浏览器都是支持多 tab 的。
	//你不能保证你关闭浏览器了后，你本地的 Cookie 立刻过期，你上次的会话已经结束。
	//上图中所谓的攻击网站，可能是一个存在其他漏洞的可信任的经常被人访问的网站。
	//
	//因此对于用户来说很难避免在登陆一个网站之后不点击一些链接进行其他操作，所以随时可能成为 CSRF 的受害者。
	//
	//CSRF 攻击主要是因为 Web 的隐式身份验证机制，Web 的身份验证机制虽然可以保证一个请求是来自于某个用户的浏览器，
	//但却无法保证该请求是用户批准发送的。
	//
	//如何预防 CSRF
	//CSRF 的防御可以从服务端和客户端两方面着手，防御效果是从服务端着手效果比较好，现在一般的 CSRF 防御也都在服务端进行。
	//
	//服务端的预防 CSRF 攻击的方式方法有多种，但思想上都是差不多的，主要从以下 2 个方面入手：
	//
	//1. 正确使用 GET , POST 和 Cookie；
	//2. 在非 GET 请求中增加伪随机数；
	//
	//一般而言，普通的 Web 应用都是以 GET、POST 为主，还有一种请求是 Cookie 方式。我们一般都是按照如下方式设计应用：
	//
	//1.GET 常用在查看，列举，展示等不需要改变资源属性的时候；
	//2.POST 常用在下达订单，改变一个资源的属性或者做其他一些事情；
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	//这样处理后，因为我们限定了修改只能使用 POST，当 GET 方式请求时就拒绝响应，
	//所以上面图示中 GET 方式的 CSRF 攻击就可以防止了，
	//但这样就能全部解决问题了吗？当然不是，因为 POST 也是可以模拟的。
	//
	//因此我们需要实施第二步，在非 GET 方式的请求中增加随机数，这个大概有三种方式来进行：
	//
	//为每个用户生成一个唯一的 cookie token，所有表单都包含同一个伪随机值，这种方案最简单，
	//因为攻击者不能获得第三方的 Cookie (理论上)，所以表单中的数据也就构造失败，
	//但是由于用户的 Cookie 很容易由于网站的 XSS 漏洞而被盗取，所以这个方案必须要在没有 XSS 的情况下才安全。

	//每个请求使用验证码，这个方案是完美的，因为要多次输入验证码，所以用户友好性很差，所以不适合实际运用。
	//不同的表单包含一个不同的伪随机值，我们在 4.4 小节介绍 “如何防止表单多次递交” 时介绍过此方案，复用相关代码，实现如下：
	//
	//生成随机数token
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	io.WriteString(h, "ganxxx")
	token := fmt.Sprintf("%x", h.Sum(nil))
	t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, token)

	//输出token
	// <input type="hidden" name="token" value="{{.}}">

	//验证token
	r.ParseForm()
	token := r.Form.Get("token")
	if token != "" {
		// 验证合法性
	} else {
		// token 不存在
	}
	//这样基本就实现了安全的 POST，但是也许你会说如果破解了 token 的算法呢，按照理论上是，
	//但是实际上破解是基本不可能的，因为有人曾计算过，暴力破解该串大概需要 2 的 11 次方时间。

}