pkill main
cd /home/ubuntu/app/Bot/bot_31808/leafBotPlugin
git checkout .
git pull
go mod tidy
go build main.go
nohup ./main > bot.log 2>&1 &