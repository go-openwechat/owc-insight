package main

import (
	"math/rand"
	"time"

	"github.com/suntong/openwechat"
)

// wxHandshakeCheck will make sure to have recieved message within 2 min
func wxHandshakeCheck() {
	time.Sleep(100 * time.Second)
	t := time.Now()
	lr := lastReceivedRead()
	diff := t.Sub(lr)
	if diff > 2*time.Minute {
		abortOn("Lost WX ClientCheck handshake", ErrClientCheckLost)
		// try fresh HotLogin via external loop
	}
	logIf(1, "wx-clientcheck-passed", "gap", diff)
}

// Scheduled Executor
func periodicHotReload(bot *openwechat.Bot,
	self *openwechat.Self, reloadStorage openwechat.HotReloadStorage) {

	for true {
		lr := lastReceivedRead()
		// delay e.KaWait + e.KaVariety
		d := e.KaWait + rand.Intn(e.KaVariety)
		time.Sleep(time.Duration(d) * time.Minute)

		logIf(1, "scheduled-relogin")
		t := time.Now()
		diff := t.Sub(lr)
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

	lr := lastReceivedRead()
	t := time.Now()
	diff := t.Sub(lr)
	for true {
		lr = lastReceivedRead()
		t = time.Now()
		diff = t.Sub(lr)
		// delay ((e.KaWait + e.KaVariety) / e.KaBoost) - diff
		d := time.Minute *
			time.Duration(e.KaWait/e.KaBoost+rand.Intn(e.KaVariety/e.KaBoost))
		sleep := d - diff
		if sleep < 0 {
			// it last received was too long ago, > d, then disregard it
			sleep = d
		}
		logIf(2, "keep-alive-start", "gap", sleep)
		time.Sleep(sleep)

		lr = lastReceivedRead()
		t = time.Now()
		diff = t.Sub(lr)
		if diff < d {
			// too hot, wait for quieter time
			logIf(1, "keep-alive-skipped", "gap", diff)
			continue
		}

		r := rand.Intn(6)
		switch r {
		case 0:
			getFriends(self, true, 1)
			logIf(1, "keep-alive-done", "with", "getFriends", "rv", r)
		case 1, 2:
			getMps(self, true, 1)
			logIf(1, "keep-alive-done", "with", "getMps", "rv", r)
		case 3, 4:
			chatie.SendText("ding")
			logIf(1, "keep-alive-done", "with", "chatie.SendText", "rv", r)
		case 5:
			fallthrough
		default:
			getGroups(self, true, 2)
			logIf(1, "keep-alive-done", "with", "getGroups", "rv", r)
		}
	}
}

func lastReceivedUpdate() {
	lrSync.Lock()
	lastReceived = time.Now()
	lrSync.Unlock()
}

func lastReceivedRead() time.Time {
	lrSync.Lock()
	r := lastReceived
	lrSync.Unlock()
	return r
}
