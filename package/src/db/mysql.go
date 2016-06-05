/**
 * 封装下mysql的常用操作
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Conmysql struct {
	db *sql.DB
}

/**
 * 数据库连接初始化
 * @param  {[type]} con *Conmysql)    Init(conf map[string]string [description]
 * @return {[type]}     [description]
 */
func (con *Conmysql) Init(conf map[string]string) {
	/* 组合连接语句 */
	constr := conf["user"] + ":" + conf["pass"] + "@tcp(" + conf["host"] + ":" + conf["port"] + ")/" + conf["db"] + "?charset=utf8"
	db, err := sql.Open("mysql", constr)
	if err != nil {
		fmt.Println("mysql connect error:", err.Error())
		return
	}
	con.db = db
}

/**
 * 查询单条信息
 * @param  {[type]} con *Conmysql)    Findfirst(columns string, where string [description]
 * @return {[type]}     [description]
 */
func (con *Conmysql) Findfirst(columns []string, where string, tablename string) (res map[string]string) {
	var columnstr string
	columnstr = strings.Join(columns, ",")
	if where == "" {
		where = "1 = 1"
	}
	if tablename == "" {
		fmt.Println("tablename is not allowed to by empty!")
		return
	}
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}
	query, err := con.db.Query("select " + columnstr + " from " + tablename + " where " + where)
	defer query.Close()
	if err != nil {
		fmt.Println("select error: ", err.Error())
		return
	}
	query.Next()
	if err := query.Scan(scans...); err != nil {
		fmt.Println("Error")
		return
	}
	res = map[string]string{}
	for i, v := range values {
		res[string(columns[i])] = string(v)
	}
	return res
}

func main() {
	var columns = []string{"id", "keyword"}
	conf := map[string]string{"host": "127.0.0.1", "user": "root", "pass": "123456", "port": "3306", "db": "tianyunzi"}
	con := Conmysql{}
	con.Init(conf)
	rs := con.Findfirst(columns, "id = 1", "keyword")
	fmt.Println(rs)
}
