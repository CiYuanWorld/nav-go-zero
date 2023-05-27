## 综合
goctl api go -h

## 运行相关服务
```bash
进入相关目录后，执行
go run web.go -f etc/web.yaml
```

## 生成api服务
```bash
进入相关目录后，执行
goctl api go -api web.api -dir .

修改api文件定义后的操作是相同的
goctl api go -api web.api -dir .
```

## 生成rpc服务
```bash
进入相关目录后，执行
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

## 生成model文件
```bash
进入相关目录后，执行
 -c --cache
goctl model mysql ddl -src ../../../../db/t_user.sql -dir .
goctl model mysql datasource -url="root@tcp(127.0.0.1:3306)/navapp" -table="t_link" -dir .
```

## 修改本地模板文件
```bash
cd ~/.goctl
goctl template init
然后编辑对应的文件即可
```

## 生成dockerfile
goctl docker --go app/nav-web/api/web.go
docker build -f Dockerfile -t navapp-webapi:latest .

## 生成kubernets配置文件
goctl kube deploy -name web-api -namespace navapp -port 11001 -o web-api.yaml
