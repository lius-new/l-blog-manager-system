#!/bin/bash

# export PROJECT_POSITION_PATH=$HOME/Documents/coder/go-resp/src/liusnew-blog-backend-server

# stop gateway service & stop rpc service
# ports=`ps -aux | grep 'go run' | awk 'NR > 1 {print $2}'`

# for port in $ports
# do
#     kill -9 $port
# done

for port in :8080 :8081 :8082 :8083 :8084 :8888 :8889; do
    for pid in `lsof -i $port | awk 'NR > 1 {print $2}'`;do
        kill -9 $pid >> /dev/null 
        echo "$pid端口服务关闭"
    done
done
