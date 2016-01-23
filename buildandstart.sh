#!/bin/bash
killall gophergala-learning-resources-server 
cd /home/gophergala/gocode/src/github.com/gophergala2016/wwcdc_01
echo "starting server at `date`" >> service.log
/home/gophergala/gocode/src/github.com/gophergala2016/wwcdc_01/installserver.sh && /home/gophergala/gocode/bin/gophergala-learning-resources-server --tls-certificate /home/gophergala/mc1.massca.st.pem --tls-key /home/gophergala/mc1.massca.st.key --tls-port 13337 --tls-host mc1.massca.st >>service.log 2>&1 &
