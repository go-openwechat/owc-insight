////////////////////////////////////////////////////////////////////////////
// Program: owc-insight
// Purpose: OpenWeChat Insight
// Authors: Tong Sun (c) 2020-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	//"github.com/eatMoreApple/openwechat"
	"github.com/skip2/go-qrcode"
	"github.com/suntong/openwechat"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const desc = "OpenWeChat Insight"

type envConfig struct {
	LogLevel  string `env:"OWCI_LOG"`
	KaWait    int    `env:"OWCI_KA_WAIT" envDefault:"450"`   // keep-alive wait (in min)
	KaVariety int    `env:"OWCI_KA_VARIETY" envDefault:"60"` // ka variant (in min)
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "owc-insight"
	version  = "0.1.0"
	date     = "2021-07-20"

	e envConfig

	lastReceived time.Time
	lastError    time.Time
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

/*

   NOT WORKING!

type ResponseHooker struct{}

func (r ResponseHooker) BeforeRequest(req *http.Request) {}

func (r ResponseHooker) AfterRequest(response *http.Response, err error) {
	fmt.Println(response.Request.URL.Path)
	fmt.Println(response.Request.Header)
}

*/

//==========================================================================
// Main

func main() {
	// == Config handling
	err := env.Parse(&e)
	abortOn("Env config parsing error", err)
	if e.LogLevel != "" {
		di, err := strconv.ParseInt(e.LogLevel, 10, 8)
		abortOn("OWCI_LOG (int) parse error", err)
		debug = int(di)
	}

	logIf(0, desc,
		"Version", version,
		"Built-on", date,
	)
	logIf(0, "Copyright (C) 2020-2021, Tong Sun", "License", "MIT")
	logIf(0, "Program parameters",
		"log-level", e.LogLevel,
		"keep-alive-wait-min", e.KaWait,
		"keep-alive-variant-min", e.KaVariety,
	)

	bot := openwechat.DefaultBot(openwechat.Desktop)
	//bot.Caller.Client.AddHttpHook(ResponseHooker{})
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	var count int32
	bot.GetMessageErrorHandler = func(err error) {
		t := time.Now()
		if t.Sub(lastError) < 30*time.Minute {
			count++
		} else {
			count = 1
		}
		// 如果发生了三次错误,那么直接退出
		if count > 3 {
			abortOn("Too many errors", err)
		}
		logIf(0, "catch-and-skip", "count", count, "err", err)
		lastError = time.Now()
	}

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		lastReceived = time.Now()
		sender, err := msg.Sender()
		abortOn("Can't get sender", err)
		// 如果是群聊消息，该方法返回的是群聊对象(需要自己将User转换为Group对象)
		fromGroup := ""
		if msg.IsSendByGroup() {
			fromGroup = fmt.Sprintf("%s", sender)
			sender, err = msg.SenderInGroup()
			abortOn("Can't get sender in group", err)
		}
		fromUserName := fmt.Sprintf("%s", sender)

		logIf(0, "收到消息", "type",
			fmt.Sprintf("%d (%d,%d)", msg.MsgType, msg.AppMsgType, msg.SubMsgType),
			"from", fromUserName, "in", fromGroup,
			"content", "")
		if len(msg.Content) == 0 {
			fmt.Println(msg)
		} else {
			fmt.Println(msg.Content)
		}

		if msg.IsText() {
			if msg.Content == "ping" {
				msg.ReplyText("pong")
				fmt.Println("回文本消息", msg.Content)
			} else {
				fmt.Println("收到文本消息", msg.Content)
			}
		}
	}

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登陆, 不定长参数设置为true, 可在登录凭证失效后进行扫码登录
	err = bot.HotLogin(reloadStorage, true)
	_abortOn("Can't start bot", err, 9)

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	abortOn("Can't get self", err)
	logIf(0, "logged-on", "user", self)

	// == Start Scheduled Executor
	rand.Seed(time.Now().Unix())
	lastReceived = time.Now()
	lastError = time.Now()
	go func(reloadStorage openwechat.HotReloadStorage,
		self *openwechat.Self) {
		for true {
			// delay e.KaWait + e.KaVariety
			d := e.KaWait + rand.Intn(e.KaVariety)
			time.Sleep(time.Duration(d) * time.Minute)
			t := time.Now()
			diff := t.Sub(lastReceived)
			if diff < 5*time.Minute {
				// too hot, wait for quieter time
				logIf(1, "scheduled-relogin-skipped", "gap", diff)
				continue
			}

			err := bot.HotLogin(reloadStorage)
			_abortOn("Can't restart bot", err, 9)
			logIf(1, "scheduled-relogin", "user", self)
			postLogin(self)
		}
	}(reloadStorage, self)

	postLogin(self)

	// 阻塞主goroutine, 知道发生异常或者用户主动退出
	bot.Block()
}

//==========================================================================
// support functions

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToSmallString(true))
}

func postLogin(self *openwechat.Self) {
	// 获取所有的群组(最新的)
	groups, err := self.Groups(true)
	abortOn("Can't get groups", err)
	logIf(2, "groups", "list", fmt.Sprintf("%v", groups))

	// 获取所有的好友(最新的好友)
	friends, err := self.Friends(true)
	abortOn("Can't get friends", err)
	logIf(3, "friends", "list", fmt.Sprintf("%v", friends))

	// WX ClientCheck from 微信团队 will come within seconds after initially login
	// wait for ~2 minutes to confirm their arrival
	go func() {
		time.Sleep(100 * time.Second)
		t := time.Now()
		diff := t.Sub(lastReceived)
		if diff > 2*time.Minute {
			abortOn("Lost WX ClientCheck handshake", nil)
			// try fresh HotLogin via external loop
		}
		logIf(1, "wx-clientcheck-passed", "gap", diff)
	}()
}
