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
}

// periodically feed the watch dog
func periodicDogFeed(bot *openwechat.Bot,
	self *openwechat.Self) {
	// 获取当前用户所有的公众号
	mps := getMps(self, false, 1)
	chatie := mps.Search(1, func(mp *openwechat.Mp) bool { return mp.User.NickName == "Chatie" })
	logIf(1, "Chatie公众号", "rec", chatie[0].User)

	// delay (e.KaWait + e.KaVariety) / e.KaBoost
	for true {
		d := e.KaWait/e.KaBoost + rand.Intn(e.KaVariety/e.KaBoost)
		time.Sleep(time.Duration(d) * time.Minute)

		switch rand.Intn(5) {
		case 0, 1:
			getFriends(self, true, 1)
		case 2, 3:
			getMps(self, true, 1)
		// case 4, 5:
		// 	getMps(self, true, 1)
		// case 6:
		case 4:
			fallthrough
		default:
			getGroups(self, true, 2)
		}
	}
}
