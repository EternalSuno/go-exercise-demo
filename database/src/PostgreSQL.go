package src

import (
	"database/sql"
	"fmt"
)

func curl() {
	//PostgreSQL 是一个自由的对象 - 关系数据库服务器 (数据库管理系统)，它在灵活的 BSD - 风格许可证下发行。
	//它提供了相对其他开放源代码数据库系统 (比如 MySQL 和 Firebird)，和对专有系统比如 Oracle、Sybase、IBM 的 DB2
	//和 Microsoft SQL Server 的一种选择。
	//
	//PostgreSQL 和 MySQL 比较，它更加庞大一点，因为它是用来替代 Oracle 而设计的。所以在企业应用中采用 PostgreSQL 是一个明智的选择。
	//
	//MySQL 被 Oracle 收购之后正在逐步的封闭（自 MySQL 5.5.31 以后的所有版本将不再遵循 GPL 协议），鉴于此，
	//将来我们也许会选择 PostgreSQL 而不是 MySQL 作为项目的后端数据库。
	//

	//Go 实现的支持 PostgreSQL 的驱动也很多，因为国外很多人在开发中使用了这个数据库。
	//
	//github.com/lib/pq 支持 database/sql 驱动，纯 Go 写的
	//github.com/jbarham/gopgsqldriver 支持 database/sql 驱动，纯 Go 写的
	//github.com/lxn/go-pgsql 支持 database/sql 驱动，纯 Go 写的
	//

	//
	//CREATE TABLE userinfo
	//(
	//	uid serial NOT NULL,
	//	username character varying(100) NOT NULL,
	//	department character varying(500) NOT NULL,
	//	Created date,
	//	CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
	//)
	//WITH (OIDS=FALSE);
	//
	//CREATE TABLE userdetail
	//(
	//	uid integer,
	//	intro character varying(100),
	//	profile character varying(100)
	//)
	//WITH(OIDS=FALSE);

	db, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2020-06-30")
	checkErr(err)

	//pg 不支持这个函数，因为他没有类似 MySQL 的自增 ID
	//id, err := res.LastInsertId()
	//checkErr(err)
	//fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, departname, created) VALUES($1, $2, $3) returning uid;",
		"astaxie", "研发部门", "2020-06-30").Scan(&lastInsertId)

	checkErr(err)
	fmt.Println("最后插入id=", lastInsertId)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}
