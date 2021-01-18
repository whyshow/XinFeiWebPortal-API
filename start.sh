#!/bin/sh
tiem=$(date "+%Y%m%d")
chmod 777 XinFeiWebPortal-API
#创建文件夹

if [ -d "./logs/$tiem" ];then
      nohup ./XinFeiWebPortal-API >>logs/$tiem/$tiem.log 2>&1 &
    else
        mkdir -p ./logs/$tiem/
        nohup ./XinFeiWebPortal-API >>logs/$tiem/$tiem.log 2>&1 &
    fi
