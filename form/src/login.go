package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //// 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	// 注意:如果没有调用 ParseForm 方法, 下面无法获取表单数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客服端
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		// 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
	//通过上面的代码可以看到，处理文件上传我们需要调用 r.ParseMultipartForm，里面的参数表示 maxMemory，
	//调用 ParseMultipartForm 之后，上传的文件存储在 maxMemory 大小的内存里面，如果文件大小超过了 maxMemory，
	//那么剩下的部分将存储在系统的临时文件中。我们可以通过 r.FormFile 获取上面的文件句柄，然后实例中使用了 io.Copy 来存储文件。
	//
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求方法
	if r.Method == "GET" {

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}

		//token 防止多次提交
		//token := r.Form.Get("token")
		//if token != "" {
		//	//验证token的合法性
		//} else {
		//	// 不存在token报错
		//}
		//表单验证
		//必填字段
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("username 为空")
		}

		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
			fmt.Println("请输入正确的真实姓名")
		}

		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("请输入正确年龄格式")
		}
		if getint > 100 {
			fmt.Println("年龄太大了")
		}

		//正则匹配
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			fmt.Println("正则匹配年龄失败")
		}

		//英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
			fmt.Println("英文格式错误")
		}

		//电子邮件
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2, 4})$`, r.Form.Get("email")); !m {
			fmt.Println("邮箱格式错误")
		} else {
			fmt.Println("邮箱格式正确")
		}

		//手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4, 8})$`, r.Form.Get("mobile")); !m {
			fmt.Println("手机号格式错误")
		}

		//下拉菜单
		slice := []string{"apple", "pear", "banana"}

		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
				fmt.Println("选项正确")
			} else {
				fmt.Println("错误种类")
			}
		}

		//单选按钮
		gender := []string{"1", "2"}
		for _, v := range gender {
			if v == r.Form.Get("gender") {
				fmt.Println("单选正确")
			} else {
				fmt.Println("单选错误")
			}
		}

		//复选框
		//对于复选框我们的验证和单选有点不一样，因为接收到的数据是一个 slice
		//interest := []string{"football", "basketball", "tennis"}
		//a := Slice_diff(r.Form["interest"], interest)
		//if a == nil {
		//	fmt.Println("选项正确")
		//}else{
		//	fmt.Println("选项不正确")
		//}
		//
		//上面这个函数 Slice_diff 包含在我开源的一个库里面 (操作 slice 和 map 的库)，github.com/astaxie/beeku

		//时间和日期
		//Go 里面提供了一个 time 的处理包，我们可以把用户的输入年月日转化成相应的时间，然后进行逻辑判断

		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
		fmt.Printf("Go launched at %s\n", t.Local())

		//身份证号
		// 验证15位身份证 , 15位全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
			fmt.Println("15位身份证号码验证失败")
		}

		//验证18位身份证, 18位前17位为数字, 最后一位是校验位, 可能为数字或字符X.
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
			fmt.Println("18位身份证号码验证失败")
		}

		//防止多次提交表单
		//解决方案是在表单中添加一个带有唯一值的隐藏字段。在验证表单时，先检查带有该唯一值的表单是否已经递交过了。
		//如果是，拒绝再次递交；如果不是，则处理表单进行逻辑处理。另外，如果是采用了 Ajax 模式递交表单的话，
		//当表单递交后，通过 javascript 来禁用表单的递交按钮
		//
		//application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
		//multipart/form-data   不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
		//text/plain   空格转换为 "+" 加号，但不对特殊字符编码。
		//

		// 请求的是登录数据，那么执行登录的逻辑判断
		//fmt.Println("username", r.Form["username"])
		//fmt.Println("password", r.Form["password"])
		//fmt.Println(reflect.TypeOf(r.Form))
		//fmt.Println(reflect.TypeOf(r.Form["password"]))
		//fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		//fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		//template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客服端
		//v := url.Values{}
		//v.Set("name", "Ava")
		//v.Add("friend", "Jess")
		//v.Add("friend", "Sarah")
		//fmt.Println(v.Get("name"))
		//fmt.Println(v.Get("friend"))
		//fmt.Println(v["friend"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置路由
	http.HandleFunc("/login", login)         //设置路由
	http.HandleFunc("/upload", upload)       //设置路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//通过上面的代码我们可以看出获取请求方法是通过 r.Method 来完成的，这是个字符串类型的变量，返回 GET, POST, PUT 等 method 信息。
	//
	//login 函数中我们根据 r.Method 来判断是显示登录界面还是处理登录逻辑。当 GET 方式请求时显示登录界面，其他方式请求时则处理登录逻辑，
	//如查询数据库、验证登录信息等。
	//
	//request.Form 是一个 url.Values 类型，里面存储的是对应的类似 key=value 的信息，
	//下面展示了可以对 form 数据进行的一些操作:
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	fmt.Println(v.Get("name"))
	//Ava
	fmt.Println(v.Get("friend"))
	//Jess
	fmt.Println(v["friend"])
	//[Jess Sarah]

	//Request 本身也提供了 FormValue () 函数来获取用户提交的参数。
	//如 r.Form ["username"] 也可写成 r.FormValue ("username")。
	//调用 r.FormValue 时会自动调用 r.ParseForm，所以不必提前调用。
	//r.FormValue 只会返回同名参数中的第一个，若参数不存在则返回空字符串
	//

	//预防跨站脚本
	// Go 的 html/template 里面带有下面几个函数可以帮你转义
	//
	//func HTMLEscape (w io.Writer, b [] byte) // 把 b 进行转义之后写到 w
	//func HTMLEscapeString (s string) string // 转义 s 之后返回结果字符串
	//func HTMLEscaper (args ...interface {}) string // 支持多个参数一起转义，返回结果字符串
	//
	//fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
	//fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
	//template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
	//Go 的 html/template 包默认帮你过滤了 html 标签，
	//但是有时候你只想要输出这个 <script>alert()</script> 看起来正常的信息，该怎么处理？
	//请使用 text/template。请看下面的例子
	//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(out, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	//输出
	// Hello, <script>alert('you have been pwed')</script>

	//表单验证
	//必填字段
	//if len(r.Form["username"][0]) == 0 {
	// 为空处理
	//}
	//r.Form 对不同类型的表单元素的留空有不同的处理， 对于空文本框、空文本区域以及文件上传，
	//元素的值为空值，而如果是未选中的复选框和单选按钮，则根本不会在 r.Form 中产生相应条目，
	//如果我们用上面例子中的方式去获取数据时程序就会报错。所以我们需要通过 r.Form.Get() 来获取值，
	//因为如果字段不存在，通过该方式获取的是空值。但是通过 r.Form.Get() 只能获取单个的值，
	//如果是 map 的值，必须通过上面的方式来获取。
	//
	//数字
	// 你想要确保一个表单输入框中获取的只能是数字，
	//例如，你想通过表单获取某个人的具体年龄是 50 岁还是 10 岁，
	//而不是像 “一把年纪了” 或 “年轻着呢” 这种描述
	//
	//如果我们是判断正整数，那么我们先转化成 int 类型，然后进行处理
	//
	//getint,err:=strconv.Atoi(r.Form.Get("age"))
	//if err!=nil{
	//	// 数字转化出错了，那么可能就不是数字
	//}
	//
	//// 接下来就可以判断这个数字的大小范围了
	//if getint >100 {
	//	// 太大了
	//}
	//
	//正则
	//
	//if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
	//	return false
	//}

	//对于性能要求很高的用户来说，这是一个老生常谈的问题了，他们认为应该尽量避免使用正则表达式，因为使用正则表达式的速度会比较慢。
	//但是在目前机器性能那么强劲的情况下，对于这种简单的正则表达式效率和类型转换函数是没有什么差别的。
	//如果你对正则表达式很熟悉，而且你在其它语言中也在使用它，那么在 Go 里面使用正则表达式将是一个便利的方式。
	//Go 实现的正则是 RE2，所有的字符都是 UTF-8 编码的。

	//中文
	//有时候我们想通过表单元素获取一个用户的中文名字，但是又为了保证获取的是正确的中文，
	//我们需要进行验证，而不是用户随便的一些输入。对于中文我们目前有两种方式来验证，
	//可以使用 unicode 包提供的 func Is(rangeTab *RangeTable, r rune) bool 来验证，
	//也可以使用正则方式来验证，这里使用最简单的正则方式，如下代码所示
	//if m, _ := regexp.MatchString("\\p{Han}+$", r.Form.Get("realname")); !m{
	//	return false
	//}

}
