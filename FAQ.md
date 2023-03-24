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
