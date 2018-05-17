

# bindata安装

https://github.com/jteeuwen/go-bindata#installation
```SHELL
go get -u github.com/jteeuwen/go-bindata/...
```

# XORM 生成 model
```SHELL
cd src/github.com/foxiswho/shop-go/
xorm reverse mysql root:root@/shop_go?charset=utf8 template/design/goxorm
```

# 使用 casbin 作为后台权限控制
http://casbin.org/
https://github.com/casbin/casbin
https://github.com/casbin/xorm-adapter

```SHELL
go get github.com/casbin/casbin
go get github.com/casbin/xorm-adapter


```
https://github.com/labstack/echo-contrib
案例
https://studygolang.com/articles/12323
https://zupzup.org/casbin-http-role-auth/
https://github.com/zupzup/casbin-http-role-example

# 案例测试(本地测试)
## 修改本地host
修改
```SHELL
127.0.0.1 a.net b.net
```