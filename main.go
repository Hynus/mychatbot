package main

import (
	"fmt"
	"mychatbot/robot"
	"mychatbot/show"
	"strings"
)

func main() {
	var myText string
	fmt.Println("*****************************聊天机器人*****************************")
	fmt.Print("++++robot：你好，现在可以开始和我聊天了！（输入bye或者再见即可退出）")
	for {
		fmt.Printf("\n++我：")
		fmt.Scanf("%s", &myText)
		if myText == "再见" || strings.ToUpper(myText) == "BYE" {
			fmt.Println("++++robot：再见！")
			break
		} else {
			resp, _ := robot.GetRobotResp(myText, "")
			show.ShowFinalResp(robot.ParseRespJson(resp))
		}
	}

}
