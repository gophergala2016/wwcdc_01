#!/bin/bash
cd /home/gophergala/src/github.com/gophergala2016/wwcdc_01/
x=`git pull origin master | grep Already | wc -l`
if [x -eq 1] then
  echo "up to date"
else 
  restart.sh
fi
