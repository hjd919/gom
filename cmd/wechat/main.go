package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hjd919/gom"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func main() {

	//获取officialAccount实例
	gomOA := gom.NewGomOA(&gom.GomOAConf{
		RedisCacheOpts: &cache.RedisOpts{ // 公众号cache配置-redis
			Host:        "101.200.41.141:4379",
			Password:    "Xiaozi527",
			Database:    0,
			MaxIdle:     0,
			MaxActive:   0,
			IdleTimeout: 0,
		},
		OfficialAccountConf: &offConfig.Config{ // 公众号配置
			AppID:          "wx693f4c620712348b",
			AppSecret:      "8b3e323836403ad73f78453f07d54334",
			Token:          "slsw0520",
			EncodingAESKey: "SQbvays80pu3G9tasQgGU8FIIX5duPvL4mOwSJb4M9q",
		},
	})
	officialAccount := gomOA.Get()

	exampleOffAccount := NewExampleOfficialAccount(officialAccount, gomOA)
	fmt.Println(exampleOffAccount)
	// r := gin.Default()
	// r.GET("/ping", exampleOffAccount.Serve)

	// http.ListenAndServe(":8989", r)
}

//ExampleOfficialAccount 公众号操作样例
type ExampleOfficialAccount struct {
	officialAccount *officialaccount.OfficialAccount
	gomOA           *gom.GomOA
}

//ExampleOfficialAccount new
func NewExampleOfficialAccount(officialAccount *officialaccount.OfficialAccount, gomOA *GomOA) *ExampleOfficialAccount {
	return &ExampleOfficialAccount{
		officialAccount: officialAccount,
		gomOA:           gomOA,
	}
}

//Serve 处理消息
func (ex *ExampleOfficialAccount) Serve(c *gin.Context) {
	// 传入request和responseWriter
	server := ex.officialAccount.GetServer(c.Request, c.Writer)
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	//设置接收消息的处理方法
	// server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
	// 	//TODO
	// 	//回复消息：演示回复用户发送的消息
	// 	text := message.NewText(msg.Content)
	// 	return &message.Reply{MsgType: message.MsgTypeImage, MsgData: text}

	// 	//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
	// 	//articles := []*message.Article{article1}
	// 	//news := message.NewNews(articles)
	// 	//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

	// 	//voice := message.NewVoice(mediaID)
	// 	//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

	// 	//
	// 	//video := message.NewVideo(mediaID, "标题", "描述")
	// 	//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

	// 	//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
	// 	//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

	// 	//多客服消息转发
	// 	//transferCustomer := message.NewTransferCustomer("")
	// 	//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
	// })

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Println("Serve Println, err=", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Println("Send Println, err=", err)
		return
	}
	return
}
