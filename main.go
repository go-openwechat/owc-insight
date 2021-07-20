////////////////////////////////////////////////////////////////////////////
// Program: wxsg
// Purpose: WX Statistic Genie
// Authors: Tong Sun (c) 2020-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/skip2/go-qrcode"
	"github.com/eatMoreApple/openwechat"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const desc = "WX Statistic Genie"

type envConfig struct {
	LogLevel      string `env:"WXSG_LOG"`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "wxsg"
	version  = "0.1.0"
	date     = "2021-07-15"

	e   envConfig

)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// Main

func main() {
	// == Config handling
	err := env.Parse(&e)
	abortOn("Env config parsing error", err)
	if e.LogLevel != "" {
		di, err := strconv.ParseInt(e.LogLevel, 10, 8)
		abortOn("WXSG_LOG (int) parse error", err)
		debug = int(di)
	}

	logIf(0, desc,
		"Version", version,
		"Built-on", date,
	)
	logIf(0, "Copyright (C) 2020-2021, Tong Sun", "License", "MIT")

	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			if msg.Content == "ping" {
				msg.ReplyText("pong")
				fmt.Println("回文本消息", msg.Content)
			} else  {
				fmt.Println("收到文本消息", msg.Content)
			}
		}
	}

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登陆
	err = bot.HotLogin(reloadStorage)
	abortOn("Can't start bot", err)

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	abortOn("Can't get self", err)
	logIf(0, "logged-on", "user", self)

	// 获取所有的群组
	groups, err := self.Groups()
	abortOn("Can't get groups", err)
	logIf(1, "groups", "list", fmt.Sprintf("%v", groups))

	// 获取所有的好友(最新的好友)
	friends, err := self.Friends(true)
	abortOn("Can't get friends", err)
	logIf(3, "friends", "list", fmt.Sprintf("%v", friends))

	// 阻塞主goroutine, 知道发生异常或者用户主动退出
	bot.Block()
}


//==========================================================================
// support functions

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
