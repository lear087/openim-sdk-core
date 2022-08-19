package test

import (
	"encoding/json"
	"log"
	"net/url"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"open_im_sdk/ws_wrapper/test/client"
	"open_im_sdk/ws_wrapper/ws_local_server"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var totalConnNum int
var lock sync.Mutex

func StartSimulationJSClient(api, jssdkURL, userID string, i int, userIDList []string) {
	// 模拟登录 认证 ws连接初始化
	user := client.NewIMClient("", userID, api, jssdkURL, 5)
	var err error
	user.Token, err = user.GetToken()
	if err != nil {
		log.Println("generate token failed", userID, api, err.Error())
	}
	v := url.Values{}
	v.Set("sendID", userID)
	v.Set("token", user.Token)
	v.Set("platformID", utils.IntToString(5))
	c, _, err := websocket.DefaultDialer.Dial(jssdkURL+"?"+v.Encode(), nil)
	if err != nil {
		log.Println("dial:", err.Error(), "userID", userID, "i: ", i)
		return
	}
	lock.Lock()
	totalConnNum += 1
	log.Println("connect success", userID, "total conn num", totalConnNum)
	lock.Unlock()
	user.Conn = c
	// user.WsLogout()
	user.WsLogin()
	time.Sleep(time.Second * 2)

	// 模拟登录同步
	go func() {
		err = user.GetSelfUserInfo()
		err = user.GetAllConversationList()
		err = user.GetBlackList()
		err = user.GetFriendList()
		err = user.GetRecvFriendApplicationList()
		err = user.GetRecvGroupApplicationList()
		err = user.GetSendFriendApplicationList()
		err = user.GetSendGroupApplicationList()
		if err != nil {
			log.Println(err)
		}
	}()

	// 模拟监听回调
	go func() {
		for {
			resp := ws_local_server.EventData{}
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err, "error an connet failed", userID)
				return
			}
			log.Printf("recv: %s", message)
			_ = json.Unmarshal(message, &resp)
			if resp.Event == "CreateTextMessage" {
				msg := sdk_struct.MsgStruct{}
				_ = json.Unmarshal([]byte(resp.Data), &msg)
				type Data struct {
					RecvID          string `json:"recvID"`
					GroupID         string `json:"groupID"`
					OfflinePushInfo string `json:"offlinePushInfo"`
					Message         string `json:"message"`
				}
				offlinePushBytes, _ := json.Marshal(server_api_params.OfflinePushInfo{Title: "push offline"})
				messageBytes, _ := json.Marshal(msg)
				data := Data{RecvID: userID, OfflinePushInfo: string(offlinePushBytes), Message: string(messageBytes)}
				err = user.SendMsg(userID, data)
				//fmt.Println(msg)
			}
		}
	}()

	// 模拟给随机用户发消息
	go func() {
		for {
			err = user.CreateTextMessage(userID)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			if err = user.GetLoginStatus(); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second * 10)
		}
	}()
}