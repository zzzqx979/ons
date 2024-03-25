FROM ubuntu:latest
LABEL authors="Administrator"

# 复制文件到镜像中
COPY . /app

# 设置工作目录
WORKDIR /app

ENTRYPOINT ["top", "-b"]