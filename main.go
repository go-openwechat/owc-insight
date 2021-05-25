package main

import (
	"fmt"
	"net/http"

	"github.com/eatMoreApple/openwechat"
)

type ResponseHooker struct{}

func (r ResponseHooker) BeforeRequest(req *http.Request) {}

func (r ResponseHooker) AfterRequest(response *http.Response, err error) {
	fmt.Println(response.Request.URL.Path)
	fmt.Println(response.Request.Header)
}

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.Caller.Client.AddHttpHook(ResponseHooker{})

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			fmt.Println("你收到了一条新的文本消息")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)

	// 阻塞主goroutine, 知道发生异常或者用户主动退出
	bot.Block()
}
