package main

import (
	"fmt"
	"os/exec"
)

var serviceAddress = "127.0.0.1:8383"

var gameNames = []string{"cs2.exe", "r5apex.exe", "steam.exe"}

func main() {
	found("steam.exe")
}

func found(exename string) bool {
	sql := fmt.Sprintf(`tasklist /FI "IMAGENAME eq %s"`, exename)
	cmd := exec.Command(sql)
	stdout, _ := cmd.Output()

	fmt.Println(string(stdout))
	return false
}

func Shutdown() {
	// 立即关闭计算机
	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
