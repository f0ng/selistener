# selistener
## 参考项目 https://github.com/fuzz7j/JNDIServer

用于解决判断出网情况的问题，以http与ldap形式批量监听端口，打印端口连接日志.

用法:

```
-ps , ports start    端口组监听，开始
-pe , ports end      端口组监听，结束
-pn ,port ,exp:22,3306,8443  指定端口监听
```



### 场景一

在log4j2(CVE-2021-44228)下，一般JNDIExpliot的1389等ldap的端口会禁止，那么可以使用该工具进行批量端口监听，然后在数据包内设置端口号为变量，查看端口通信接收情况(ldap),举例ldap监听情况:

<img width="529" alt="image" src="https://user-images.githubusercontent.com/48286013/212856092-7b326382-9116-48b8-93c8-77eac229c6e7.png">


### 场景二

在进行内网渗透中，传输文件的方法之一就是使用python3`python3 -m http.server 8000`、python2`python2 -m SimpleHTTPServer 8000`搭建http服务，但是可能会遇到ACL开放端口的问题，导致无法访问开启的http服务，如下：

<img width="568" alt="image" src="https://user-images.githubusercontent.com/48286013/212814647-881e705a-5b96-4ab7-830b-ecaa2c6bf7bc.png">

可以直接通过iptables去看端口的情况，但是避免不了端口占用的情况，所以可以使用该工具进行批量端口监听，查看端口通信接收情况(http),举例http监听情况:

<img width="613" alt="image" src="https://user-images.githubusercontent.com/48286013/212856583-f43dbb7e-2ce4-446d-92d5-5c7c3da322fd.png">


### 场景三

在进行反弹shell的时候，有时会禁止一些目的端口访问，可以通过使用该工具进行批量端口监听，查看端口通信接收情况(socket)。

<img width="677" alt="image" src="https://user-images.githubusercontent.com/48286013/212856726-342c12e5-b1e9-4a6d-a47c-04b91a8785c1.png">
