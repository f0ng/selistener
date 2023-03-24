## 延伸用法
1. 搭配log4j2burpscanner探测内网的Log4j(CVE-2021-44228)漏洞，参考文章https://mp.weixin.qq.com/s/NJ3gocQ_LojYlJk_0yWm6A

### 联动[log4j2burpscanner](https://github.com/f0ng/log4j2burpscanner)
配置如下:

```bash
10.211.x.x:9999/%20{HOSTURI}【ip改为内网ip】
```

```bash
http://10.211.x.x:65535/resp?token=f0ng&words={HOSTURI}【ip改为内网ip、token为生成的或者自定义的】
```
### 必须设置`{HOSTURI}`，这是定位到漏洞点的参数
### 必须设置`{HOSTURI}`，这是定位到漏洞点的参数
### 必须设置`{HOSTURI}`，这是定位到漏洞点的参数

<img width="776" alt="image" src="https://user-images.githubusercontent.com/48286013/227464369-fc1d90a8-9318-4b08-b4aa-be8d48784202.png">

启动环境，[Log4_demo-0.0.1-SNAPSHOT.jar](https://github.com/f0ng/selistener/blob/master/Log4_demo-0.0.1-SNAPSHOT.jar)

`java -jar Log4_demo-0.0.1-SNAPSHOT.jar`

漏洞数据包
```bash
GET /cvetext?cmd=1 HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
```

在数据包右键`send to passive scan`即可看到结果

<img width="770" alt="image" src="https://user-images.githubusercontent.com/48286013/227476958-34537046-1d19-4709-b9e7-7f86e9bca16b.png">
