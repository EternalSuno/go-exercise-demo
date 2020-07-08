package src

import "regexp"

func filter() {
	//大多数 Web 应用的漏洞都是因为没有对用户输入的数据进行恰当过滤所引起的。
	//
	//我们介绍的过滤数据分成三个步骤：
	//
	//1、识别数据，搞清楚需要过滤的数据来自于哪里
	//2、过滤数据，弄明白我们需要什么样的数据
	//3、区分已过滤及被污染数据，如果存在攻击数据那么保证过滤之后可以让我们使用更安全的数据
	//

	//识别数据
	//“识别数据” 作为第一步是因为在你不知道 “数据是什么，它来自于哪里” 的前提下，你也就不能正确地过滤它。
	//这里的数据是指所有源自非代码内部提供的数据。例如：所有来自客户端的数据，但客户端并不是唯一的外部数据源，
	//数据库和第三方提供的接口数据等也可以是外部数据源。
	//
	//由用户输入的数据我们通过 Go 非常容易识别，Go 通过 r.ParseForm 之后，把用户 POST 和 GET 的数据全部放在了 r.Form 里面。
	//其它的输入要难识别得多，例如，r.Header 中的很多元素是由客户端所操纵的。
	//常常很难确认其中的哪些元素组成了输入，所以，最好的方法是把里面所有的数据都看成是用户输入。
	//(例如 r.Header.Get("Accept-Charset") 这样的也看做是用户输入，虽然这些大多数是浏览器操纵的)
	//

	//数据过滤
	//过滤数据主要采用如下一些库来操作：
	//
	//strconv 包下面的字符串转化相关函数，因为从 Request 中的 r.Form 返回的是字符串，
	//而有些时候我们需要将之转化成整 / 浮点数，Atoi、ParseBool、ParseFloat、ParseInt 等函数就可以派上用场了。
	//string 包下面的一些过滤函数 Trim、ToLower、ToTitle 等函数，能够帮助我们按照指定的格式获取信息。
	//regexp 包用来处理一些复杂的需求，例如判定输入是否是 Email、生日之类。

	//过滤数据除了检查验证之外，在特殊时候，还可以采用白名单。即假定你正在检查的数据都是非法的，除非能证明它是合法的。
	//
	//使用这个方法，如果出现错误，只会导致把合法的数据当成是非法的，而不会是相反，尽管我们不想犯任何错误，
	//但这样总比把非法数据当成合法数据要安全得多。
	//

	//
	//区分过滤数据
	//在编写 Web 应用的时候我们还需要区分已过滤和被污染数据，因为这样可以保证过滤数据的完整性，而不影响输入的数据。
	//我们约定把所有经过过滤的数据放入一个叫全局的 Map 变量中 (CleanMap)。这时需要用两个重要的步骤来防止被污染数据的注入：
	//
	//每个请求都要初始化 CleanMap 为一个空 Map。
	//加入检查及阻止来自外部数据源的变量命名为 CleanMap。

	//<form action="/whoami" method="POST">
	//    我是谁:
	//    <select name="name">
	//        <option value="astaxie">astaxie</option>
	//        <option value="herry">herry</option>
	//        <option value="marry">marry</option>
	//    </select>
	//    <input type="submit" />
	//</form>

	//在处理这个表单的编程逻辑中，非常容易犯的错误是认为只能提交三个选择中的一个。其实攻击者可以模拟 POST 操作，
	//递交 name=attack 这样的数据，所以在此时我们需要做类似白名单的处理

	// r.ParseForm()
	//name := r.Form.Get("name")
	//CleanMap := make(map[string]interface{}, 0)
	//if name == "astaxie" || name == "herry" || name == "marry" {
	//	CleanMap["name"] = name
	//}
	//上面代码中我们初始化了一个 CleanMap 的变量，当判断获取的 name 是 astaxie、herry、marry 三个中的一个之后
	//，我们把数据存储到了 CleanMap 之中，这样就可以确保 CleanMap ["name"] 中的数据是合法的，从而在代码的其它部分使用它。
	//当然我们还可以在 else 部分增加非法数据的处理，一种可能是再次显示表单并提示错误。但是不要试图为了友好而输出被污染的数据。
	//
	//上面的方法对于过滤一组已知的合法值的数据很有效，但是对于过滤有一组已知合法字符组成的数据时就没有什么帮助。
	//例如，你可能需要一个用户名只能由字母及数字组成：
	//

	r.ParseForm()
	username := r.Form.Get("username")
	CleanMap := make(map[string]interface{}, 0)
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username); ok {
		CleanMap["username"] = username
	}

}
