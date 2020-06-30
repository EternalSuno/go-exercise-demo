package src

import (
	"container/heap"
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func beego() {
	//beego orm 是我开发的一个 Go 进行 ORM 操作的库，它采用了 Go style 方式对数据库进行操作，
	//实现了 struct 到数据表记录的映射。beego orm 是一个十分轻量级的 Go ORM 框架，
	//开发这个库的本意降低复杂的 ORM 学习曲线，尽可能在 ORM 的运行效率和功能之间寻求一个平衡，
	//beego orm 是目前开源的 Go ORM 框架中实现比较完整的一个库，而且运行效率相当不错，功能也基本能满足需求。
	//
	//beego orm 是支持 database/sql 标准接口的 ORM 库，所以理论上来说，
	//只要数据库驱动支持 database/sql 接口就可以无缝的接入 beego orm。目前我测试过的驱动包括下面几个
	//Mysql: github/go-mysql-driver/mysql
	//PostgreSQL: github.com/lib/pq
	//SQLite: github.com/mattn/go-sqlite3
	//Mysql: github.com/ziutek/mymysql/godrv
	//暂未支持数据库:
	//MsSql: github.com/denisenkom/go-mssqldb
	//MS ADODB: github.com/mattn/go-adodb
	//Oracle: github.com/mattn/go-oci8
	//ODBC: bitbucket.org/miquella/mgodbc
	//
	//安装
	//go get github.com/astaxie/beego

	//注册驱动
	//orm.RegisterDriver("mysql", orm.DR_MYSQL)
	//设置默认数据库
	//orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
	//注册定义的model
	//orm.RegisterModel(new(User))

	//创建 table
	//orm.RunSyncdb("default", false, true)

	//PostgreSQL配置
	//导入驱动
	// _ "github.com/lib/pq
	//注册驱动
	//orm.RegisterDriver("postgres", orm.DR_Postgres)

	//设置默认数据库
	//PostgresQL 用户: postgres , 密码: zxxx , 数据库名称: test , 数据库别名:default
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=zxxx dbname=test host=127.0.0.1 port-5432 " +
	//	"sslmode=disable")

	//Mysql 配置
	//导入驱动
	// _ "github.com/go-sql-driver/mysql"

	//注册驱动
	//orm.RegisterDriver("mysql", orm.DR_MYSQL)

	//设置默认数据库
	//mysql用户: root 密码 zxxx 数据库:test 数据库别名 default
	//orm.RegisterDataBase("default", "mysql", "root:zxxx@/test?charset=utf8")

	//Sqlite 配置
	//导入驱动
	// _ "github.com/mattn/go-sqlite3"

	//注册驱动
	//orm.RegisterDriver("sqlite", orm.DR_sqlite)

	//设置默认数据库
	//数据库存放位置 ./datas/test.db 数据库别名 default
	//orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")

	//导入必须的 package 之后，我们需要打开到数据库的链接，
	//然后创建一个 beego orm 对象（以 MySQL 为例)，如下所示
	//beego orm:

	//func main() {
	//	o := orm.NewOrm()
	//}

}
