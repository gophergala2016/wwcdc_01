#!/bin/bash
killall gophergala-learning-resources-server 
echo "$$" > service.pid
./installserver.sh && ./startserver.sh &
