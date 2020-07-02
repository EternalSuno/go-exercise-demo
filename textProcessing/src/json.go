package src

import (
	"encoding/json"
	"fmt"
)

//生成JSON
//我们开发很多应用的时候，最后都是要输出 JSON 数据串，那么如何来处理呢？JSON 包里面通过 Marshal 函数来处理，函数定义如下：
//func Marsha1(v interface{}) ([]byte, error)
type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//输出内容
	//
	//{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}
	//
	//我们看到上面的输出字段名的首字母都是大写的，如果你想用小写的首字母怎么办呢？把结构体的字段名改成首字母小写的？
	//JSON 输出的时候必须注意，只有导出的字段才会被输出，
	//如果修改字段名，那么就会发现什么都不会输出，所以必须通过 struct tag 定义来实现：
	//
	//
	type Server struct {
		ServerName string `json:"serverName"`
		ServerIP   string `json:"serverIP"`
	}

	type Serverslice struct {
		Servers []Server `json:"servers"`
	}
	//通过修改上面的结构体定义，输出的 JSON 串就和我们最开始定义的 JSON 串保持一致了。
	//针对 JSON 的输出，我们在定义 struct tag 的时候需要注意的几点是:
	//字段的 tag 是 "-"，那么这个字段不会输出到 JSON
	//tag 中带有自定义名称，那么这个自定义名称会出现在 JSON 的字段名中，例如上面例子中 serverName
	//tag 中如果带有 "omitempty" 选项，那么如果该字段值为空，就不会输出到 JSON 串中
	//如果字段类型是 bool, string, int, int64 等，而 tag 中带有 ",string" 选项，
	//那么这个字段在输出到 JSON 的时候会把该字段对应的值转换成 JSON 字符串
	//

	//type Server struct {
	//	ID int `json:"-"` //ID 不会导出到JSON中
	//
	//	// ServerName2 的只会进行二次 JSON 编码
	//	ServerName string `json:"serverName"`
	//	ServerName2 string `json:"serverName2,string"`
	//
	//	//如果 ServerIP 为空 , 则不输出到JSON 串中
	//	ServerIP string `json:"serverIP,omitempty"`
	//}
	//
	//s := Server {
	//	ID: 3,
	//	ServerName: `Go "1.0"`,
	//	ServerName2: `Go "1.0"`,
	//	ServerIP: ``,
	//}
	//
	//b, _ := json.Marshal(s)
	//os.Stdout.Write(b)
	//输出内容
	//{"serverName":"Go \"1.0\" ", "serverName2": "\"Go \\\"1.0\\\" \""}
	//Marshal 函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
	//JSON 对象只支持 string 作为 key，所以要编码一个 map，那么必须是 map [string] T 这种类型 (T 是 Go 语言中任意的类型)
	//Channel, complex 和 function 是不能被编码成 JSON 的
	//嵌套的数据是不能编码的，不然会让 JSON 编码进入死循环
	//指针在编码的时候会输出指针指向的内容，而空指针会输出 null
	//
}

//type Server struct {
//	ServerName string
//	ServerIP   string
//}
//
//type Serverslice struct {
//	Servers []Server
//}
//
//func main() {
//	var s Serverslice
//	str := `{{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
//				{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}}`
//	json.Unmarshal([]byte(str), &s)
//	fmt.Println(s)

