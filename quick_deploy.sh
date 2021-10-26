#!/bin/bash
screen -r leafBot
sudo git checkout .
sudo git pull
go build main.go
echo $(date +%F%n%T) > do.log

