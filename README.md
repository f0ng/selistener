### 参考项目 https://github.com/fuzz7j/JNDIServer

# selistener

用于解决判断出网情况的问题，以http、ldap、rmi以及socket形式批量监听端口，在web界面进行结果查看，结果呈现形式类似dnslog，可用于内网log4j(CVE-2021-44228)等漏洞的检测，默认访问`http://x.x.x.x:65535/resp`即可查看到请求

- http请求

<img width="506" alt="image" src="https://user-images.githubusercontent.com/48286013/225658757-fd73c497-65eb-4962-888b-3537c062f1cd.png">


- ldap请求

<img width="593" alt="image" src="https://user-images.githubusercontent.com/48286013/225658473-2f2cb9d6-f676-4fe2-909b-c2c8214202d3.png">

- rmi请求

<img width="524" alt="image" src="https://user-images.githubusercontent.com/48286013/225659304-bf29616b-d275-4634-9881-96507fe34fe8.png">



用法:

```
-ps , ports start    端口组监听，开始
-pe , ports end      端口组监听，结束
-pn ,port ,exp:22,3306,8443  指定端口监听
-wp ,wport ,exp: 65535    指定端口运行http服务查看结果
```



### 场景一

在log4j2(CVE-2021-44228)下，一般JNDIExpliot的1389等ldap的端口会禁止，那么可以使用该工具进行批量端口监听，然后在数据包内设置端口号为变量，查看端口通信接收情况(ldap),举例ldap监听情况(rmi同理):

<img width="529" alt="image" src="https://user-images.githubusercontent.com/48286013/212856092-7b326382-9116-48b8-93c8-77eac229c6e7.png">


### 场景二

在windows系统无法上传文件情况下，需要进行文件下载，如certutil落地文件、powershell上线等等，会遇到某些端口无法出网，可以将certuil探测的端口设置为变量，使用本工具进行监听，如下

<img width="648" alt="image" src="https://user-images.githubusercontent.com/48286013/225650077-118fbe32-3762-42ef-8f33-15c8e1f9ef25.png">


### 场景三

在进行反弹shell的时候，有时会禁止一些目的端口访问，可以通过使用该工具进行批量端口监听，查看端口通信接收情况(socket)。

<img width="777" alt="image" src="https://user-images.githubusercontent.com/48286013/212856726-342c12e5-b1e9-4a6d-a47c-04b91a8785c1.png">
