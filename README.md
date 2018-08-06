SHOP-GO 还在完善中
----------

# WIKI 
SHOP 说明介绍地址

[SHOP-GO WIKI](https://github.com/foxiswho/shop-go/wiki)


[https://github.com/foxiswho/shop-go/wiki](https://github.com/foxiswho/shop-go/wiki)




# ECHO GO SHOP Go (Echo Web)
Go web framework Echo example. 
RBAC权限,JWT、Socket,session,cookie,缓存,登录,注册,上传,db数据库操作,生成models,service演示

> 本案例是是 对 [echo-web](https://github.com/hb-go/echo-web) 的增强版，站在巨人的肩膀上

> Echo中文文档 [go-echo.org](http://go-echo.org/)

> Requires
- go1.8+
- Echo V3

# 启动

https://github.com/foxiswho/shop-go/wiki/90.run


##### 5.子域名
```shell
# ./conf/conf.toml
[server]
addr = ":8080"
domain_api = "echo.api.localhost.com"
domain_www = "echo.www.localhost.com"
domain_socket = "echo.socket.localhost.com"

# 改host
$ vi /etc/hosts
127.0.0.1       echo.api.localhost.com
127.0.0.1       echo.www.localhost.com
127.0.0.1       echo.www.localhost.com

# Nginx配置，可选
server{
    listen       80;
    server_name  echo.www.localhost.com echo.api.localhost.com echo.socket.localhost.com;

    charset utf-8;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host $host;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```


## Confd管理配置
```bash
# 安装confd
$ go get github.com/kelseyhightower/confd
```
```bash
# 将配置写入etcd，统一前缀echo-web
$ etcdctl ls --recursive --sort /echo-web
/echo-web/app
/echo-web/app/name
......
/echo-web/tmpl/suffix
/echo-web/tmpl/type

$ cd {pwd}/echo-web
$ confd -onetime -confdir conf  -backend etcd -node http://127.0.0.1:4001 -prefix echo-web
```

