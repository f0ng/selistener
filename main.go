package main

//ldap监听

import (
	hex2 "encoding/hex"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"

	//"os"
	"strings"
	"sync"
	"time"
)


var (
	res      []string
	port     int
	exitChan = make(chan int)
	wg       sync.WaitGroup
)



func main() {

	ports := []int{}
	// 定义几个变量，用于接收命令行的参数值
	var ps        int
	var pe    int
	var port        string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.IntVar(&ps, "ps", 65535, "ports start")
	flag.IntVar(&pe, "pe", 65535, "ports end")
	flag.StringVar(&port, "pn", "", "port, exp:22,3306,8443")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	// 输出结果
	//fmt.Println(ps)
	//fmt.Println(pe)

	if (ps < 65535 && pe < 65536 && ps < pe){

		portsStrings := NewSlice(ps, pe, 1)  // 批量端口
		ports = String2Int(portsStrings)
		fmt.Printf("ports start %v \n" , ps)
		fmt.Printf("ports end %v \n" , pe)

	} else if(ps == 65535 && pe == 65535 && port != ""){

		portsss := strings.Split(port , ",")
		ports = String2Int(portsss)

		fmt.Println("ports start \n" , ports)

	}else {

		ports = []int{9999, 9997, 999, 9981, 995, 9944, 9943, 994, 993, 992, 990, 9876, 9870, 9869, 9864, 9801, 98, 9711, 9700, 9668, 9653, 9600, 9595, 9530, 9527, 9446, 9444, 9443, 9418, 9334, 9333, 9306, 9300, 9295, 9292, 9229, 9200, 9191, 9151, 9100, 9091, 9090, 9083, 9080, 9051, 9050, 9042, 9030, 902, 9011, 9010, 901, 9009, 9008, 9007, 9006, 9005, 9004, 9003, 9002, 9001, 9000, 900, 8999, 8983, 898, 89, 8899, 8890, 8889, 8888, 8887, 8886, 8885, 8884, 8883, 8882, 8881, 8880, 888, 8834, 8800, 880, 88, 8765, 876, 873, 87, 8688, 8686, 8649, 8600, 86, 8567, 8554, 8546, 8545, 853, 8529, 8500, 85, 8480, 8444, 8443, 84, 8388, 8378, 8377, 8351, 8334, 8333, 8332, 83, 8291, 8222, 8200, 82, 8194, 8182, 8181, 8161, 8159, 8140, 8139, 8138, 8129, 8126, 8125, 8123, 8118, 8112, 8111, 81, 8099, 8098, 8097, 8096, 8095, 8094, 8093, 8092, 8091, 8090, 8089, 8088, 8087, 8086, 8085, 8084, 8083, 8082, 8081, 8080, 808, 8069, 8060, 8058, 8040, 8032, 8030, 8025, 8020, 8010, 801, 8009, 8008, 8007, 8006, 8005, 8004, 8003, 8002, 8001, 8000, 800, 80, 7911, 79, 789, 7788, 7780, 7779, 7778, 7777, 7776, 777, 771, 7676, 7657, 7634, 7548, 7547, 7537, 7500, 7493, 7479, 7474, 7443, 7402, 7401, 7288, 7272, 7199, 7187, 7180, 7171, 7170, 7145, 7144, 7100, 7080, 7077, 7071, 7070, 705, 7014, 7010, 7007, 7005, 7004, 7003, 7002, 7001, 7000, 70, 7, 6998, 6969, 69, 6881, 6868, 6782, 6780, 67, 6699, 6697, 6669, 6668, 6667, 6666, 6665, 6664, 666, 6600, 6590, 6588, 6581, 6565, 6560, 6544, 6502, 6488, 64738, 646, 6443, 6379, 6363, 636, 6346, 631, 626, 623, 62078, 620, 61616, 61613, 6103, 6082, 6080, 6068, 6060, 60443, 6010, 6009, 6008, 6007, 6006, 6005, 6004, 60030, 6003, 6002, 60010, 6001, 60001, 6000, 5986, 5985, 5984, 5938, 59110, 5903, 5902, 5901, 5900, 587, 5820, 5802, 5801, 5800, 5683, 5678, 5673, 5672, 564, 5632, 5631, 5598, 5577, 5560, 55555, 55553, 5555, 5554, 5550, 55442, 554, 548, 5443, 5432, 5427, 5405, 5400, 540, 5357, 5353, 5351, 53413, 53, 52869, 5280, 5269, 5258, 523, 5222, 520, 515, 512, 5111, 51106, 51, 5095, 5094, 5093, 5084, 5080, 5061, 5060, 5051, 5050, 505, 5038, 503, 502, 50111, 50100, 5010, 50090, 5009, 5008, 50075, 50070, 5007, 50060, 5006, 50050, 5005, 5004, 5003, 5002, 5001, 50000, 5000, 500, 49674, 49673, 49672, 49671, 49670, 49669, 49668, 49667, 49666, 49665, 49664, 4949, 49159, 49158, 49157, 49156, 49155, 49154, 49153, 49152, 49151, 4911, 49, 48899, 4880, 4848, 4842, 4840, 4800, 4786, 4782, 47808, 4730, 4712, 4711, 4664, 4660, 465, 4567, 45554, 4506, 4505, 4500, 449, 44818, 445, 4444, 4443, 4440, 444, 4433, 4430, 443, 44158, 4369, 4300, 43, 42873, 427, 4200, 41795, 4155, 4070, 4064, 4063, 4050, 4040, 4022, 40001, 40000, 4000, 391, 389, 38, 3790, 3784, 37810, 3780, 37777, 3749, 37215, 37020, 3702, 37, 3690, 3689, 3671, 36, 3567, 3542, 3541, 3531, 3528, 3525, 3524, 3523, 3522, 3520, 34964, 34963, 34962, 3478, 3460, 34599, 34567, 3443, 3391, 3390, 33890, 3389, 3388, 33848, 3372, 3352, 3337, 33338, 3333, 3312, 3311, 3310, 3307, 3306, 3299, 3288, 3283, 3280, 32773, 32771, 32770, 32768, 3260, 32414, 32412, 32400, 32, 31337, 3128, 311, 31, 3097, 3075, 30718, 3052, 3050, 30313, 30312, 30311, 30310, 3005, 3002, 3001, 30005, 30001, 3000, 30, 29999, 29876, 28784, 2869, 2809, 28080, 28017, 280, 2715, 27017, 27016, 27015, 26470, 264, 2638, 2628, 26214, 2604, 2601, 2600, 26, 25565, 2525, 25105, 25010, 2501, 25000, 25, 2480, 2455, 2443, 2427, 2425, 2424, 2406, 2404, 2401, 2396, 2379, 2376, 2375, 2362, 23424, 2332, 2323, 2306, 23023, 23, 2252, 22335, 2223, 22222, 2222, 22105, 22, 2181, 2160, 2154, 2152, 2123, 2121, 211, 21, 2096, 2095, 2094, 20880, 2087, 2086, 2083, 2082, 2080, 2077, 2064, 2055, 20547, 2053, 2052, 2051, 2049, 20332, 2030, 2022, 2020, 2010, 2002, 2001, 20005, 20000, 2000, 20, 19999, 1993, 1991, 199, 19888, 1967, 1962, 1947, 1935, 19150, 1911, 1901, 1900, 19, 1883, 1880, 1863, 1830, 18264, 18246, 18245, 1812, 18086, 18081, 18080, 18001, 18000, 17988, 179, 1777, 177, 175, 1741, 1723, 1720, 17185, 1701, 17000, 17, 16993, 16992, 16923, 16922, 1645, 162, 1610, 161, 1604, 16030, 16010, 16000, 1588, 1554, 1521, 1515, 1505, 1503, 15000, 15, 1494, 1471, 14534, 14443, 1443, 1434, 1433, 143, 14265, 14147, 14000, 1400, 139, 138, 13722, 13720, 137, 13666, 13579, 135, 1344, 1314, 1311, 1302, 13, 12999, 1290, 1260, 1248, 1241, 12345, 1234, 12300, 123, 1214, 1212, 121, 1201, 12000, 1200, 11965, 1194, 119, 1177, 11371, 11310, 11300, 113, 11211, 111, 11001, 110, 11, 1099, 1080, 10554, 10443, 1042, 104, 10333, 10332, 1027, 1026, 10255, 10250, 1025, 10243, 1024, 1023, 1022, 102, 10162, 1010, 10035, 10030, 10005, 10003, 10001, 10000, 1000}
		fmt.Println("default ports start \n" )
	}

	for _, v := range ports {
		go func(port int) { //每个端口都扔进一个goroutine中去监听
			var bytes []byte
			ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

			if err != nil {
				fmt.Println(err.Error())
				exitChan <- 1
			}
			defer ln.Close()
			for {

				conn, _ := ln.Accept()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fb, err := ReadTagAndLength(conn, &bytes)
				//fmt.Print(conn)
				if err == nil {
					go ChooseMode(fb, conn, port)
				}
			}
		}(v)
	}
	select {}
	//}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Host)
	fmt.Print(r.Host + "\n")
}

