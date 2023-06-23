package process

import (
	"chatroom/client/utils/menu"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"os"
)

func ShowHome() (err error) {
	var num int
	for {
		for {
			if num = menu.Home(); num != -1 {
				break
			}
		}
		if err = processHomePage(num); err != nil {
			fmt.Println(err)
			return
		} 
	}
}

func processHomePage(num int) (err error){
	if num == 1 {
		fmt.Println("请输入您的账号")
		var account, password string
		fmt.Scanln(&account)
		fmt.Println("请输入您的密码")
		fmt.Scanln(&password)

		pro := &UserProcess{}
		err = pro.Login(account, password)
		if err != nil {
			return
		} else {
			fmt.Println("登陆操作完成")
		}
	} else if num == 2 {
		fmt.Println("请输入要注册的账号")
		var account, password string
		fmt.Scanln(&account)
		fmt.Println("请输入该账号的密码")
		fmt.Scanln(&password)

		pro := &UserProcess{}
		err = pro.Register(account, password)
		if err != nil {
			return
		} else {
			fmt.Println("注册成功，现在您可以使用账号进行登陆了")
		}
	} else {
		os.Exit(0)
	}
	return
}

func ShowLogin() (err error){
	var num int
	for {
		for {
			if num = menu.AfterLogin(); num != -1 {
				break
			}
		}
		if err = processLogin(num); err != nil {
			return
		}
	}
}

func processLogin(num int) (err error) {
	if num == 1 {
		MyInfo.ShowOnlineFriend()
	} else if num == 2 {
		fmt.Println("请输入你想对大家说的话：")
		var data string
		fmt.Scanln(&data)
		sms := SmsProcess{
			data: data,
		}
		sms.SendSms()
	}
	return
}

func ProcessMessage(m message.Message) (err error) {
	switch m.Type {
	case message.NotifyLoginType :
		err = processNotify(&m)
	case message.SmsSendType :
		err = processSms(&m)
	}
	return
}

func processNotify(m *message.Message) (err error) {
	var notifyMsg message.NotifyOnlineMessage
	err = json.Unmarshal([]byte(m.Content), &notifyMsg)
	if err != nil {
		return
	}

	if(notifyMsg.User.Status == message.OnLine) {
		fmt.Println(notifyMsg.User.Account, "刚刚上线了，快去找他聊聊吧")
		MyInfo.Friend[notifyMsg.User.Account] = notifyMsg.User.Account
	}

	return
}

func processSms(m *message.Message) (err error) {
	var smsMsg message.SmsMessage
	err = json.Unmarshal([]byte(m.Content), &smsMsg)
	if err != nil {
		return
	}
	fmt.Println(smsMsg.Content)
	return
}