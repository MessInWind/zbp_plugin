package aatest

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	zero.OnFullMatch("你好").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("你好，我是机器人")
	})

	zero.OnKeyword("天气").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("今天的天气是晴天")
	})
}
