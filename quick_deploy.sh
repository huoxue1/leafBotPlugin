pkill main
git checkout .
git pull
echo $(date +%F%n%T) > log.log
go mod tidy
go build main.go
nohup ./main > bot.log 2>&1 &
