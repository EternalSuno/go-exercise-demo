package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		fmt.Println(reflect.TypeOf(r.Form))
		fmt.Println(reflect.TypeOf(r.Form["password"]))
		v := url.Values{}
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend"))
		fmt.Println(v["friend"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置路由
	http.HandleFunc("/login", login)         //设置路由
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
}
