package src

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

func protect() {
	//session 劫持是一种广泛存在的比较严重的安全威胁，在 session 技术中，
	//客户端和服务端通过 session 的标识符来维护会话， 但这个标识符很容易就能被嗅探到，
	//从而被其他人利用。它是中间人攻击的一种类型。
	//
	//session 劫持防范
	//cookieonly 和 token
	//
	//有效防止session劫持的方法
	// 方案一
	//其中一个解决方案就是 sessionID 的值只允许 cookie 设置，而不是通过 URL 重置方式设置，
	//同时设置 cookie 的 httponly 为 true, 这个属性是设置是否可通过客户端脚本访问这个设置的 cookie，
	//第一这个可以防止这个 cookie 被 XSS 读取从而引起 session 劫持，第二 cookie 设置不会像 URL 重置方式那么容易获取 sessionID。
	//第二步就是在每个请求里面加上 token，实现类似前面章节里面讲的防止 form 重复递交类似的功能，
	//我们在每个请求里面加上一个隐藏的 token，然后每次验证这个 token，从而保证用户的请求都是唯一性。
	//

}

func token() {
	h := md5.New()
	salt := "astaxie%^7&8888"
	io.WriteString(h, salt+time.Now().String())
	token := fmt.Sprintf("%x", h.Sum(nil))
	if r.Form["tone"] != token {
		//提示登录
	}
	sess.Set("token", token)
}

//间隔生成新的SID
//还有一个解决方案就是，我们给 session 额外设置一个创建时间的值，一旦过了一定的时间，我们销毁这个 sessionID，
//重新生成新的 session，这样可以一定程度上防止 session 劫持的问题。
//
func refreshsession() {
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 60) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
}

//session 启动后，我们设置了一个值，用于记录生成 sessionID 的时间。通过判断每次请求是否过期 (这里设置了 60 秒) 定期生成新的 ID，
//这样使得攻击者获取有效 sessionID 的机会大大降低。
//
//上面两个手段的组合可以在实践中消除 session 劫持的风险，一方面，由于 sessionID 频繁改变，
//使攻击者难有机会获取有效的 sessionID；另一方面，因为 sessionID 只能在 cookie 中传递，
//然后设置了 httponly，所以基于 URL 攻击的可能性为零，同时被 XSS 获取 sessionID 也不可能。
//最后，由于我们还设置了 MaxAge=0，这样就相当于 session cookie 不会留在浏览器的历史记录里面。
//

//session 劫持过程
// count计数器

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
