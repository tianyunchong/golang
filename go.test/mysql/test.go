/**
 * 简单的数据库链接测试
 */
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type ConMysql struct {
    db *sql.DB
}

func (con *ConMysql) Init() (db *sql.DB) {
    db, err := sql.Open("mysql", "root:gc7232275@tcp(192.168.8.18:3306)/tianyunzi?charset=utf8")
    if err != nil {
        fmt.Println("database init error : ", err.Error())
        return
    }
    return db
}

func (con *ConMysql) Readinfo(db *sql.DB) {
    query, err := db.Query("select * from pinyin limit 20")
    if err != nil {
        fmt.Println("fetch data failed:", err.Error())
    }
    defer query.Close()
    cols, _ := query.Columns()
    values := make([]sql.RawBytes, len(cols))
    scans := make([]interface{}, len(cols))
    for i := range values {
        scans[i] = &values[i]
    }
    results := make(map[int]map[string]string)
    i := 0
    for query.Next() {
        if err := query.Scan(scans...); err != nil {
            fmt.Println("Error")
            return
        }
        row := make(map[string]string)
        for j, v := range values {
            key := cols[j]
            row[key] = string(v)
        }
        results[i] = row
        i++
    }
    // 打印结果
    for i, m := range results {
        fmt.Println(i)
        for k, v := range m {
            fmt.Println(k, " : ", v)
        }
        fmt.Println("========================")
    }
}

func (con ConMysql) Read() {
    fmt.Println(con.db)
}

func main() {
    con := ConMysql{}
    con.db = con.Init()
    defer con.db.Close()
    con.Readinfo(con.db)

}
