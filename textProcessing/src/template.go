package src

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

//Go 模板使用
//  Go 语言中，我们使用 template 包来进行模板处理，使用类似 Parse、ParseFile、Execute 等方法从文件或者字符串加载模板，
// 然后执行类似上面图片展示的模板的 merge 操作。
//
func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")
	t, _ = t.ParseFiles("tmpl/welcome.html")
	user := GetUser()  //获取当前用户信息
	t.Execute(w, user) //执行模板的merger操作
}

//为了演示和测试代码的方便，我们在接下来的例子中采用如下格式的代码
//
//使用 Parse 代替 ParseFiles，因为 Parse 可以直接测试一个字符串，而不需要额外的文件
//不使用 handler 来写演示代码，而是每个测试一个 main，方便测试
//使用 os.Stdout 代替 http.ResponseWriter，因为 os.Stdout 实现了 io.Writer 接口
//

// 字段操作
// Go 语言的模板通过 {{}} 来包含需要在渲染时被替换的字段，{{.}} 表示当前的对象，
//这和 Java 或者 C++ 中的 this 类似，如果要访问当前对象的字段通过 {{.FieldName}}，
//但是需要注意一点：这个字段必须是导出的 (字段首字母必须是大写的)，否则在渲染的时候就会报错
//
//type Person struct {
//	UserName string
//}
//
//func tempp() {
//	t := template.New("fieldname example")
//	t, _ = t.Parse("hello {{.UserName}}")
//	p := Person{UserName: "Astaxie"}
//	t.Execute(os.Stdout, p)
//}

//上面的代码我们可以正确的输出 hello Astaxie，但是如果我们稍微修改一下代码，在模板中含有了未导出的字段，那么就会报错
//type Person struct {
//	UserName string
//	email    string //未到处的字段， 首字母是小写的
//}

//t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
//上面的代码就会报错，因为我们调用了一个未导出的字段，但是如果我们调用了一个不存在的字段是不会报错的，而是输出为空。
//
//如果模板中输出 {{.}}，这个一般应用于字符串对象，默认会调用 fmt 包输出字符串的内容。
//

//输出嵌套字段内容
//上面我们例子展示了如何针对一个对象的字段输出，那么如果字段里面还有对象，如何来循环的输出这些内容呢？
//我们可以使用 {{with …}}…{{end}} 和 {{range …}}{{end}} 来进行数据的输出。
//
//{{range}} 这个和 Go 语法里面的 range 类似，循环操作数据
//{{with}} 操作是指当前对象的值，类似上下文的概念

type Friend struct {
	Fname string
}

type Persion struct {
	UserName string
	Emails   []string
	Friend   []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	//find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	//replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func maintest() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Email}}
				an emails {{.|emailDeal}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
	`)

	p := Persion{
		UserName: "Astaxie",
		Emails:   []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}

	t.Execute(os.Stdout, p)
}

//其实，在模板包内部已经有内置的实现函数，下面代码截取自模板包里面
//
//var builtins = FuncMap{
//	"and":      and,
//	"call":     call,
//	"html":     HTMLEscaper,
//	"index":    index,
//	"js":       JSEscaper,
//	"len":      length,
//	"not":      not,
//	"or":       or,
//	"print":    fmt.Sprint,
//	"printf":   fmt.Sprintf,
//	"println":  fmt.Sprintln,
//	"urlquery": URLQueryEscaper,
//}

//Must 操作
//模板包里面有一个函数 Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写。
//接下来我们演示一个例子，用 Must 来判断模板是否正确：
func musttest() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }}"))
}

//输出
//
//The first one parsed OK.
//The second one parsed OK.
//The next one ought to fail.
//panic: template: check parse error with Must:1: unexpected "}" in command
//模板包里面有一个函数 Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，
//变量是否正确的书写。接下来我们演示一个例子，用 Must 来判断模板是否正确：

//嵌套模板
//我们平常开发 Web 应用的时候，经常会遇到一些模板有些部分是固定不变的，然后可以抽取出来作为一个独立的部分，
//例如一个博客的头部和尾部是不变的，而唯一改变的是中间的内容部分。所以我们可以定义成 header、content、footer 三个部分。
//Go 语言中通过如下的语法来申明
// {{define "子模板名称"}}内容{{end}}
// {{template "子模板名称"}}

func main() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

//通过上面的例子我们可以看到通过 template.ParseFiles 把所有的嵌套模板全部解析到模板里面，
//其实每一个定义的 {{define}} 都是一个独立的模板，他们相互独立，是并行存在的关系，
//内部其实存储的是类似 map 的一种关系 (key 是模板的名称，value 是模板的内容)，
//然后我们通过 ExecuteTemplate 来执行相应的子模板内容，我们可以看到 header、footer 都是相对独立的，
//都能输出内容，content 中因为嵌套了 header 和 footer 的内容，就会同时输出三个的内容。
//但是当我们执行 s1.Execute，没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西。
//

func temps() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func temp() {
	//条件处理
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipepline if demo: {{if `anything`}} 内容 输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else 部分." +
		"{{end}}\n"))

	tIfElse.Execute(os.Stdout, nil)
	//注意：if 里面无法使用条件判断，例如 .Mail=="astaxie@gmail.com"，这样的判断是不正确的，if 里面只能是 bool 值

	//pipelines
	//Unix 用户已经很熟悉什么是 pipe 了，ls | grep "beego" 类似这样的语法你是不是经常使用，
	//过滤当前目录下面的文件，显示含有 "beego" 的数据，表达的意思就是前面的输出可以当做后面的输入，
	//最后显示我们想要的数据，而 Go 语言模板最强大的一点就是支持 pipe 数据，在 Go 语言里面任何 {{}} 里面的都是 pipelines 数据，
	//例如我们上面输出的 email 里面如果还有一些可能引起 XSS 注入的，那么我们如何来进行转化呢？
	// {{. | html}}
	//在 email 输出的地方我们可以采用如上方式可以把输出全部转化 html 的实体，
	//上面的这种方式和我们平常写 Unix 的方式是不是一模一样，操作起来相当的简便，调用其他的函数也是类似的方式。

	//模板变量
	//with``range``if 过程中申明局部变量，这个变量的作用域是 {{end}} 之前，Go 语言通过申明的局部变量格式如下所示
	//	$variable := pipeline
	//

	//{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
	//{{with $x := "output"}}{{printf "%q" $x}}{{end}}
	//{{with $x := "output"}}{{$x | printf "%q"}}{{end}}

	//模板函数
	//模板在输出对象的字段值时，采用了 fmt 包把对象转化成了字符串。
	//但是有时候我们的需求可能不是这样的，例如有时候我们为了防止垃圾邮件发送者通过采集网页的方式来发送给我们的邮箱信息，
	//我们希望把 @ 替换成 at 例如：astaxie at beego.me，如果要实现这样的功能，我们就需要自定义函数来做这个功能。
	//
	//
	//type FuncMap map[string]interface{}
	//如果我们想要的 email 函数的模板函数名是 emailDeal，它关联的 Go 函数名称是 EmailDealWith,
	//那么我们可以通过下面的方式来注册这个函数
	//t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	//
	//func EmailDealWith(args …interface{}) string
}
