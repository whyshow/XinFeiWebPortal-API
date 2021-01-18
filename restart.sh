#!/bin/bash

pid=`ps aux|grep XinFeiWebPortal-API|grep -v grep|awk '{print $2}'`
kill -9 $pid
chmod 777 XinFeiWebPortal-API
tiem=$(date "+%Y%m%d")
#创建文件夹

if [ -d "./logs/$tiem" ];then
      nohup ./XinFeiWebPortal-API >>logs/$tiem/$tiem.log 2>&1 &
    else
        mkdir -p ./logs/$tiem/
        nohup ./XinFeiWebPortal-API >>logs/$tiem/$tiem.log 2>&1 &
    fi
