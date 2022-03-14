package rtc

import (
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/log"
	api "open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
)

func (s *LiveSignaling) SetDefaultReq(req *api.InvitationInfo) {
	if req.RoomID == "" {
		req.RoomID = utils.OperationIDGenerator()
	}
	if req.Timeout == 0 {
		req.Timeout = 60 * 60
	}
	req.InviterUserID = s.loginUserID
}

func (s *LiveSignaling) InviteInGroup(signalInviteInGroupReq string, callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		log.Error(operationID, "callback is nil")
		return
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalInviteInGroupReq)
		req := &api.SignalReq_InviteInGroup{InviteInGroup: &api.SignalInviteInGroupReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalInviteInGroupReq, req, callback, operationID)
		s.SetDefaultReq(req.InviteInGroup.Invitation)
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback: finished")
	}()
}

func (s *LiveSignaling) Invite(callback open_im_sdk_callback.Base, signalInviteReq string, operationID string) {
	if callback == nil {
		log.Error(operationID, "callback is nil")
		return
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalInviteReq)
		req := &api.SignalReq_Invite{Invite: &api.SignalInviteReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalInviteReq, req, callback, operationID)
		s.SetDefaultReq(req.Invite.Invitation)
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback: finished")
	}()
}

func (s *LiveSignaling) Accept(callback open_im_sdk_callback.Base, signalAcceptReq string, operationID string) {
	if callback == nil {
		log.Error(operationID, "callback is nil")
		return
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalAcceptReq)
		req := &api.SignalReq_Accept{Accept: &api.SignalAcceptReq{Invitation: &api.SignalInviteReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalAcceptReq, req, callback, operationID)
		s.SetDefaultReq(req.Accept.Invitation.Invitation)
		req.Accept.InviteeUserID = s.loginUserID
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback finished")
	}()
}

func (s *LiveSignaling) Reject(callback open_im_sdk_callback.Base, signalRejectReq string, operationID string) {
	if callback == nil {
		log.NewError(operationID, "callback is nil")
		return
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalRejectReq)
		req := &api.SignalReq_Reject{Reject: &api.SignalRejectReq{Invitation: &api.SignalInviteReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalRejectReq, req, callback, operationID)
		s.SetDefaultReq(req.Reject.Invitation.Invitation)
		req.Reject.InviteeUserID = s.loginUserID
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback finished")
	}()
}

func (s *LiveSignaling) Cancel(callback open_im_sdk_callback.Base, signalCancelReq string, operationID string) {
	if callback == nil {
		log.NewError(operationID, "callback is nil")
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalCancelReq)
		req := &api.SignalReq_Cancel{Cancel: &api.SignalCancelReq{Invitation: &api.SignalInviteReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalCancelReq, req, callback, operationID)
		s.SetDefaultReq(req.Cancel.Invitation.Invitation)
		req.Cancel.InviterUserID = s.loginUserID
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback finished")
	}()
}

func (s *LiveSignaling) HungUp(callback open_im_sdk_callback.Base, signalHungUpReq string, operationID string) {
	if callback == nil {
		log.NewError(operationID, "callback is nil")
	}
	if s.listener == nil {
		log.Error(operationID, "listener is nil")
		callback.OnError(3004, "listener is nil")
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", signalHungUpReq)
		req := &api.SignalReq_HungUp{HungUp: &api.SignalHungUpReq{Invitation: &api.SignalInviteReq{Invitation: &api.InvitationInfo{}, OfflinePushInfo: &api.OfflinePushInfo{}}}}
		var signalReq api.SignalReq
		common.JsonUnmarshalCallback(signalHungUpReq, req, callback, operationID)
		s.SetDefaultReq(req.HungUp.Invitation.Invitation)
		req.HungUp.UserID = s.loginUserID
		signalReq.Payload = req
		s.handleSignaling(&signalReq, callback, operationID)
		log.NewInfo(operationID, fName, " callback finished")
	}()
}