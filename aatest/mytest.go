package aatest

import (
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	ZeroBot.OnGroupMessage().SetBlock(true).FirstPriority().Handle(func(ctx *ZeroBot.Ctx) {
		if ctx.Event.GroupID == 925535081 && ctx.Event.Message.String() == "获取密钥" {
			ctx.SendChain(message.Text("你的密钥是123"))
		}
	})
}
