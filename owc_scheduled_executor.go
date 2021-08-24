package main

import (
	"math/rand"
	"time"

	"github.com/suntong/openwechat"
)

// Scheduled Executor
func periodicHotReload(bot *openwechat.Bot,
	self *openwechat.Self, reloadStorage openwechat.HotReloadStorage) {

	for true {
		// delay e.KaWait + e.KaVariety
		d := e.KaWait + rand.Intn(e.KaVariety)
		time.Sleep(time.Duration(d) * time.Minute)

		logIf(1, "scheduled-relogin")
		t := time.Now()
		diff := t.Sub(lastReceived)
		if diff < 5*time.Minute {
			// too hot, wait for quieter time
			logIf(1, "scheduled-relogin-skipped", "gap", diff)
			continue
		}

		err := bot.HotLogin(reloadStorage, false)
		_abortOn("Can't restart bot", err, 9)
		logIf(1, "scheduled-relogin-ok", "user", self)
		postLogin(self)
	}
}

// periodically feed the watch dog
func periodicDogFeed(bot *openwechat.Bot,
	self *openwechat.Self) {
	// 获取当前用户所有的公众号
	mps := getMps(self, false, 1)
	chatie := mps.SearchByNickName(1, "Chatie")[0]
	logIf(1, "keep-alive", "chatie", chatie.User)

	// delay (e.KaWait + e.KaVariety) / e.KaBoost
	for true {
		d := e.KaWait/e.KaBoost + rand.Intn(e.KaVariety/e.KaBoost)
		time.Sleep(time.Duration(d) * time.Minute)

		logIf(2, "keep-alive")
		switch rand.Intn(7) {
		case 0, 1:
			getFriends(self, true, 1)
			logIf(1, "keep-alive-done", "with", "getFriends")
		case 2, 3:
			getMps(self, true, 1)
			logIf(1, "keep-alive-done", "with", "getMps")
		case 4, 5:
			chatie.SendText("ding")
			logIf(1, "keep-alive-done", "with", "chatie.SendText")
		case 6:
			fallthrough
		default:
			getGroups(self, true, 2)
			logIf(1, "keep-alive-done", "with", "getGroups")
		}
	}
}
