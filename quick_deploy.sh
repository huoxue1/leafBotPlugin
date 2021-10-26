#!/bin/bash
pkill main
sudo git checkout .
sudo git pull
/home/ubuntu/go/bin/go build main.go
nohup ./main > bot.log 2>&1 &
echo $(date +%F%n%T) > do.log

