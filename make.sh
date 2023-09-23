#!/bin/bash
# 上面这一行指明了脚本使用的是 bash 解释器
rm -rf application.exe
# window 打包镜像
docker build -f .\deployment\dockerfile.yaml . -t changhao970520/fs
printf "执行打包go"
# 执行打包go
GOOS=linux go build -o foodSpider .\cmd\app\application.go
sleep 3
# 停掉并删除已经运行的容器
docker container stop gongyi
docker container rm  gongyi

# 删除已经存在的镜像
docker rmi changhao970520/gysrv


# 执行 Docker build
echo "Building Docker image..."
docker build -t changhao970520/gysrv .

# 运行容器
#docker run -p 6666:8888 -env=release --name gongyi -v /Users/hope/GolandProjects/cld-quick-cli/conf:/app/conf changhao970520/gysrv                                                                                                                     ─╯
docker run  --name gongyi -v /Users/hope/GolandProjects/cld-quick-cli/conf:/app/conf -p 6666:8888 changhao970520/gysrv
docker run  --name gongyi -v /home/gongyiMM/conf:/app/conf -v /home/gongyiMM/static:/app/static -p 9999:8888 changhao970520/gysrv # release

docker run --name spider -v C:\Users\19333\GolandProjects\foods-spider-server\conf:/app/conf  -p 9999:8888 fs