package aatest

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	//群聊和私聊都会触发
	zero.OnFullMatch("你好").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("你好, 我是无尽人生Bot v0.1.3")
	})

	zero.OnKeyword("天气").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("今天的天气是晴天")
	})
	zero.OnFullMatch("账号注册教程").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("先加机器人好友，然后私聊机器人，输入 申请密钥 四个字，机器人会自动帮你申请。游戏账号是你的QQ号，密码是机器人发给你的密钥。")
	})

	zero.OnRegex("^用法$").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		ctx.Send("------用法------\n" +
			"1. 申请密钥\n" +
			"2. 查看密钥\n" +
			"3. 账号状态\n" +
			"4. 上次登录时间\n" +
			"5. 账号注册教程\n" +
			"私聊先加机器人好友！")
	})
	//私聊时触发
	zero.OnFullMatch("连接数据库").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "ticketServer:ticketServer@tcp(47.236.248.235:53306)/testDB")
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
			db, err := sql.Open("mysql", "ticketServer:ticketServer@tcp(47.236.248.235:53306)/testDB")
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
	zero.OnFullMatch("申请密钥").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "ticketServer:ticketServer@tcp(47.236.248.235:53306)/testDB")
			ctx.Send("正在连接数据库...")
			if err != nil {
				fmt.Println(err)
				ctx.Send("连接数据库失败, 请联系管理员")
				return
			}
			ctx.Send("连接数据库成功")
			//获取qq号
			qqID := ctx.Event.UserID
			ctx.Send(fmt.Sprintf("你的QQ号是: %d", qqID))
			//查询是否已经申请过密钥
			row := db.QueryRow("SELECT login_key FROM ticketServer_tickets WHERE email = ?", qqID)
			var loginKey string
			err2 := row.Scan(&loginKey)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					// 没有找到匹配的记录
					ctx.Send("正在帮你申请密钥")
					//-----------------申请密钥-----------------
					data := url.Values{}
					data.Set("name", "bot")
					data.Set("email", strconv.FormatInt(qqID, 10))
					data.Set("tags", "april_9")

					req, err3 := http.NewRequest("POST", "http://web.hyx520.asia/ticketServer/server.php?action=sell_ticket", strings.NewReader(data.Encode()))
					if err3 != nil {
						fmt.Println(err3)
						ctx.Send("http请求失败, 请联系管理员")
						return
					}

					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					req.Header.Set("Referer", "http://web.hyx520.asia/ticketServer/server.php")
					req.Header.Set("Accept-Encoding", "gzip, deflate")
					req.Header.Set("Host", "web.hyx520.asia")
					// req.Header.Set("Cookie", "ticketServer_cookie_password_hash=1710055413_3e088f97002df99b8f2235673381d7a3")

					client := &http.Client{}
					resp, err4 := client.Do(req)
					if err4 != nil {
						fmt.Println(err4)
						ctx.Send("服务器未响应, 请联系管理员")
						return
					}
					defer resp.Body.Close()

					fmt.Println("response")
					fmt.Println(resp)
					if resp.StatusCode == 200 {
						fmt.Println("success!")
						ctx.Send("申请密钥成功!")
					}

					//-----------------返回密钥-----------------
					row := db.QueryRow("SELECT login_key FROM ticketServer_tickets WHERE email = ?", qqID)
					var loginKey string
					err5 := row.Scan(&loginKey)

					if err5 != nil {
						if err5 == sql.ErrNoRows {
							// 没有找到匹配的记录
							ctx.Send("申请密钥失败, 请联系管理员")
						} else {
							// 数据库错误
							fmt.Println(err5)
							ctx.Send("查询数据库失败, 请联系管理员")
						}
						return
					}

					ctx.Send(fmt.Sprintf("你的登录密钥是: %s", loginKey))

				} else {
					// 数据库错误
					fmt.Println(err2)
					ctx.Send("查询数据库失败, 请联系管理员")
				}
				return
			}

			ctx.Send(fmt.Sprintf("你已申请过，你的登录密钥是: %s", loginKey))

			defer db.Close()
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
	zero.OnFullMatch("账号状态").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "ticketServer:ticketServer@tcp(47.236.248.235:53306)/testDB")
			ctx.Send("正在连接数据库...")
			if err != nil {
				fmt.Println(err)
				ctx.Send("连接数据库失败, 请联系管理员")
				return
			}
			ctx.Send("连接数据库成功")

			qqID := ctx.Event.UserID
			// ctx.Send(fmt.Sprintf("你的QQ号是: %d", qqID))

			row := db.QueryRow("SELECT blocked FROM ticketServer_tickets WHERE email = ?", qqID)
			var blocked int8
			err2 := row.Scan(&blocked)
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
			if blocked == 1 {
				ctx.Send("你的账号已被封禁")
				ctx.Send("若误封，请联系管理员解封")
			} else {
				ctx.Send("你的账号信用记录正常")
			}
			defer db.Close()
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
	zero.OnFullMatch("上次登录时间").SetBlock(true).FirstPriority().Handle(func(ctx *zero.Ctx) {
		if ctx.Event.MessageType == "private" {
			db, err := sql.Open("mysql", "ticketServer:ticketServer@tcp(47.236.248.235:53306)/testDB")
			ctx.Send("正在连接数据库...")
			if err != nil {
				fmt.Println(err)
				ctx.Send("连接数据库失败, 请联系管理员")
				return
			}
			ctx.Send("连接数据库成功")

			qqID := ctx.Event.UserID

			row := db.QueryRow("SELECT last_activity FROM ticketServer_tickets WHERE email = ?", qqID)
			var lastActivity string
			err2 := row.Scan(&lastActivity)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					// 没有找到匹配的记录
					ctx.Send("你还未申请密钥, 请先申请密钥")
				} else {
					// 数据库错误
					fmt.Println(err2)
					ctx.Send("查询数据库失败, 请联系管理员")
				}
				return
			}

			lastActivityTime, err3 := time.Parse("2006-01-02 15:04:05", lastActivity)
			if err3 != nil {
				// 处理错误
				fmt.Println(err3)
				ctx.Send("解析时间失败, 请联系管理员")
			}

			ctx.Send(fmt.Sprintf("你上次登录时间是: %s", lastActivityTime.Format("2006-01-02 15:04:05")))
			defer db.Close()
		} else {
			ctx.Send("请私聊我(加好友)")
		}
	})
}