func ChooseMode(b byte, conn net.Conn,port int) {
	switch b {
	case 0x47:
		HttpServer(conn ,port)
	case 0x30:
		LDAPServer(conn ,port )
	//case 0x4a:
	//	RMIServer(conn)
	case 0x62:
		SOCKETServer(conn , port)
	default:
		//fmt.Print(b)
		conn.Close()
		break
	}
}


func ReadTagAndLength(conn net.Conn, bytes *[]byte) (fb byte, err error) {
	var b byte
	b, err = ReadBytes(conn, bytes, 1)
	if err != nil {
		return
	}
	return b, err
}

func ReadBytes(conn net.Conn, bytes *[]byte, length int) (b byte, err error) {
	newbytes := make([]byte, length)
	n, err := conn.Read(newbytes)
	if n != length {
		fmt.Errorf("%d bytes read instead of %d", n, length)
	} else if err != nil {
		return
	}
	*bytes = append(*bytes, newbytes...)
	b = (*bytes)[len(*bytes)-1]
	return
}


func HttpServer(conn net.Conn ,port int) {
	rev := make([]byte, 1024)
	_, err := conn.Read(rev)
	defer conn.Close()
	if err == nil {
		session := strings.Split(string(rev), " ")[1]
		fmt.Printf("port: %v \n%v HTTP Query \"%v\" From %v\n",port, time.Now().Format("2006-01-02 15:04:05"), strings.TrimSpace(session), conn.RemoteAddr())

	}
}

