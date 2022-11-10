#                                 http-test-tool

## 介绍

http-test-tool是一款用于渗透测试中对40x页面进行bypass测试的扫描工具，通过多种方式尝试绕过状态码为40x的页面，采用go编写

## 使用说明

### 参数:

```
  -a, --AddHeader strings   the headers you will add,for explame cookie: 123,Referer: https://www.baidu.com
  -b, --bypassip strings    the ip you will add behind some header like x-client-ip: 192.168.1.1,for example,-b 192.168.1.1,30.1.1.1 and the default values are 127.0.0.1,localhost
  -h, --help                help 
  -p, --proxy string        the proxy you will use,support http and socks5 ,for example: -p http://127.0.0.1:8080 or -p socks5://127.0.0.1:8080
  -u, --url string          the target url
  -m, --urls string         bulk URLs
```



### 功能:

**其中bypass语句配置文件在http.yaml中,可自行添加或删除，若直接下载releases，请同时手动添加http.yaml**

1. 测试常见http方法

2. 添加常见的http头进行测试，分别使用GET和POST方法发送，如X-Host，X-Forward-For，X-Real-IP，Forwarded-For等

3. 在url最后添加部分字符进行测试

4. 在url中间添加字符进行测试

#### example:

简单用法:

```
./main -u http://www.google.com/admin/
```
批量urls用法：
```
./main -m http://www.google.com/admin/,http://www.google.com/
```

全参数用法:

```
./main -u http://www.google.com/admin/ -a cookie:123,Origin:https://30.110.0.0 -p http://127.0.0.1:60090 -b 39.122.1.1
```



