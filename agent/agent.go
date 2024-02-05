package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/dollarkillerx/urllib"
)

var serviceAddress = "127.0.0.1:8383"

var gameNames = []string{"cs2.exe", "r5apex.exe", "steam.exe", "RainbowSix.exe", "r5apex_dx12.exe"}

type TBS struct {
	Stop bool `json:"stop"`
}

type JCS struct {
	Abbreviation string      `json:"abbreviation"`
	ClientIp     string      `json:"client_ip"`
	Datetime     string      `json:"datetime"`
	DayOfWeek    int         `json:"day_of_week"`
	DayOfYear    int         `json:"day_of_year"`
	Dst          bool        `json:"dst"`
	DstFrom      interface{} `json:"dst_from"`
	DstOffset    int         `json:"dst_offset"`
	DstUntil     interface{} `json:"dst_until"`
	RawOffset    int         `json:"raw_offset"`
	Timezone     string      `json:"timezone"`
	Unixtime     int         `json:"unixtime"`
	UtcDatetime  string      `json:"utc_datetime"`
	UtcOffset    string      `json:"utc_offset"`
	WeekNumber   int         `json:"week_number"`
}

func main() {
	fmt.Println("Agent Start")
	os.WriteFile("agent.lock", []byte("1 2 3"), 0644)

	go func() {
		for {
			time.Sleep(time.Second * 10)

			uri := "http://192.227.234.228:8782/info"

			var tbs TBS

			err := urllib.Get(uri).FromJson(&tbs)
			if err != nil {
				time.Sleep(time.Second * 5)
				continue
			}

			if tbs.Stop {
				Shutdown()
			}
		}
	}()

	for {
		time.Sleep(time.Second * 30)
		uri := "https://worldtimeapi.org/api/timezone/Asia/Tokyo"

		var jcs JCS
		err := urllib.Get(uri).FromJson(&jcs)
		if err != nil {
			fmt.Println("获取时间失败")
			continue
		}

		// 使用time包解析字符串
		parsedTime, err := time.Parse(time.RFC3339Nano, jcs.Datetime)
		if err != nil {
			fmt.Println("解析日期时间失败:", err)
			return
		}

		// 获取当前时间
		currentTime := time.Now()

		// 输出解析得到的时间和当前时间的小时和分钟部分
		fmt.Println("解析得到的时间:", parsedTime)
		fmt.Println("当前时间:", currentTime)

		// 获取解析得到的时间的小时部分
		hour := parsedTime.Hour()

		// 执行逻辑判断
		if hour >= 0 && hour < 6 {
			fmt.Println("当前时间在0-6点，执行打印helloworld的代码")

			var ex bool
			for _, v := range gameNames {
				if found(v) {
					Kill(v)
					ex = true
				}
			}
			if ex {
				Shutdown()
			}
		} else if hour >= 23 && parsedTime.Minute() >= 30 {
			fmt.Println("当前时间在晚上11:30-12:00，执行打印helloworld2的代码")

			var ex bool
			for _, v := range gameNames {
				if found(v) {
					Kill(v)
					ex = true
				}
			}
			if ex {
				Shutdown()
			}
		} else {
			fmt.Printf("当前时间不在指定范围内，不执行特定逻辑 %d \n", hour)
		}
	}

	//for {
	//	time.Sleep(time.Second * 5)
	//
	//	now := time.Now()
	//
	//	// 计算早上 1 点
	//	oneOClock := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.Local)
	//
	//	// 计算早上 6 点
	//	sixOClock := oneOClock.Add(6 * time.Hour)
	//	six12OClock := oneOClock.Add(23 * time.Hour)
	//
	//	fmt.Println(oneOClock.Unix())
	//	fmt.Println(sixOClock.Unix())
	//	fmt.Println(six12OClock.Unix())
	//
	//	// 判断当前时间是否在 1 点到 6 点之间
	//	if now.Unix() >= oneOClock.Unix() && now.Unix() <= sixOClock.Unix() {
	//		fmt.Println("当前时间在 1 点到 6 点之间")
	//		var ex bool
	//		for _, v := range gameNames {
	//			if found(v) {
	//				Kill(v)
	//				ex = true
	//			}
	//		}
	//		if ex {
	//			Shutdown()
	//		}
	//	}
	//
	//	if now.Unix() >= six12OClock.Unix() {
	//		fmt.Println("当前时间在 12 点之间")
	//
	//		var ex bool
	//		for _, v := range gameNames {
	//			if found(v) {
	//				Kill(v)
	//				ex = true
	//			}
	//		}
	//		if ex {
	//			Shutdown()
	//		}
	//	}
	//}
}

func found(processName string) bool {
	cmd := exec.Command("taskkill", "/FI", fmt.Sprintf("IMAGENAME eq %s", processName))
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	stdout, _ := cmd.Output()
	if len(string(stdout)) == 40 {
		return false
	}

	return true
}

func Shutdown() {
	// 立即关闭计算机
	cmd := exec.Command("cmd", "/C", "shutdown", "/s")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func Kill(processName string) {
	// 创建一个 *exec.Cmd 对象
	cmd := exec.Command("taskkill", "/F", "/IM", processName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	stdout, _ := cmd.Output()

	fmt.Println(string(stdout))
}
