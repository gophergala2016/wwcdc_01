#!/bin/bash
killall gophergala-learning-resources-server 
cd /home/gophergala/gocode/src/github.com/gophergala2016/wwcdc_01
echo "starting server at `date`" >> service.log
/home/gophergala/gocode/src/github.com/gophergala2016/wwcdc_01/installserver.sh && /home/gophergala/gocode/bin/gophergala-learning-resources-server --tls-certificate /home/gophergala/boolmeover.com.pem --tls-key /home/gophergala/boolmeover.com.key --tls-port 13337 --tls-host boolmeover.com >>service.log 2>&1 &
