pkill main
git checkout .
git pull
/usr/local/go/bin/go mod tidy
/usr/local/go/bin/go build main.go
nohup ./main > bot.log 2>&1 &
echo $(date +%F%n%T) > do.log