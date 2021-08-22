package main

import (
	"math/rand"
	"time"

	"github.com/suntong/openwechat"
)

// Scheduled Executor
func periodicHotReload(bot *openwechat.Bot,
	reloadStorage openwechat.HotReloadStorage, self *openwechat.Self) {

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
