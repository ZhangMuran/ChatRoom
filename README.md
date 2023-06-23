# ChatRoom
海量用户通信系统-Golang

server模块分层结构

![](https://s3.bmp.ovh/imgs/2023/06/12/45e7d18bb5511ba3.png)

client模块分层结构

![](https://s3.bmp.ovh/imgs/2023/06/12/fe90f1182b906b51.png)



List to do

- 目前的上线用户查看依赖第一次上线时收到的数据，没有考虑后续用户上下线更新以及好友等情况
- 客户端登录和退出时，服务器记录下客户端的名字
- 客户端登录后页面也记录下用户名字
- 重构 message 包
- 离线通知
- bug: 客户端断开连接时服务器有几率报错，但不影响运行
- 客户端的userprocess没有成员变量，看看怎么优化