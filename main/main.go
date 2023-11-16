package main

import (
	"SocialMediaMonitor_Tgbot"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"sync"
)

const (
	BotToken = "Your Bot Token"
	// ChatID : Your ChatID,uint.
	ChatID    = -0000000
	Twitter   = "Twitter UserName"
	DC        = "Discord Server Invite Link"
	TGChat    = "Link of TG Chat Group"
	TGChannel = "Link of TG Channel"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go SetCorn()
	wg.Wait()
}

func SetCorn() {
	c := cron.New()

	err := c.AddFunc("0 0 0,6,9,18 * * ?", func() {
		Monitor()
	})
	if err != nil {
		return
	}

	c.Start()

	select {}
}

func Monitor() {

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if err != nil {
		fmt.Println(err)
	}

	tw := SocialMediaMonitor_Tgbot.TwFollowersCount(Twitter)
	tgChat, chatOnline := SocialMediaMonitor_Tgbot.TGChatMembersCount(TGChat)
	tgChannel := SocialMediaMonitor_Tgbot.TGChannelMembersCount(TGChannel)
	dc, dcOnlie := SocialMediaMonitor_Tgbot.DCMembersCount(DC)

	twInt64 := int64(tw)
	tg1Int64, _ := strconv.ParseInt(tgChat, 10, 64)
	tg2Int64, _ := strconv.ParseInt(tgChannel, 10, 64)
	total := twInt64 + tg2Int64 + tg1Int64 + int64(dc)
	totalStr := strconv.FormatInt(total, 10)
	dcStr := strconv.FormatInt(int64(dc), 10)
	twStr := strconv.FormatInt(int64(tw), 10)
	dcOnlineStr := strconv.FormatInt(int64(dcOnlie), 10)

	text := "Twitter: " + twStr + ";\n" + "\n" + "Telegram Chat:  " + tgChat + ";\n" + "Online:	" + chatOnline + ";\n" + "\n" + "Telegram Announcement:  " + tgChannel + ";\n" + "\n" + "Discord:  " + dcStr + ";\n" + "Online:	" + dcOnlineStr + ";\n" + "\n" + "Total:  " + totalStr

	msg2 := tgbotapi.NewMessage(ChatID, text)
	msg2.ParseMode = tgbotapi.ModeMarkdown

	_, err = bot.Send(msg2)
}
