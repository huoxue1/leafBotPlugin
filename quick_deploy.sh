pkill main
git checkout .
git pull
go mod tidy
go build main.go
nohup ./main > bot.log 2>&1 &
echo $(date +%F%n%T) > do.log