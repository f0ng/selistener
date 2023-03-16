package UI

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin" // https://gin-gonic.com/zh-cn/docs/examples/ascii-json/
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GinAction() {
	r := gin.Default()
	var dirnow = ""
	if find := strings.Contains(dirnow, "selistener.db"); find {
		dirnow = dirnow
	}else{
		dirnow = dirnow + "./selistener.db"
	}

	r.GET("/resp", func(c *gin.Context) {
		database, err := sql.Open("sqlite3", "file:" + dirnow + "?cache=shared&mode=rwc")
		stmt, _ := database.Prepare("create table if not exists notesprotocol(id integer primary key ,protocol text, ip text , port text, content text,time text)")
		stmt.Exec()

		//stmt, _ = database.Prepare("insert into notesprotocol( protocol, ip, port, content, time) values(?, ? ,? , ? , ? )")
		//stmt.Exec("ldap","172.253.237.5:41578", "1234","/aaaa","2023-03-16 11:06:19")
		if err != nil {
			log.Fatal(err)
		}

		defer database.Close()

		var id int
		var ip string
		var port string
		var content string
		var time string
		var protocol string
		rows, err := database.Query("select id, protocol, ip, port, content, time  from notesprotocol")
		if nil != err {
			fmt.Println(err)
		}
		var total string
		total = "["
		for rows.Next() {
			rows.Scan(&id, &protocol, &ip, &port, &content ,&time)
			total = total + "{\"" + strconv.Itoa(id) + "\":{\"ip\": \"" + ip + "\",\"port\":\"" + port + "\",\"protocol\":\"" + protocol +"\",\"content\":\"" + content + "\",\"time\":\"" + time +"\"}},"
			fmt.Println(strconv.Itoa(id) + ": " + ip + " " + content + " " + time)
		}
		total = string([]byte(total)[0:len(total)-1]) + "]"

		json_str := total
		var Ids []map[string]interface{}
		json.Unmarshal([]byte(json_str), &Ids)

		c.AsciiJSON(http.StatusOK, Ids)
	})



	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":65535")
}