package test

import (
	"encoding/json"
	"fmt"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"time"
)

type BaseSuccessFailed struct {
	successData string
	errCode     int
	errMsg      string
	funcName    string
	time        time.Time
}

func (b *BaseSuccessFailed) OnError(errCode int32, errMsg string) {
	b.errCode = -1
	b.errMsg = errMsg
	log.Error("login failed", errCode, errMsg)

}

func (b *BaseSuccessFailed) OnSuccess(data string) {
	b.errCode = 1
	b.successData = data
	log.Info("login success", data, time.Since(b.time))
}

func InOutDoTest(uid, tk, ws, api string) {
	var cf sdk_struct.IMConfig
	cf.ApiAddr = api
	cf.Platform = 1
	cf.WsAddr = ws
	cf.DataDir = "./"
	cf.LogLevel = LogLevel
	cf.ObjectStorage = "minio"
	cf.IsCompression = true
	cf.IsExternalExtensions = true

	b, _ := json.Marshal(cf)
	s := string(b)
	fmt.Println(s)
	var testinit testInitLister

	operationID := utils.OperationIDGenerator()
	if !open_im_sdk.InitSDK(&testinit, operationID, s) {
		fmt.Println("", "InitSDK failed")
		return
	}

	var testConversation conversationCallBack
	open_im_sdk.SetConversationListener(&testConversation)

	var testUser userCallback
	open_im_sdk.SetUserListener(testUser)

	var msgCallBack MsgListenerCallBak
	open_im_sdk.SetAdvancedMsgListener(&msgCallBack)

	var batchMsg BatchMsg
	open_im_sdk.SetBatchMsgListener(&batchMsg)

	var friendListener testFriendListener
	open_im_sdk.SetFriendListener(friendListener)

	var groupListener testGroupListener
	open_im_sdk.SetGroupListener(groupListener)

	var signalingListener testSignalingListener
	open_im_sdk.SetSignalingListener(&signalingListener)

	var workMomentsListener testWorkMomentsListener
	open_im_sdk.SetWorkMomentsListener(workMomentsListener)

	InOutlllogin(uid, tk)
}

func InOutlllogin(uid, tk string) {
	var callback BaseSuccessFailed
	callback.time = time.Now()
	callback.funcName = utils.GetSelfFuncName()
	operationID := utils.OperationIDGenerator()
	open_im_sdk.Login(&callback, operationID, uid, tk)
	for {
		if callback.errCode == 1 {
			return
		} else if callback.errCode == -1 {
			time.Sleep(100 * time.Millisecond)
		} else {
			time.Sleep(100 * time.Millisecond)
			log.Info(operationID, "waiting login ")
		}
	}
}

func InOutLogout() {
	var callback BaseSuccessFailed
	callback.funcName = utils.GetSelfFuncName()
	opretaionID := utils.OperationIDGenerator()
	open_im_sdk.Logout(&callback, opretaionID)
}