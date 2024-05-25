#!/bin/bash

# 源代码文件所在路径, 以此来表示项目根路径
PROJECT_POSITION_PATH=$HOME/Documents/coder/go-resp/src/liusnew-blog-backend-server

# running rpc service
# 执行util rpc service
go run $PROJECT_POSITION_PATH/rpc/utils/utils.go -f $PROJECT_POSITION_PATH/rpc/utils/etc/utils.yaml & 
# 执行content rpc service
go run $PROJECT_POSITION_PATH/rpc/content/content.go -f $PROJECT_POSITION_PATH/rpc/content/etc/content.yaml &
# 执行analyzer rpc service
go run $PROJECT_POSITION_PATH/rpc/analyzer/analyzer.go -f $PROJECT_POSITION_PATH/rpc/analyzer/etc/analyzer.yaml &
# 执行user rpc  service
go run $PROJECT_POSITION_PATH/rpc/user/user.go -f $PROJECT_POSITION_PATH/rpc/user/etc/user.yaml &
# 执行authorization rpc service
go run $PROJECT_POSITION_PATH/rpc/authorization/authorization.go -f $PROJECT_POSITION_PATH/rpc/authorization/etc/authorization.yaml &

# running gateway service 
# 执行user gateway service 
# go run  $PROJECT_POSITION_PATH/api/user/user.go -f $PROJECT_POSITION_PATH/api/user/etc/user-api.yaml &
# 执行article gateway service 
# go run  $PROJECT_POSITION_PATH/api/article/article.go -f $PROJECT_POSITION_PATH/api/article/etc/article-api.yaml &
