package mysql
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
		return res
	}
	query.Next()
	if err := query.Scan(scans...); err != nil {
		fmt.Println(err)
		return res
	}
	res = map[string]string{}
	for i, v := range values {
		res[string(columns[i])] = string(v)
	}
	return res
}

/**
 * mysql插入操作
 */
func (conn *Conmysql)Insert(data map[string]string, table string) int64{
	var keys []string
	var values []interface{}
	var rep []string
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
		rep = append(rep, "?")
	}
	//keyStr := strings.Join(keys, ",")
	//fmt.Println("INSERT "+table+" ("+keyStr+") values ("+strings.Join(rep, ",")+")")
	//return 0
	//stmt, err := conn.db.Prepare("INSERT "+table+" ("+keyStr+") values ("+strings.Join(rep, ",")+")")
	stmt, err := conn.db.Prepare(`INSERT cj_urls (url, addtime, uptime) values (?, ?, ?)`)
	if err != nil {
		fmt.Println("please check you sql!")
		return 0
	}
	fmt.Println(stmt)
	return 0
	res, err := stmt.Exec(values...)
	if (err != nil) {
		fmt.Println("insert is failed!")
		return 0
	}
	id, err := res.LastInsertId()
	if (err != nil) {
		fmt.Println("can not get the lastid")
		return 0
	}
	return id
}