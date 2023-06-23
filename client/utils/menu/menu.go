package menu

import (
	"fmt"
	"os"
)

func Home() int {
	fmt.Println("------------------------------------------")
	fmt.Println("\t请输入:")
	fmt.Println("\t\t1. 登陆账号")
	fmt.Println("\t\t2. 注册账号")
	fmt.Println("\t\t3. 退出程序")
	fmt.Println("------------------------------------------")
	var num int
	fmt.Scanln(&num)
	switch num {
		case 1 :
			fmt.Println("正在准备登陆...")
		case 2 :
			fmt.Println("正在准备注册...")
		case 3 :
			fmt.Println("退出系统")
		default :
			fmt.Println("请重新输入")
			num = -1
	}
	return num
}

func AfterLogin() int {
	fmt.Println("------------------------------------------")
	fmt.Println("\t\txxx，欢迎回来")
	fmt.Println("\t\t1. 显示在线用户列表")
	fmt.Println("\t\t2. 发送消息")
	fmt.Println("\t\t3. 消息列表")
	fmt.Println("\t\t4. 退出系统")
	fmt.Println("------------------------------------------")
	var num int
	fmt.Scanln(&num)
	switch num {
		case 1 :
			
		case 2 :
			fmt.Println("群聊模式")
		case 3 :
			fmt.Println("消息列表")
			num = -1
		case 4 :
			fmt.Println("退出系统")
			os.Exit(0)
		default :
			fmt.Println("请重新输入")
			num = -1
	}
	return num
}