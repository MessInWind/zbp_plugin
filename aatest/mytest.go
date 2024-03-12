package aatest

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	//群聊和私聊都会触发
	zero.OnFullMatch("你好").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("你好，我是机器人 v0.0.8")
	})

	zero.OnKeyword("天气").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("今天的天气是晴天")
	})

	zero.OnRegex("^用法$").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("用法：\n群聊:\n1. 你好\n2. 天气\n私聊:\n1. 查看qq号\n2. 查看数据库\n3. 查看密钥\n4. 申请密钥")
	})
	//私聊时触发
	zero.OnFullMatch("查看qq号").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			qqID := ctx.Event.UserID
			ctx.Send(fmt.Sprintf("你的QQ号是: %d", qqID))
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
	zero.OnFullMatch("查看数据库").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "root:root@tcp(47.236.248.235:3306)/testDB")
			ctx.Send("正在连接数据库...")
			if err != nil {
				fmt.Println(err)
				ctx.Send("连接数据库失败, 请联系管理员")
				return
			}
			ctx.Send("连接数据库成功")
			defer db.Close()
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
	zero.OnFullMatch("查看密钥").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "root:root@tcp(47.236.248.235:3306)/testDB")
			ctx.Send("正在连接数据库...")
			if err != nil {
				fmt.Println(err)
				ctx.Send("连接数据库失败, 请联系管理员")
				return
			}
			ctx.Send("连接数据库成功")

			qqID := ctx.Event.UserID
			ctx.Send(fmt.Sprintf("你的QQ号是: %d", qqID))

			row := db.QueryRow("SELECT login_key FROM ticketServer_tickets WHERE email = ?", qqID)
			var loginKey string
			err2 := row.Scan(&loginKey)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					// 没有找到匹配的记录
					ctx.Send("你还未申请密钥, 请先申请密钥")
				} else {
					// 数据库错误
					fmt.Println(err)
					ctx.Send("查询数据库失败, 请联系管理员")
				}
				return
			}
			ctx.Send(fmt.Sprintf("你的登录密钥是: %s", loginKey))

			defer db.Close()
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
}
