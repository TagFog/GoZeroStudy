#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=User
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=172.27.44.42
port=3306
dbname=go-zero
username=root
passwd=20020110


echo "开始创建库：$dbname 的表：$tables"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=false --style=goZero