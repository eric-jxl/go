package unittest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// ConnectMysql() Mysql Select data
func ConnectMysql() string {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/migu")
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	var (
		session_key string
	)
	if rows, err := db.Query("SELECT session_key FROM ops.django_session;"); err == nil {
		fmt.Println(">.>>>>>")
		for rows.Next() {
			err = rows.Scan(session_key)
			if err != nil {
				return err.Error()
			}
			fmt.Println(">.>>>")
		}
	}
	return session_key
}
