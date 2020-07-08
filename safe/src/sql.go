package src

func sqltest() {
	//SQL 注入攻击（SQL Injection），简称注入攻击，是 Web 开发中最常见的一种安全漏洞。可以用它来从数据库获取敏感信息，
	//或者利用数据库的特性执行添加用户，导出文件等一系列恶意操作，甚至有可能获取数据库乃至系统用户最高权限。
	//
	//而造成 SQL 注入的原因是因为程序没有有效过滤用户的输入，使攻击者成功的向服务器提交恶意的 SQL 查询代码，
	//程序在接收后错误的将攻击者的输入作为查询语句的一部分执行，导致原始的查询逻辑被改变，额外的执行了攻击者精心构造的恶意代码。
	//
	//SQL 注入实例
	//
	//
	//表单
	//<form action="/login" method="POST">
	//<p>Username: <input type="text" name="username" /></p>
	//<p>Password: <input type="password" name="password" /></p>
	//<p><input type="submit" value="登陆" /></p>
	//</form>

	//服务端
	//
	//username:=r.Form.Get("username")
	//password:=r.Form.Get("password")
	//sql:="SELECT * FROM user WHERE username='"+username+"' AND password='"+password+"'"

	// 输入
	//myuser' or 'foo' = 'foo' --

	//SQL变成
	//SELECT * FROM user WHERE username='myuser' or 'foo' = 'foo' --'' AND password='xxx'
	//在 SQL 里面 -- 是注释标记，所以查询语句会在此中断。这就让攻击者在不知道任何合法用户名和密码的情况下成功登录了。

	//对于 MSSQL 还有更加危险的一种 SQL 注入，就是控制系统，下面这个可怕的例子将演示如何在某些版本的 MSSQL 数据库上执行系统命令。
	//
	//sql:="SELECT * FROM products WHERE name LIKE '%"+prod+"%'"
	//Db.Exec(sql)
	//如果攻击提交 a%' exec master..xp_cmdshell 'net user test testpass /ADD' -- 作为变量 prod 的值，那么 sql 将会变成
	//
	//sql:="SELECT * FROM products WHERE name LIKE '%a%' exec master..xp_cmdshell 'net user test testpass /ADD'--%'"

	//MSSQL 服务器会执行这条 SQL 语句，包括它后面那个用于向系统添加新用户的命令。
	//如果这个程序是以 sa 运行而 MSSQLSERVER 服务又有足够的权限的话，攻击者就可以获得一个系统帐号来访问主机了
	//

	//如何预防 SQL 注入
	//也许你会说攻击者要知道数据库结构的信息才能实施 SQL 注入攻击。确实如此，但没人能保证攻击者一定拿不到这些信息，一旦他们拿到了，数据库就存在泄露的危险。如果你在用开放源代码的软件包来访问数据库，比如论坛程序，攻击者就很容易得到相关的代码。如果这些代码设计不良的话，风险就更大了。目前 Discuz、phpwind、phpcms 等这些流行的开源程序都有被 SQL 注入攻击的先例。
	//
	//这些攻击总是发生在安全性不高的代码上。所以，永远不要信任外界输入的数据，特别是来自于用户的数据，包括选择框、表单隐藏域和 cookie。就如上面的第一个例子那样，就算是正常的查询也有可能造成灾难。
	//
	//SQL 注入攻击的危害这么大，那么该如何来防治呢？下面这些建议或许对防治 SQL 注入有一定的帮助。
	//
	//严格限制 Web 应用的数据库的操作权限，给此用户提供仅仅能够满足其工作的最低权限，从而最大限度的减少注入攻击对数据库的危害。
	//检查输入的数据是否具有所期望的数据格式，严格限制变量的类型，例如使用 regexp 包进行一些匹配处理，
	//或者使用 strconv 包对字符串转化成其他基本类型的数据进行判断。

	//对进入数据库的特殊字符（'"\ 尖括号 &*; 等）进行转义处理，或编码转换。Go 的 text/template 包里面的 HTMLEscapeString 函数可以对字符串进行转义处理。
	//所有的查询语句建议使用数据库提供的参数化查询接口，参数化的语句使用参数而不是将用户输入变量嵌入到 SQL 语句中，即不要直接拼接 SQL 语句。例如使用 database/sql 里面的查询函数 Prepare 和 Query，或者 Exec(query string, args ...interface{})。
	//在应用发布之前建议使用专业的 SQL 注入检测工具进行检测，以及时修补被发现的 SQL 注入漏洞。网上有很多这方面的开源工具，例如 sqlmap、SQLninja 等。
	//避免网站打印出 SQL 错误信息，比如类型错误、字段不匹配等，把代码里的 SQL 语句暴露出来，以防止攻击者利用这些错误信息进行 SQL 注入。
	//
	//

}
