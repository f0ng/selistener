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
	"strings"
)

func GinAction( wport string ,token string) {
	//token := randStr(10) // 随机token
	//token := "f0ng"  // 自定义token
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	var dirnow = ""

	dirnow = "/selistener.db"


	str, _ := os.Getwd()
	fmt.Println(str + dirnow)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Thanks for using selistener!！Please add /resp to visit the results!\n\n\nGithub:https://github.com/f0ng/selistener")
	})

	r.GET("/resp", func(c *gin.Context) {
		words := c.DefaultQuery("words","")
		ipc := c.DefaultQuery("ip","")
		portc := c.DefaultQuery("port","")
		protocolc := c.DefaultQuery("protocol","")
		tokenc := c.DefaultQuery("token","")

		b := tokenc == token
		if (b) {

			database, err := sql.Open("sqlite3", str+dirnow+"?cache=shared&mode=rwc")
			if err != nil {
				log.Fatal(err)
			}

			var id int
			var ip string
			var port string
			var content string
			var time string
			var protocol string
			rows, err := database.Query("select id, protocol, ip, port, content, time  from notesprotocol where content like '%" + words + "%' and ip like '%" + ipc + "%' and port like '%" + portc + "%' and protocol like '%" + protocolc + "%'")
			if nil != err {
				fmt.Println(err)
			}
			var total string
			total = "["
			for rows.Next() {
				rows.Scan(&id, &protocol, &ip, &port, &content, &time)
				//fmt.Println(strings.Replace(strings.TrimSpace(content),"\\","/",-1))
				//fmt.Println(strings.Replace(strings.Replace(strings.TrimSpace(content),"\\","/",-1),"\r\n","",-1))
				total = total + "{\"" + strconv.Itoa(id) + "\":{\"ip\": \"" + ip + "\",\"port\":\"" + port + "\",\"protocol\":\"" + protocol + "\",\"content\":\"" + strings.Replace(strings.Replace(strings.TrimSpace(content),"\\","/",-1),"\r\n","",-1) + "\",\"time\":\"" + time + "\"}},"
				fmt.Println(strconv.Itoa(id) + ": " + ip + " " + content + " " + time)
			}
			total = string([]byte(total)[0:len(total)-1]) + "]"

			json_str := total
			var Ids []map[string]interface{}
			json.Unmarshal([]byte(json_str), &Ids)

			c.AsciiJSON(http.StatusOK, Ids)
		}else{
			c.String(http.StatusUnauthorized, "Error,Please input token!")
		}
	})


	fmt.Println("[default] web result port : " + wport + "\ntoken如下:" + token)
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":"+wport)
}



