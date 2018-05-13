

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