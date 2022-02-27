package course

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/huoxue1/leafbot"
	"github.com/huoxue1/leafbot/message"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func init() {
	cmd()
	readFile()
}

var (
	binds       = make(map[int64]string)
	lock        = sync.Mutex{}
	defaultFile = "19网工.yml"
	weekTable   = map[string]int{
		"一": 1,
		"二": 2,
		"三": 3,
		"四": 4,
		"五": 5,
		"六": 6,
		"七": 7,
		"日": 7,
	}
	ming = map[string]int{
		"明": 1,
		"后": 2,
		"今": 0,
		"昨": -1,
		"前": -2,
	}
)

func cmd() {
	plugin := leafbot.NewPlugin("课程表")
	//plugin.SetHelp(map[string]string{
	//	"课表":            "获取当前账号所绑定的课表的今日课表",
	//	"课表列表":          "查看bot所记录的课表列表",
	//	"绑定":            "绑定当前账号到一个课表，例如：绑定 19网工.yml",
	//	"我的绑定":          "查看当前账号绑定的列表",
	//	"前|昨|今|明|后天课表":  "查看对应的课表",
	//	"x周周x课表":        "获取对应课表，例如：8周周一课表",
	//	"xxxx年xx月xx日课表": "获取对应日期课表",
	//})
	// plugin.OnRegex(`增加第(\d+)周周(.*?)[\n]id:(\d+)[\n]节数:(\d+)[\n]老师:(.*?)[\n]地点:(.*?)[\n]`).AddHandle(
	//	func(ctx.Event leafbot.Event, bot leafbot.Api, state *leafbot.State) {
	//		week,err := strconv.Atoi(state.RegexResult[1])
	//		if err != nil {
	//			return
	//		}
	//		i,ok := weekTable[state.RegexResult[2]]
	//		if !ok {
	//			return
	//		}
	//		id, err := strconv.Atoi(state.RegexResult[3])
	//		if err != nil {
	//			return
	//		}
	//		jie, err := strconv.Atoi(state.RegexResult[4])
	//		if err != nil {
	//			return
	//		}
	//
	//
	//	})
	plugin.OnCommand("bind", leafbot.Option{
		Weight: 5,
		Block:  false,
		Allies: []string{"绑定"},
		Rules:  []leafbot.Rule{leafbot.OnlyToMe},
	}).Handle(func(ctx *leafbot.Context) {
		if len(ctx.State.Data) < 1 {
			ctx.Event.Send("请输入正确的参数")
			return
		}
		binds[ctx.Event.UserId] = ctx.State.Args[0]
		loadFile()
		ctx.Event.Send(message.Text("绑定成功"))
	})
	plugin.OnCommand("课表", leafbot.Option{
		Weight: 10,
		Block:  true,
		Allies: []string{"课程表"},
		Rules:  nil,
	}).Handle(func(ctx *leafbot.Context) {
		lock.Lock()
		value, ok := binds[int64(ctx.Event.UserId)]
		lock.Unlock()
		var file string
		if ok {
			file = value
		} else {
			file = defaultFile
		}
		week, day := getWeek(time.Now())
		log.Infoln(fmt.Sprintf("当前第%d周周%d", week, day))
		if len(ctx.State.Args) > 0 && ctx.State.Args[0] == "all" {
			dir, err := os.ReadDir("./config/course/")
			if err != nil {
				return
			}

			for _, entry := range dir {
				if entry.IsDir() {
					continue
				}
				course, err := getCourse(week, day, entry.Name())
				if err != nil {
					log.Errorln(err.Error())
					return
				}
				ctx.Event.Send(message.Message{message.Text(entry.Name()), message.Image("base64://" + draw(course))})
			}
			return
		}
		course, err := getCourse(week, day, file)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		log.Infoln(course)
		ctx.Event.Send(message.Image("base64://" + draw(course)))
	})

	plugin.OnRegex(`^(\d+)年(\d+)月(\d+)[号|日]课表`).Handle(func(ctx *leafbot.Context) {
		lock.Lock()
		value, ok := binds[int64(ctx.Event.UserId)]
		lock.Unlock()
		var file string
		if ok {
			file = value
		} else {
			file = defaultFile
		}

		t, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v-%v-%v 12:00:00", ctx.State.RegexResult[1], ctx.State.RegexResult[2], ctx.State.RegexResult[3]))
		if err != nil {
			return
		}
		week, day := getWeek(t)
		course, err := getCourse(week, day, file)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		ctx.Event.Send(message.Image("base64://" + draw(course)))
	})

	plugin.OnRegex(`^(.*?)天课表`).Handle(func(ctx *leafbot.Context) {
		lock.Lock()
		value, ok := binds[int64(ctx.Event.UserId)]
		lock.Unlock()
		var file string
		if ok {
			file = value
		} else {
			file = defaultFile
		}
		i, ok := ming[ctx.State.RegexResult[1]]
		if !ok {
			ctx.Event.Send(message.Text("请输入正确的查询"))
			return
		}
		week, day := getWeek(time.Now().AddDate(0, 0, i))
		if len(ctx.State.Args) > 0 && ctx.State.Args[0] == "all" {
			dir, err := os.ReadDir("./config/course/")
			if err != nil {
				return
			}

			for _, entry := range dir {
				if entry.IsDir() {
					continue
				}
				course, err := getCourse(week, day, entry.Name())
				if err != nil {
					log.Errorln(err.Error())
					return
				}
				ctx.Event.Send(message.Message{message.Text(entry.Name()), message.Image("base64://" + draw(course))})
			}
			return
		}
		course, err := getCourse(week, day, file)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		ctx.Event.Send(message.Image("base64://" + draw(course)))
	})

	plugin.OnRegex(`^(\d+)周周(.*?)课表`).Handle(func(ctx *leafbot.Context) {
		lock.Lock()
		value, ok := binds[int64(ctx.Event.UserId)]
		lock.Unlock()
		var file string
		if ok {
			file = value
		} else {
			file = defaultFile
		}
		week, err := strconv.Atoi(ctx.State.RegexResult[1])
		if err != nil {
			return
		}
		day, ok := weekTable[ctx.State.RegexResult[2]]
		if !ok {
			ctx.Event.Send("请输入正确的内容")
			return
		}
		if len(ctx.State.Args) > 0 && ctx.State.Args[0] == "all" {
			dir, err := os.ReadDir("./config/course/")
			if err != nil {
				return
			}

			for _, entry := range dir {
				if entry.IsDir() {
					continue
				}
				course, err := getCourse(week, day, entry.Name())
				if err != nil {
					log.Errorln(err.Error())
					return
				}
				ctx.Event.Send(message.Message{message.Text(entry.Name()), message.Image("base64://" + draw(course))})
			}
			return
		}
		course, err := getCourse(week, day, file)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		ctx.Event.Send(message.Image("base64://" + draw(course)))
	})

	plugin.OnCommand("我的绑定").Handle(func(ctx *leafbot.Context) {
		value, ok := binds[int64(ctx.Event.UserId)]
		if ok {
			ctx.Event.Send(message.Text("你的绑定信息为：\n" + value))
		} else {
			ctx.Event.Send(message.Text("你还未进行课表绑定"))
		}
	})

	plugin.OnCommand("课表列表", leafbot.Option{
		Weight: 10,
		Block:  false,
		Allies: nil,
		Rules:  nil,
	}).Handle(func(ctx *leafbot.Context) {
		dir, err := os.ReadDir("./config/course/")
		if err != nil {
			return
		}
		results := ""
		for _, entry := range dir {
			if entry.IsDir() {
				continue
			}
			results += entry.Name() + "\n"
		}
		ctx.Event.Send(message.Text(results))
	})
}

func readFile() {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.ReadFile("./config/course.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(file, &binds)
	if err != nil {
		return
	}
}

func loadFile() {
	lock.Lock()
	defer lock.Unlock()
	data, err := yaml.Marshal(&binds)
	if err != nil {
		return
	}
	err = os.WriteFile("./config/course.yml", data, 0o666)
	if err != nil {
		return
	}
}
