#!/bin/bash

pid=`ps aux|grep XinFeiWebPortal-API|grep -v grep|awk '{print $2}'`
kill -9 $pid