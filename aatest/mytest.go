package aatest

import (
	"fmt"

	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	//群聊和私聊都会触发
	zero.OnFullMatch("你好").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("你好，我是机器人")
	})

	zero.OnKeyword("天气").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("今天的天气是晴天")
	})

	zero.OnRegex("^用法$").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("用法：\n群聊：\n1. 你好\n2. 天气\n私聊：\n1. 查看qq号")
	})
	//私聊时触发
	zero.OnFullMatch("查看qq号").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			qqID := ctx.Event.UserID
			ctx.Send(fmt.Sprintf("你的QQ号是：%d", qqID))
		}
	})
}
