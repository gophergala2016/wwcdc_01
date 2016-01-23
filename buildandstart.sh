#!/bin/bash
killall gophergala-learning-resources-server 
./installserver.sh && ./startserver.sh &
