package src

import (
	"database/sql"
	"fmt"
	"time"
)

func sqlite()  {
	//SQLite 是一个开源的嵌入式关系数据库，实现自包容、零配置、支持事务的 SQL 数据库引擎。
	//其特点是高度便携、使用方便、结构紧凑、高效、可靠。
	//与其他数据库管理系统不同，SQLite 的安装和运行非常简单，在大多数情况下，
	//只要确保 SQLite 的二进制文件存在即可开始创建、连接和使用数据库。
	//如果您正在寻找一个嵌入式数据库项目或解决方案，SQLite 是绝对值得考虑。SQLite 可以是说开源的 Access。
	//

	//驱动
	//Go 支持 sqlite 的驱动也比较多，但是好多都是不支持 database/sql 接口的
	// github.com/mattn/go-sqlite3 支持 database/sql 接口，基于 cgo (关于 cgo 的知识请参看官方文档或者本书后面的章节) 写的
	//github.com/feyeleanor/gosqlite3 不支持 database/sql 接口，基于 cgo 写的
	//github.com/phf/go-sqlite3 不支持 database/sql 接口，基于 cgo 写的
	//
	//目前支持 database/sql 的SQLite数据库驱动只有第一个，我目前也是采用它来开发项目的。
	//采用标准接口有利于以后出现更好的驱动的时候做迁移。
	//

	//
	//CREATE TABLE `userinfo` (
	//	`uid` INTEGER PRIMARY KEY AUTOINCREMENT,
	//	`username` VARCHAR(64) NULL,
	//	`department` VARCHAR(64) NULL,
	//	`created` DATE NULL
	//);
	//CREATE TABLE `userdetail` (
	//	`uid` INT(10) NULL,
	//	`intro` TEXT NULL,
	//	`profile` TEXT NULL,
	//	PRIMARY KEY (`uid`)
	//);

	//下面 Go 程序是如何操作数据库表数据：增删改查
}

func curd() {
	db, err := sql.Open("sqlite3", './foo.db')
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2020-06-30")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
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
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
	//sqlite 管理工具：sqliteadmin.orbmu2k.de/
	//
	//可以方便的新建数据库管理。
}