func SOCKETServer(conn net.Conn ,port int) {
	rev := make([]byte, 1024)
	_, err := conn.Read(rev)
	defer conn.Close()
	if err == nil {
		session := strings.Split(string(rev), " ")[1]
		fmt.Printf("port: %v \n%v SOCKET Query \"%v\" From %v\n",port, time.Now().Format("2006-01-02 15:04:05"), strings.TrimSpace(session), conn.RemoteAddr())

	}
}


func LDAPServer(conn net.Conn ,port int ) {

	rev := make([]byte, 1024)
	data, _ := hex2.DecodeString("300c02010161070a010004000400\n")
	_, err := conn.Read(rev)
	defer conn.Close()
	if err == nil {
		conn.Write(data)
		_, err := conn.Read(rev)
		if err == nil {
			res = append(res, string(rev[8:30]))
			fmt.Printf("port: %v \n%v LDAP Query \"%v\" From %v\n",port,time.Now().Format("2006-01-02 15:04:05"), strings.TrimSpace(string(rev[8:30])), conn.RemoteAddr())
		}
	}
}

func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

func NewSlice(start, end, step int) []string {
	if step <= 0 || end < start {
		return []string{}
	}
	s := make([]string, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, strconv.FormatInt(int64(start),10) )
		start += step
	}
	return s
}