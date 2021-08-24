package main

import (
	"fmt"

	"github.com/suntong/openwechat"
)

// 处理文本消息
func textMessageHandle(msg *openwechat.Message) {
	lastReceivedUpdate()
	sender, err := msg.Sender()
	abortOn("Can't get sender", err)
	sendUser := sender.NickName
	// 如果是群聊消息，该方法返回的是群聊对象(需要自己将User转换为Group对象)
	if msg.IsSendByGroup() {
		// 取出消息在群里面的发送者
		sender, _ = msg.SenderInGroup()
		sendUser = fmt.Sprintf("[%s:%s]", sendUser, sender.NickName)
		// might not be able to get sender in group, ignore error
	}

	logIf(0, "收到消息", "type",
		fmt.Sprintf("%d (%d,%d)", msg.MsgType, msg.AppMsgType, msg.SubMsgType),
		"from", sendUser, "content", "")
	if len(msg.Content) == 0 {
		fmt.Println(msg)
	} else {
		fmt.Println(msg.Content)
	}

	if msg.IsText() {
		if msg.Content == "foo" {
			msg.ReplyText("bar")
			fmt.Println("回文本消息", msg.Content)
		} else {
			fmt.Println("收到文本消息", msg.Content)
		}
	}
}
