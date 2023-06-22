package main

import (
	"chatroom/client/process"
	"chatroom/client/utils/menu"
	"fmt"
)

func main() {
	var num int
	for {
		if num = menu.Home(); num != -1 {
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