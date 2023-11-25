package main

import (
	"fmt"
	"os/exec"
	"time"
)

var serviceAddress = "127.0.0.1:8383"

var gameNames = []string{"cs2.exe", "r5apex.exe", "steam.exe", "RainbowSix.exe"}

func main() {
	fmt.Println("Agent Start")
	for {
		time.Sleep(time.Second * 5)

		now := time.Now()

		// 计算早上 1 点
		oneOClock := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.Local)

		// 计算早上 6 点
		sixOClock := oneOClock.Add(6 * time.Hour)
		six12OClock := oneOClock.Add(12 * time.Hour)

		//fmt.Println(oneOClock.Unix())
		//fmt.Println(sixOClock.Unix())
		//fmt.Println(six12OClock.Unix())

		// 判断当前时间是否在 1 点到 6 点之间
		if now.Unix() >= oneOClock.Unix() && now.Unix() <= sixOClock.Unix() {
			//fmt.Println("当前时间在 1 点到 6 点之间")
			var ex bool
			for _, v := range gameNames {
				if found(v) {
					Kill(v)
				}

			}
			if ex {
				Shutdown()
			}
		}

		if now.Unix() >= six12OClock.Unix() {
			var ex bool
			for _, v := range gameNames {
				if found(v) {
					Kill(v)
				}

			}
			if ex {
				Shutdown()
			}
		}
	}
}

func found(processName string) bool {
	cmd := exec.Command("taskkill", "/FI", fmt.Sprintf("IMAGENAME eq %s", processName))
	stdout, _ := cmd.Output()
	if len(string(stdout)) == 40 {
		return false
	}

	return true
}

func Shutdown() {
	// 立即关闭计算机
	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func Kill(processName string) {
	// 创建一个 *exec.Cmd 对象
	cmd := exec.Command("taskkill", "/F", "/IM", processName)

	stdout, _ := cmd.Output()

	fmt.Println(string(stdout))
}