//在上面的示例代码中，我们首先定义了与 json 数据对应的结构体，数组对应 slice，
//字段名对应 JSON 里面的 KEY，在解析的时候，如何将 json 数据与 struct 字段相匹配呢？
//例如 JSON 的 key 是 Foo，那么怎么找对应的字段呢？
//
//首先查找 tag 含有 Foo 的可导出的 struct 字段 (首字母大写)
//其次查找字段名是 Foo 的导出字段
//最后查找类似 FOO 或者 FoO 这样的除了首字母之外其他大小写不敏感的导出字段
//
//聪明的你一定注意到了这一点：能够被赋值的字段必须是可导出字段 (即首字母大写）。同时 JSON 解析的时候只会解析能找得到的字段，
//找不到的字段会被忽略，这样的一个好处是：当你接收到一个很大的 JSON 数据结构而你却只想获取其中的部分数据的时候，
//你只需将你想要的数据对应的字段名大写，即可轻松解决这个问题。
//
//解析到interface
//我们知道 interface {} 可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的 json 数据的结果。
//JSON 包中采用 map [string] interface {} 和 [] interface {} 结构来存储任意的 JSON 对象和数组。
//Go 类型和 JSON 类型的对应关系如下：
//
// bool 代表JSON booleans
// float64 代表 JSON numbers
// string 代表JSON strings
// nil 代表 JSON null
//
//假设有如下JSON数据
// b := []byte(`{"Name": "Wednesday", "Age":6, "Parents":["Gomez", "Morticia"]}`)
// 如果在我们不知道他的结构的情况下，我们把他解析到 interface {} 里面
//var f interface{}
//err := json.Unmarshal(b, &f)
//这个时候 f 里面存储了一个 map 类型，他们的 key 是 string，值存储在空的 interface {} 里
//f = map[string]interface{}{
//	"Name": "Wednesday",
//	"Age": 6,
//	"Parents": []interface{}{
//		"Gomez",
//		"Morticia",
//	},
//}
// 通过断言的方式来访问这些数据
//m := f.(map[string]interface{})
//断言之后,你就可以访问里面的数据了
//for k, v :=range m {
//	switch vv := v.(type) {
//	case string:
//		fmt.Println(k, "is string", vv)
//	case int:
//		fmt.Println(k, "is int", vv)
//	case float64:
//		fmt.Println(k, "is float64", vv)
//	case []interface{}:
//		fmt.Println(k, "is an array:")
//		for i, u := range vv {
//			fmt.Println(i, u)
//		}
//	default:
//		fmt.Println(k, "is of a type I don't know how to handle")
//	}
//}
//通过上面的示例可以看到，通过 interface {} 与 type assert 的配合，我们就可以解析未知结构的 JSON 数了。
//
//上面这个是官方提供的解决方案，其实很多时候我们通过类型断言，操作起来不是很方便，目前 bitly 公司开源了一个叫做 simplejson 的包，
//在处理未知结构体的 JSON 时相当方便，详细例子如下所示：
//
//js, err := simplejson.NewJson([]byte(`{
//	"test": {
//		"array": [1, "2", 3],
//		"int": 10,
//		"float": 5.150,
//		"bignum": 9223372036854775807,
//		"string": "simplejson",
//		"bool": true
//	}
//}`))
//
//arr, _ := js.Get("test").Get("array").Array()
//i, _ := js.Get("test").Get("int").Int()
//ms := js.Get("test").Get("string").MustString()

//}

func jsonTest() {
	//JSON（Javascript Object Notation）是一种轻量级的数据交换语言，以文字为基础，具有自我描述性且易于让人阅读。
	//尽管 JSON 是 Javascript 的一个子集，但 JSON 是独立于语言的文本格式，并且采用了类似于 C 语言家族的一些习惯。
	//JSON 与 XML 最大的不同在于 XML 是一个完整的标记语言，而 JSON 不是。
	//JSON 由于比 XML 更小、更快，更易解析，以及浏览器的内建快速解析支持，使得其更适用于网络数据传输领域。
	//目前我们看到很多的开放平台，基本上都是采用了 JSON 作为他们的数据交互的接口。
	//

	/*
		{
		    "servers": [
		        {
		            "serverName": "Shanghai_VPN",
		            "serverIP": "127.0.0.1"
		        },
		        {
		            "serverName": "Beijing_VPN",
		            "serverIP": "127.0.0.2"
		        }
		    ]
		}
	*/
	//解析JSON
	// 解析到结构体
	// Go 的 JSON 包中有如下函数
	//
	//func Unmarshal(data []byte, v interface{}) error

}
