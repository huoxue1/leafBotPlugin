#!/bin/bash
screen -r leafBot
pkill main
sudo git checkout .
sudo git pull
/usr/local/go/bin/go build main.go
nohup ./main > bot.log 2>&1 &
echo $(date +%F%n%T) > do.log

