#!/bin/bash

export PROJECT_POSITION_PATH=$HOME/Documents/coder/go-resp/src/liusnew-blog-backend-server

# running rpc service
go run $PROJECT_POSITION_PATH/rpc/utils/utils.go -f $PROJECT_POSITION_PATH/rpc/utils/etc/utils.yaml &
go run $PROJECT_POSITION_PATH/rpc/user/user.go -f $PROJECT_POSITION_PATH/rpc/user/etc/user.yaml &
go run $PROJECT_POSITION_PATH/rpc/authorization/authorization.go -f $PROJECT_POSITION_PATH/rpc/authorization/etc/authorization.yaml &

# running gateway service 
go run  $PROJECT_POSITION_PATH/api/user/user.go -f $PROJECT_POSITION_PATH/api/user/etc/user-api.yaml &
go run  $PROJECT_POSITION_PATH/api/article/article.go -f $PROJECT_POSITION_PATH/api/article/etc/article-api.yaml &