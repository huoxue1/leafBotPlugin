package plugin_course

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v2"
)

type Course struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Teacher  string `json:"teacher"`
	Weeks    []int  `json:"weeks"`
	Time     int    `json:"times"`
	Location string `json:"location"`
}

func getCourseFromYaml(file string) (map[int][]Course, error) {
	contents := make(map[int][]Course, 7)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func getCourse(time time.Time) ([]Course, error) {
	week, day := getWeek(time)
	fmt.Println(week, day)
	xlsx, err := getCourseFromYaml("./config/course.yml")
	if err != nil {
		return nil, err
	}
	courses := xlsx[day]
	var cs []Course
	for _, course := range courses {
		for _, w := range course.Weeks {
			if w == week {
				cs = append(cs, course)
			}
		}
	}
	return cs, nil
}

var c = []string{"D", "F", "H", "K", "L", "O", "P"}

func parseXlsx(file string) (map[int][]Course, error) {
	contents := make(map[int][]Course, 7)
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	for index, s := range c {
		var courses []Course
		for i := 4; i <= 9; i++ {
			value, err := xlsx.GetCellValue("Sheet1", fmt.Sprintf("%v%d", s, i))
			if err != nil {
				return nil, err
			}
			//value = strings.ReplaceAll(value, " ", "")
			if value != "" {
				i2 := parse(value + "\n")
				for _, course := range i2 {
					courses = append(courses, course)
				}
			}
		}
		contents[index+1] = courses
	}
	toFile(contents)
	return contents, nil
}

func toFile(data map[int][]Course) {
	datas, err := yaml.Marshal(&data)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("course.yml", datas, 0666)
	if err != nil {
		return
	}
}

func parse(content string) []Course {
	var courses []Course
	content = strings.ReplaceAll(content, " ", "")
	compile := regexp.MustCompile(`\[(\d+)](.*?)[\n](.*?)\[(.*?)]\[(.*?)](.*?)[\n]`)
	strings := compile.FindAllStringSubmatch(content, -1)
	for _, s := range strings {
		times := parseTime(s[5])
		for _, time := range times {
			var course Course
			id, err := strconv.Atoi(s[1])
			if err != nil {
				continue
			}
			course.ID = id
			course.Name = s[2]
			course.Teacher = s[3]
			course.Weeks = parseWeek(s[4])
			course.Time = time
			course.Location = s[6]
			courses = append(courses, course)
		}
	}
	return courses
}

func parseTime(data string) []int {
	var times []int
	data = strings.ReplaceAll(data, "节", "")
	datas := strings.Split(data, "-")
	if len(datas) < 2 {
		start, _ := strconv.Atoi(datas[0])
		times = append(times, start)
		return times
	}
	start, err1 := strconv.Atoi(datas[0])
	if err1 != nil {
		return nil
	}
	end, err1 := strconv.Atoi(datas[1])
	if err1 != nil {
		return nil
	}
	for i := start; i <= end; i++ {
		times = append(times, i)
	}
	return times
}

func parseWeek(data string) []int {
	var weeks []int
	data = strings.ReplaceAll(data, "周", "")
	contents := strings.Split(data, ",")
	for _, content := range contents {
		dan := "all"
		if strings.Contains(content, "单") {
			dan = "dan"
			content = strings.ReplaceAll(content, "单", "")
		} else if strings.Contains(content, "双") {
			dan = "shuang"
			content = strings.ReplaceAll(content, "双", "")
		}

		week, err := strconv.Atoi(content)
		if err != nil {
			datas := strings.Split(content, "-")
			start, err1 := strconv.Atoi(datas[0])
			if err1 != nil {
				continue
			}
			end, err1 := strconv.Atoi(datas[1])
			if err1 != nil {
				continue
			}
			for i := start; i <= end; i++ {
				if (i%2 == 0 && dan == "shuang") || dan == "all" {
					weeks = append(weeks, i)
				} else if (i%2 != 0 && dan == "dan") || dan == "all" {
					weeks = append(weeks, i)
				}
			}
		} else {
			weeks = append(weeks, week)
		}
	}
	return weeks
}
