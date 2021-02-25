package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syohex/go-texttable"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/statics","./statics")
	r.StaticFile("/529c3fcc09d41.jpg","./statics/images/529c3fcc09d41.jpg")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"在线留言.html",gin.H{
			"title":"hello"+c.Request.Method,
		})
	})
	r.POST("/", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:@tcp(172.29.38.170:4000)/")
		fmt.Println("connect success")
		defer db.Close()
		//fmt.Println(c.)
		sql := `/*--user=root;--password=jijizhazha;--host=172.29.38.170;--port=3306;--check=1;*/
    inception_magic_start;
    use  test;
    create table t1(id int primary key);
    alter table t1 add index idx_id (id);
    create table t2(jid int primary key);
    inception_magic_commit;`

		rows, err := db.Query(sql)
		fmt.Println("querying")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(cols)
		tbl := &texttable.TextTable{}

		tbl.SetHeader(cols[0], cols[1], cols[2], cols[3], cols[4], cols[5], cols[6], cols[7], cols[8], cols[9], cols[10], cols[11])

		for rows.Next() {
			var order_id, affected_rows, stage, error_level, stage_status, error_message, sql, sequence, backup_dbname, execute_time, sqlsha1, backup_time []uint8
			err = rows.Scan(&order_id, &stage, &error_level, &stage_status, &error_message, &sql, &affected_rows, &sequence, &backup_dbname, &execute_time, &sqlsha1, &backup_time)
			tbl.AddRow(string(order_id), string(affected_rows), string(stage), string(error_level), string(stage_status), string(error_message), string(sql), string(sequence), string(backup_dbname), string(execute_time))
			// tbl.AddRow(string(nil_process(sqlsha1)))
		}
		fmt.Println(tbl.Draw())
		c.HTML(http.StatusOK,"在线留言.html",gin.H{
			"title": tbl.Draw(),
		})
	})
	r.Run()
}
