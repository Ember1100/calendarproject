#!/bin/bash

# 停止并删除旧的容器
CONTAINER_ID=$(docker ps -aqf "name=gin-app")
if [ -n "$CONTAINER_ID" ]; then
    docker stop $CONTAINER_ID
    echo "停止成功"
    docker rm $CONTAINER_ID
    echo "删除成功"
else
    echo "容器 gin-app 不存在"
fi

# 删除旧的镜像
IMAGE_ID=$(docker images -q gin-app)
if [ -n "$IMAGE_ID" ]; then
    docker rmi $IMAGE_ID
    echo "rmi 成功"
else
    echo "镜像 gin-app 不存在"
fi

# 构建新镜像
docker build -t gin-app .
echo " 构建新镜像 gin-app 成功"
# 查看容器日志文件路径
LOG_PATH=$(docker inspect --format='{{.LogPath}}' gin-app)
echo "容器日志文件路径：$LOG_PATH"
# 运行新容器
docker run --name gin-app -p 8016:8016 --privileged=true -d gin-app:latest
echo "启动成功"

# 查看容器日志
docker logs -fn 200 gin-app