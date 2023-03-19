package UI

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin" // https://gin-gonic.com/zh-cn/docs/examples/ascii-json/
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GinAction( wport string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	os.Getwd()
	var dirnow = ""

	dirnow = "/selistener.db"


	str, _ := os.Getwd()
	fmt.Println(str + dirnow)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Thanks for using selistener!！Please add /resp to visit the results!\n\n\nGithub:https://github.com/f0ng/selistener")
	})

	r.GET("/resp", func(c *gin.Context) {
		database, err := sql.Open("sqlite3", str +  dirnow + "?cache=shared&mode=rwc")
		//stmt, _ := database.Prepare("create table if not exists notesprotocol(id integer primary key ,protocol text, ip text , port text, content text,time text)")
		//stmt.Exec()

		//stmt, _ = database.Prepare("insert into notesprotocol( protocol, ip, port, content, time) values(?, ? ,? , ? , ? )")
		//stmt.Exec("ldap","172.253.237.5:41578", "1234","/aaaa","2023-03-16 11:06:19")
		if err != nil {
			log.Fatal(err)
		}

		//defer database.Close()

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


	fmt.Println("[default] web result port : " + wport)
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":"+wport)
}