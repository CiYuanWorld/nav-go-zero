## 综合
goctl api go -h

## 生成api服务
```bash
进入相关目录后，执行
goctl api go -api web.api -dir .

修改api文件定义后的操作是相同的
goctl api go -api web.api -dir .
```

## 运行相关服务
```bash
进入相关目录后，执行
go run web.go -f etc/web.yaml
```

## 修改本地模板文件
```bash
cd ~/.goctl
goctl template init
然后编辑对应的文件即可
```
