package main

import (
	"chatroom/client/process"
	"fmt"
)

func main() {
	var num int
	isLoop := true
	for {
		fmt.Println("------------------------------------------")
		fmt.Println("\t请输入:")
		fmt.Println("\t\t1. 登陆账号")
		fmt.Println("\t\t2. 注册账号")
		fmt.Println("\t\t3. 退出程序")
		fmt.Println("------------------------------------------")
		fmt.Scanln(&num)
		switch num{
			case 1 :
				fmt.Println("正在准备登陆...")
				isLoop = false
			case 2 :
				fmt.Println("正在准备注册...")
				isLoop = false
			case 3 :
				fmt.Println("退出系统")
				isLoop = false
			default :
				fmt.Println("请重新输入")
		}
		if !isLoop {
			break
		}
	}
	if num == 1 {
		fmt.Println("请输入您的账号")
		var account, password string
		fmt.Scanln(&account)
		fmt.Println("请输入您的密码")
		fmt.Scanln(&password)

		pro := &process.UserProcess{}
		err := pro.Login(account, password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("登陆操作完成")
		}
	} else if num == 2 {
		fmt.Println("请输入要注册的账号")
		var account, password string
		fmt.Scanln(&account)
		fmt.Println("请输入该账号的密码")
		fmt.Scanln(&password)

		pro := &process.UserProcess{}
		err := pro.Register(account, password)
		if err != nil {
			
		}
	} else {
		return
	}

}